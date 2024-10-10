package date

import (
	"fmt"
	"strings"
	"time"

	"github.com/6tail/lunar-go/calendar"
	"github.com/o98k-ok/lazy/v2/alfred"
)

type Date struct {
	Date      string
	DayOfWeek string
	IconPath  string
	Lunar     string
	Jieqi     string
	Festivals string
}

func (d Date) ToAlfredElem() *alfred.Item {
	return &alfred.Item{
		Title: func() string {
			if d.Festivals != "" {
				return fmt.Sprintf("%s %s", d.DayOfWeek, d.Festivals)
			}
			if d.Jieqi != "" {
				return fmt.Sprintf("%s %s", d.DayOfWeek, d.Jieqi)
			}
			return d.DayOfWeek
		}(),
		SubTitle: func() string {
			builder := strings.Builder{}
			builder.WriteString(fmt.Sprintf("%s %s", d.Date, d.Lunar))
			if d.Festivals != "" {
				builder.WriteString(" ")
				builder.WriteString(d.Festivals)
			}
			if d.Jieqi != "" {
				builder.WriteString(" ")
				builder.WriteString(d.Jieqi)
			}
			return builder.String()
		}(),
		Icon: &alfred.Icon{Path: d.IconPath},
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

func ChineseDayOfWeek(date time.Time) string {
	dayOfWeek := date.Weekday()
	switch dayOfWeek {
	case time.Sunday:
		return "周日"
	case time.Monday:
		return "周一"
	case time.Tuesday:
		return "周二"
	case time.Wednesday:
		return "周三"
	case time.Thursday:
		return "周四"
	case time.Friday:
		return "周五"
	case time.Saturday:
		return "周六"
	default:
		return ""
	}
}

func Lunar(date time.Time) string {
	cal := calendar.NewLunarFromDate(date)
	return fmt.Sprintf("%s月-%s", cal.GetMonthInChinese(), cal.GetDayInChinese())
}

func Festivals(date time.Time) string {
	cal := calendar.NewLunarFromDate(date)
	l := cal.GetFestivals()

	var festivals []string
	for i := l.Front(); i != nil; i = i.Next() {
		festivals = append(festivals, i.Value.(string))
	}
	return strings.Join(festivals, " ")
}

func JieQi(date time.Time) string {
	cal := calendar.NewLunarFromDate(date)
	return cal.GetJieQi()
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
			DayOfWeek: ChineseDayOfWeek(date),
			IconPath:  iconPath(date, currentMode(date, now)),
			Lunar:     Lunar(date),
			Festivals: Festivals(date),
			Jieqi:     JieQi(date),
		})
	}
	return dates
}
