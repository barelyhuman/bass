package bass

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"sync"

	"github.com/vito/bass/pkg"
	"go.uber.org/zap/zapcore"
)

// Scope contains bindings from symbols to values, and parent scopes to
// delegate to during symbol lookup.
type Scope struct {
	// an optional name for the scope, used to prettify .String on 'standard'
	// environments
	Name string

	Parents  []*Scope
	Bindings Bindings
	Order    []Symbol
}

// Bindings maps Symbols to Values in a scope.
type Bindings map[Symbol]Value

// NewScope constructs a new *Scope with initial bindings.
func (bindings Bindings) Scope(parents ...*Scope) *Scope {
	return NewScope(bindings, parents...)
}

// NewEmptyScope constructs a new scope with no bindings and
// optional parents.
func NewEmptyScope(parents ...*Scope) *Scope {
	return &Scope{
		Parents:  parents,
		Bindings: Bindings{},
	}
}

// NewScope constructs a new scope with the given bindings and
// optional parents.
func NewScope(bindings Bindings, parents ...*Scope) *Scope {
	scope := NewEmptyScope(parents...)
	for k, v := range bindings {
		scope.Set(k, v)
	}

	// ensure stable order since we preserve it in the output
	//
	// this should almost exclusively affect tests
	sort.Slice(scope.Order, func(i, j int) bool {
		return scope.Order[i].String() < scope.Order[j].String()
	})

	return scope
}

func Assoc(obj *Scope, kv ...Value) (*Scope, error) {
	clone := obj.Copy()

	var k Symbol
	var v Value
	for i := 0; i < len(kv); i++ {
		if i%2 == 0 {
			err := kv[i].Decode(&k)
			if err != nil {
				return nil, err
			}
		} else {
			err := kv[i].Decode(&v)
			if err != nil {
				return nil, err
			}

			clone.Set(k, v)

			k = ""
			v = nil
		}
	}

	return clone, nil
}

// Bindable is any value which may be used to destructure a value into bindings
// in a scope.
type Bindable interface {
	Value

	// Bind assigns values to symbols in the given scope.
	Bind(context.Context, *Scope, Cont, Value, ...Annotated) ReadyCont

	// EachBinding calls the fn for each symbol that will be bound.
	EachBinding(func(Symbol, Range) error) error
}

func BindConst(a, b Value) error {
	if !a.Equal(b) {
		return BindMismatchError{
			Need: a,
			Have: b,
		}
	}

	return nil
}

// Assert that Scope is a Value.
var _ Value = (*Scope)(nil)

func (value *Scope) String() string {
	if value.Name != "" {
		return fmt.Sprintf("<scope: %s>", value.Name)
	}

	if value.isPrinting() {
		return "{...}" // handle recursion or general noisiness
	}

	value.startPrinting()
	defer value.finishPrinting()

	bind := []Value{}

	_ = value.eachShadow(value, func(k Symbol, v Value) error {
		bind = append(bind, k.Keyword(), v)
		return nil
	}, map[Symbol]bool{})

	return formatList(NewList(bind...), "{", "}")
}

var printing = map[*Scope]struct{}{}
var printingL = new(sync.Mutex)

func (value *Scope) isPrinting() bool {
	printingL.Lock()
	defer printingL.Unlock()
	_, p := printing[value]
	return p
}

func (value *Scope) startPrinting() {
	printingL.Lock()
	defer printingL.Unlock()
	printing[value] = struct{}{}
}

func (value *Scope) finishPrinting() {
	printingL.Lock()
	defer printingL.Unlock()
	delete(printing, value)
}

func (value *Scope) IsSubsetOf(other *Scope) bool {
	var errMismatch = errors.New("mismatch")
	err := value.Each(func(k Symbol, v Value) error {
		ov, found := other.Get(k)
		if !found || !v.Equal(ov) {
			// use an error to short-circuit
			return errMismatch
		}

		return nil
	})
	if err != nil {
		return false
	}

	return true
}

func (value *Scope) Equal(o Value) bool {
	var other *Scope
	if err := o.Decode(&other); err != nil {
		return false
	}

	if other == value {
		return true
	}

	return value.IsSubsetOf(other) && other.IsSubsetOf(value)
}

func (value *Scope) IsEmpty() bool {
	empty := true
	_ = value.Each(func(Symbol, Value) error {
		empty = false
		return errors.New("im convinced")
	})

	return empty
}

func (value *Scope) Decode(dest any) error {
	switch x := dest.(type) {
	case **Scope:
		*x = value
		return nil
	case *Scope:
		*x = *value
		return nil
	case *Value:
		*x = value
		return nil
	case *zapcore.ObjectMarshaler:
		*x = value
		return nil
	case Decodable:
		return x.FromValue(value)
	case Value:
		return DecodeError{
			Source:      value,
			Destination: dest,
		}
	default:
		rt := reflect.TypeOf(dest)
		if rt.Kind() != reflect.Ptr {
			return fmt.Errorf("decode into non-pointer %T", dest)
		}

		re := rt.Elem()

		if re.Kind() == reflect.Map {
			return decodeMap(value, dest)
		}

		return decodeStruct(value, dest)
	}
}

func (value *Scope) Copy() *Scope {
	copied := NewEmptyScope()
	_ = value.Each(func(k Symbol, v Value) error {
		copied.Set(k, v)
		return nil
	})

	return copied
}

// Reduce calls f for each binding-value pair mapped by the scope.
//
// Note that shadowed bindings will be skipped.
func (value *Scope) Each(f func(Symbol, Value) error) error {
	return value.eachShadow(value, f, map[Symbol]bool{})
}

func (value *Scope) eachShadow(top *Scope, f func(Symbol, Value) error, called map[Symbol]bool) error {
	for _, p := range value.Parents {
		err := p.eachShadow(top, f, called)
		if err != nil {
			return err
		}
	}

	for _, k := range value.Order {
		if called[k] {
			continue
		}

		called[k] = true

		v, found := top.Get(k)
		if !found {
			// TODO: this should be impossible, since the value is present here, but
			// someday we might want copy-on-write remove semantics. i think this
			// could be handled by .Get - e.g. if the value is _ (ignore), pretend
			// it's not there
			continue
		}

		err := f(k, v)
		if err != nil {
			return fmt.Errorf("scope each: %s: %w", k, err)
		}
	}

	return nil
}

func (value *Scope) MarshalJSON() ([]byte, error) {
	buf := new(bytes.Buffer)

	buf.WriteString("{")

	first := true
	err := value.Each(func(k Symbol, v Value) error {
		if !first {
			buf.WriteString(",")
		} else {
			first = false
		}

		key, err := json.Marshal(k.JSONKey())
		if err != nil {
			return err
		}

		val, err := MarshalJSON(v)
		if err != nil {
			return err
		}

		buf.Write(key)
		buf.WriteString(":")
		buf.Write(val)

		return nil
	})
	if err != nil {
		return nil, err
	}

	buf.WriteString("}")

	return buf.Bytes(), nil
}

func (value *Scope) UnmarshalJSON(payload []byte) error {
	return UnmarshalJSON(payload, value)
}

// Eval returns the value.
func (value *Scope) Eval(_ context.Context, _ *Scope, cont Cont) ReadyCont {
	return cont.Call(value, nil)
}

// Set assigns the value in the local bindings.
func (scope *Scope) Set(binding Symbol, value Value, docs ...string) {
	if len(docs) > 0 {
		value = annotate(value, docs...)
	}

	_, found := scope.Bindings[binding]
	if !found {
		scope.Order = append(scope.Order, binding)
	}

	scope.Bindings[binding] = value
}

// Get fetches the given binding.
//
// If a value is set in the local bindings, it is returned.
//
// If not, the parent scopes are queried in order.
//
// If no value is found, false is returned.
func (scope *Scope) Get(binding Symbol) (Value, bool) {
	val, found := scope.Bindings[binding]
	if found {
		return val, found
	}

	for _, parent := range scope.Parents {
		val, found = parent.Get(binding)
		if found {
			return val, found
		}
	}

	return nil, false
}

// GetDecode fetches the given binding and Decodes its value.
func (scope *Scope) GetDecode(binding Symbol, dest any) error {
	val, found := scope.Get(binding)
	if !found {
		return UnboundError{binding, scope}
	}

	return val.Decode(dest)
}

// Binds returns true if the scope provides the given binding.
func (scope *Scope) Binds(binding Symbol) bool {
	_, found := scope.Get(binding)
	return found
}

// Complete queries the scope for bindings beginning with the given prefix.
//
// Local bindings are listed before parent bindings, with shorter binding names
// listed first.
func (scope *Scope) Complete(prefix string) []CompleteOpt {
	shadowed := map[Symbol]bool{}

	var opts []CompleteOpt
	for name, val := range scope.Bindings {
		if strings.HasPrefix(name.String(), prefix) {
			var annotated Annotated
			if err := val.Decode(&annotated); err != nil {
				annotated = Annotated{
					Value: val,
					Meta:  NewEmptyScope(),
				}
			}

			opts = append(opts, CompleteOpt{
				Binding: name,
				Value:   annotated,
			})

			shadowed[name] = true
		}
	}

	sort.Sort(CompleteOpts(opts))

	for _, parent := range scope.Parents {
		for _, opt := range parent.Complete(prefix) {
			if shadowed[opt.Binding] {
				continue
			}

			opts = append(opts, opt)
			shadowed[opt.Binding] = true
		}
	}

	return opts
}

var _ zapcore.ObjectMarshaler = (*Scope)(nil)

func (scope *Scope) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	return scope.Each(func(k Symbol, v Value) error {
		key := k.String()

		var str string
		var num int
		var bol bool
		var am zapcore.ArrayMarshaler
		var om zapcore.ObjectMarshaler
		if v.Decode(&str) == nil {
			enc.AddString(key, str)
		} else if v.Decode(&num) == nil {
			enc.AddInt(key, num)
		} else if v.Decode(&bol) == nil {
			enc.AddBool(key, bol)
		} else if v.Decode(&am) == nil {
			enc.AddArray(key, am)
		} else if v.Decode(&om) == nil {
			enc.AddObject(key, om)
		} else {
			enc.AddString(key, v.String())
		}

		return nil
	})
}

type CompleteOpt struct {
	Binding Symbol
	Value   Annotated
}

type CompleteOpts []CompleteOpt

func (opts CompleteOpts) Len() int      { return len(opts) }
func (opts CompleteOpts) Swap(i, j int) { opts[i], opts[j] = opts[j], opts[i] }
func (opts CompleteOpts) Less(i, j int) bool {
	if len(opts[i].Binding) < len(opts[j].Binding) {
		return true
	}

	if len(opts[i].Binding) > len(opts[j].Binding) {
		return false
	}

	return opts[i].Binding < opts[j].Binding
}

var (
	// DocMetaBinding is the binding in meta that stores the value's doc string.
	DocMetaBinding Symbol = "doc"

	// FileMetaBinding is the binding in meta that stores the value's definition file.
	FileMetaBinding Symbol = "file"

	// LineMetaBinding is the binding in meta that stores the value's definition line.
	LineMetaBinding Symbol = "line"

	// ColumnMetaBinding is the binding in meta that stores the value's definition column.
	ColumnMetaBinding Symbol = "column"

	// DeprecatedMetaBinding is the binding in meta that indicates that the value is deprecated.
	//
	// It should define a string message indicating the alternative to use.
	DeprecatedMetaBinding Symbol = "deprecated"
)

func annotate(val Value, docs ...string) Annotated {
	meta := Bindings{
		DocMetaBinding: String(strings.Join(docs, "\n\n")),
	}

	_, here, _, ok := runtime.Caller(0)
	if ok {
		pkgDir := path.Dir(path.Dir(here))

		_, file, line, ok := runtime.Caller(2)
		if ok {
			if strings.HasPrefix(file, pkgDir) {
				// use embedded filesystem for embedded source
				sub := strings.TrimPrefix(strings.TrimPrefix(file, pkgDir), "/")
				meta[FileMetaBinding] = NewFSPath(pkg.FS, ParseFileOrDirPath(sub))
			} else {
				meta[FileMetaBinding] = ParseHostPath(file)
			}

			meta[LineMetaBinding] = Int(line)
			meta[ColumnMetaBinding] = Int(0)
		}
	}

	var ann Annotated
	if err := val.Decode(&ann); err == nil {
		ann.Meta = NewEmptyScope(ann.Meta, meta.Scope())
		return ann
	}

	return Annotated{
		Value: val,
		Meta:  meta.Scope(),
	}
}

func isOptional(segs []string) bool {
	for _, seg := range segs {
		if seg == "omitempty" {
			return true
		}
	}
	return false
}

func decodeStruct(value *Scope, dest any) error {
	rt := reflect.TypeOf(dest)
	if rt.Kind() != reflect.Ptr {
		return fmt.Errorf("decode into non-pointer %T", dest)
	}

	re := rt.Elem()
	rv := reflect.ValueOf(dest).Elem()

	if re.Kind() != reflect.Struct {
		return fmt.Errorf("%s != %s: %w", re.Kind(), reflect.Struct, DecodeError{
			Source:      value,
			Destination: dest,
		})
	}

	for i := 0; i < re.NumField(); i++ {
		field := re.Field(i)

		tag := field.Tag.Get("json")
		segs := strings.Split(tag, ",")
		name := segs[0]
		if name == "" {
			continue
		}

		sym := SymbolFromJSONKey(name)

		var found bool
		val, found := value.Get(sym)
		if !found {
			if isOptional(segs) {
				continue
			}

			return fmt.Errorf("missing key %s", sym)
		}

		if rv.Field(i).Kind() == reflect.Ptr {
			x := reflect.New(field.Type.Elem())

			err := val.Decode(x.Interface())
			if err != nil {
				return fmt.Errorf("decode (%T).%s: %w", dest, field.Name, err)
			}

			rv.Field(i).Set(x)
		} else {
			err := val.Decode(rv.Field(i).Addr().Interface())
			if err != nil {
				return fmt.Errorf("decode (%T).%s: %w", dest, field.Name, err)
			}
		}
	}

	return nil
}

func decodeMap(value *Scope, dest any) error {
	rt := reflect.TypeOf(dest)
	if rt.Kind() != reflect.Ptr {
		return fmt.Errorf("decode into non-pointer %T", dest)
	}

	re := rt.Elem()
	rv := reflect.ValueOf(dest).Elem()

	if re.Kind() != reflect.Map {
		return fmt.Errorf("%s != %s: %w", re.Kind(), reflect.Map, DecodeError{
			Source:      value,
			Destination: dest,
		})
	}

	if re.Key() != reflect.TypeOf("") {
		return fmt.Errorf("map key type: %s != %s: %w", re.Key(), reflect.String, DecodeError{
			Source:      value,
			Destination: dest,
		})
	}

	rv.Set(reflect.MakeMapWithSize(re, len(value.Bindings)))

	return value.Each(func(key Symbol, val Value) error {
		x := reflect.New(re.Elem())

		err := val.Decode(x.Interface())
		if err != nil {
			return fmt.Errorf("decode map value: %w", err)
		}

		rv.SetMapIndex(reflect.ValueOf(key.String()), x.Elem())

		return nil
	})
}
