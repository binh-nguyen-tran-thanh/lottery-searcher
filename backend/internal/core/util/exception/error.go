package exception

const (
	TypeInternal = "InternalError"
	TypeNotFound = "NotFoundError"
)

type Err = map[string][]string

type Exception struct {
	Message string `json:"-"`
	Type    string `json:"-"`
	Errors  Err    `json:"errors"`
	Cause   error  `json:"-"`
}

func New(kind, message string, cause error) *Exception {
	return &Exception{
		Type:    kind,
		Message: message,
		Cause:   cause,
		Errors:  make(map[string][]string),
	}
}

func (e Exception) HasError() bool {
	return len(e.Errors) > 0
}

func (e *Exception) AddError(key, message string) *Exception {
	e.Errors[key] = append(e.Errors[key], message)
	return e
}

func Into(err error) *Exception {
	if err == nil {
		return nil
	}

	fail, ok := err.(*Exception)

	if ok {
		return fail
	}

	return New(TypeInternal, err.Error(), err)
}

func (e *Exception) Error() string {
	if e.Cause == nil {
		return ""
	}

	return e.Cause.Error()
}
