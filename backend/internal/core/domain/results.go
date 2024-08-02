package domain

type Result struct {
	TurnNum  string
	OpenNum  string
	OpenTime string
	Detail   string
	Region   string
}

func NewResult(arg Result) *Result {
	return &Result{
		TurnNum:  arg.TurnNum,
		OpenNum:  arg.OpenNum,
		OpenTime: arg.OpenTime,
		Detail:   arg.Detail,
		Region:   arg.Region,
	}
}
