package mustbe

import "fmt"

type Error struct {
	Base  error
	Class *ErrorClass
}

func (e *Error) Error() (s string) {
	s = e.Class.FullTitle()
	if s != "" {
		s += ": "
	}
	s += e.Base.Error()
	return
}

type ErrorClass struct {
	Title     string
	BaseClass *ErrorClass
}

func NewErrorClass(title string) *ErrorClass { return &ErrorClass{Title: title} }

func (c *ErrorClass) Subclass(title string) *ErrorClass {
	return &ErrorClass{BaseClass: c, Title: title}
}

func (c *ErrorClass) FullTitle() (s string) {
	if c.BaseClass != nil {
		s = c.BaseClass.FullTitle()
	}
	if s != "" {
		s += ": "
	}
	s += c.Title
	return
}

func (c *ErrorClass) Wrap(err error) error {
	if err != nil {
		return &Error{Base: err, Class: c}
	}
	return nil
}

func (c *ErrorClass) Error(format string, args ...interface{}) error {
	return c.Wrap(fmt.Errorf(format, args...))
}

// 'IsA' возвращает true, если 'с' или один из его родительских классов имеет класс 'cRef'
func (c *ErrorClass) IsA(cRef *ErrorClass) bool {
	return c != nil && (c == cRef || c.BaseClass.IsA(cRef))
}

// 'Contains' возвращает true, если 'err' имеет класс 'c' или один из его родительских классов
func (c *ErrorClass) Contains(err error) (ok bool) {
	var e *Error
	if e, ok = err.(*Error); ok {
		ok = e.Class.IsA(c)
	}
	return
}
