package util

import "time"

var (
	DateFormat = "2006-01-02 15:04:05"
)

func GetToDay() string {
	return time.Now().Format(DateFormat)
}

func GetToDayAsDatabaseTime() string {
	return time.Now().Format(DateFormat)
}

func ParseToFormattedDate(date string) (time.Time, error) {
	loc, _ := time.LoadLocation("Asia/Saigon")
	return time.ParseInLocation(DateFormat, date, loc)
}
