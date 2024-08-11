package domain

const (
	RankSpecial = 0
	RankFirst   = 1
	RankSecond  = 2
	RankForth   = 3
	RankFifth   = 4
	RankSixth   = 5
	RankSeventh = 6
	RankEighth  = 7
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
}

func NewOpenNum(arg OpenNum) *OpenNum {
	return &OpenNum{
		ID:       arg.ID,
		ResultID: arg.ResultID,
		Numbs:    arg.Numbs,
		Rank:     arg.Rank,
	}
}
