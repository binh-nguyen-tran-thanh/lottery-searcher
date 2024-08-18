package domain

const (
	RankSpecial          = 0
	RankFirst            = 1
	RankSecond           = 2
	RankThird            = 3
	RankForth            = 4
	RankFifth            = 5
	RankSixth            = 6
	RankSeventh          = 7
	RankEighth           = 8
	SearchModeFirstThree = "firstThree"
	SearchModeLastThree  = "lastThree"
	SearchModeAll        = "all"
)

type Result struct {
	TurnNum  string
	OpenNum  string
	OpenTime string
	Region   string
	ID       uint
	Detail   string
}

func NewResult(arg Result) *Result {
	return &Result{
		TurnNum:  arg.TurnNum,
		OpenNum:  arg.OpenNum,
		OpenTime: arg.OpenTime,
		Region:   arg.Region,
		ID:       arg.ID,
		Detail:   arg.Detail,
	}
}

type OpenNum struct {
	ID       uint
	ResultID uint
	Numbs    string
	Rank     int8
	Result   Result
}

func NewOpenNum(arg OpenNum) *OpenNum {
	return &OpenNum{
		ID:       arg.ID,
		ResultID: arg.ResultID,
		Numbs:    arg.Numbs,
		Rank:     arg.Rank,
		Result:   arg.Result,
	}
}
