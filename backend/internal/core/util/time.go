package util

import "time"

var (
	DateFormat     = "2006-01-02 15:04:05"
	OpenDateFormat = "02/02/2006"
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

func IsBeforeNow(date string, format string) (bool, error) {
	dateLayout := format
	if format == "" {
		dateLayout = OpenDateFormat
	}
	openDate, err := time.Parse(dateLayout, date)

	if err != nil {
		return false, err
	}

	nowDate := time.Now().Truncate(24 * time.Hour)

	return openDate.Before(nowDate), nil
}
