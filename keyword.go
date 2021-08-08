package bass

import (
	"context"
	"fmt"
	"strings"
)

type Keyword string

var _ Value = Keyword("")

func hyphenate(value Keyword) string {
	return strings.ReplaceAll(string(value), "_", "-")
}

func (value Keyword) String() string {
	return fmt.Sprintf(":%s", hyphenate(value))
}

func (value Keyword) Equal(other Value) bool {
	var o Keyword
	return other.Decode(&o) == nil && value == o
}

func (value Keyword) Decode(dest interface{}) error {
	switch x := dest.(type) {
	case *Keyword:
		*x = value
		return nil
	case *Applicative:
		*x = value
		return nil
	case *Combiner:
		*x = value
		return nil
	case *Bindable:
		*x = value
		return nil
	case *Value:
		*x = value
		return nil
	default:
		return DecodeError{
			Source:      value,
			Destination: dest,
		}
	}
}

// Eval returns the value.
func (value Keyword) Eval(ctx context.Context, env *Env, cont Cont) ReadyCont {
	return cont.Call(value, nil)
}

var _ Applicative = Keyword("")

func (app Keyword) Unwrap() Combiner {
	return KeywordOperative{app}
}

var _ Combiner = Keyword("")

func (combiner Keyword) Call(ctx context.Context, val Value, env *Env, cont Cont) ReadyCont {
	return Wrapped{KeywordOperative{combiner}}.Call(ctx, val, env, cont)
}

type KeywordOperative struct {
	Keyword Keyword
}

var _ Value = KeywordOperative{}

func (value KeywordOperative) String() string {
	return fmt.Sprintf("(unwrap %s)", value.Keyword)
}

func (value KeywordOperative) Equal(other Value) bool {
	var o KeywordOperative
	return other.Decode(&o) == nil && value == o
}

func (value KeywordOperative) Decode(dest interface{}) error {
	switch x := dest.(type) {
	case *KeywordOperative:
		*x = value
		return nil
	case *Combiner:
		*x = value
		return nil
	case *Value:
		*x = value
		return nil
	default:
		return DecodeError{
			Source:      value,
			Destination: dest,
		}
	}
}

func (value KeywordOperative) Eval(ctx context.Context, env *Env, cont Cont) ReadyCont {
	return cont.Call(value, nil)
}

func (op KeywordOperative) Call(ctx context.Context, val Value, env *Env, cont Cont) ReadyCont {
	var list List
	err := val.Decode(&list)
	if err != nil {
		return cont.Call(nil, fmt.Errorf("call keyword: %w", err))
	}

	var obj Object
	err = list.First().Decode(&obj)
	if err != nil {
		return cont.Call(nil, err)
	}

	val, found := obj[op.Keyword]
	if found {
		return cont.Call(val, nil)
	}

	var rest List
	err = list.Rest().Decode(&rest)
	if err != nil {
		return cont.Call(nil, err)
	}

	var empty Empty
	err = rest.Decode(&empty)
	if err == nil {
		return cont.Call(Null{}, nil)
	}

	return cont.Call(rest.First(), nil)
}

var _ Bindable = Keyword("")

func (binding Keyword) Bind(env *Env, val Value) error {
	return BindConst(binding, val)
}
