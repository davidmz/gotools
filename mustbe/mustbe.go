package mustbe

import "reflect"

type errBag struct{ err error }

func OK(err error) {
	if err != nil {
		panic(errBag{err: err})
	}
}

var Throw = OK

func Catched(err *error, nullify ...interface{}) {
	switch e := recover().(type) {
	case nil:
		// pass
	case errBag:
		*err = e.err
		for _, n := range nullify {
			v := reflect.ValueOf(n)
			ind := reflect.Indirect(v)
			if v.Kind() == reflect.Ptr && ind.CanSet() {
				ind.Set(reflect.Zero(ind.Type()))
			}
		}
	default:
		panic(e)
	}
}

type elser bool

func (e elser) Else(err error) {
	if !bool(e) {
		Throw(err)
	}
}

func True(test bool) elser { return elser(test) }
