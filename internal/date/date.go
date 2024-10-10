package date

import (
	"fmt"
	"time"

	"github.com/o98k-ok/lazy/v2/alfred"
)

type Date struct {
	Date      string
	DayOfWeek string
	IconPath  string
}

func (d Date) ToAlfredElem() *alfred.Item {
	return &alfred.Item{
		Title:    d.Date,
		SubTitle: d.DayOfWeek,
		Icon:     &alfred.Icon{Path: d.IconPath},
	}
}

func iconPath(date time.Time, mode string) string {
	_, _, day := date.Date()
	return fmt.Sprintf("icon/date_%d_%s.png", day, mode)
}

func currentMode(date time.Time, now time.Time) string {
	if date.Month() != now.Month() {
		return "white"
	}
	if date.Day() == time.Now().Day() {
		return "color"
	}
	return "black"
}

func GetDates() []Date {
	now := time.Now()
	dow := now.Weekday()
	startDate := now.AddDate(0, 0, -int(dow)-7)

	var dates []Date
	for i := 0; i < 35; i++ {
		date := startDate.AddDate(0, 0, i)
		dates = append(dates, Date{
			Date:      date.Format("2006-01-02"),
			DayOfWeek: date.Weekday().String(),
			IconPath:  iconPath(date, currentMode(date, now)),
		})
	}
	return dates
}
