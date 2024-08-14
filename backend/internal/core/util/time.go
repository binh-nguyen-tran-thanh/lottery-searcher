package util

import (
	"time"
)

var (
	DateFormat     = "2006-01-02 15:04:05"
	OpenDateFormat = "02/02/2006"
	TimeZone       = "Asia/Saigon"
	DateOnlyFormat = "2006-01-02"
)

func GetToDay() string {
	return time.Now().Local().Format(DateFormat)
}

func GetToDayDate() string {
	return time.Now().Local().Format(DateOnlyFormat)
}

func GetToDayAsDatabaseTime() string {
	return time.Now().Format(DateFormat)
}

func ToDatabaseFormat(date time.Time) string {
	return date.Format(DateFormat)
}

func ParseToFormattedDate(date string) (time.Time, error) {
	loc, _ := time.LoadLocation(TimeZone)
	return time.ParseInLocation(DateFormat, date, loc)
}

func IsBeforeNow(date string, format string) (bool, error) {
	dateLayout := format
	if format == "" {
		dateLayout = OpenDateFormat
	}
	loc, _ := time.LoadLocation(TimeZone)
	openDate, err := time.ParseInLocation(dateLayout, date, loc)

	if err != nil {
		return false, err
	}
	nowDate := time.Now().Truncate(24 * time.Hour)

	result := openDate.Truncate(24 * time.Hour).Before(nowDate)

	return result, nil
}

func GenerateBeginOfDate() time.Time {
	loc, _ := time.LoadLocation(TimeZone)
	nowTime := time.Now()
	return time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, loc)
}

func GenerateEndOfDate() time.Time {
	loc, _ := time.LoadLocation(TimeZone)
	nowTime := time.Now()
	return time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 23, 59, 59, 999, loc)
}
