package date

import (
	"fmt"
	"strings"
	"time"

	"github.com/6tail/lunar-go/HolidayUtil"
	"github.com/6tail/lunar-go/calendar"
	"github.com/o98k-ok/lazy/v2/alfred"
)

const (
	DateTitle     = "日期"
	LunarTitle    = "农历"
	HolidayTitle  = "假期"
	FestivalTitle = "节日"
	JieqiTitle    = "节气"
	WeekTitle     = "星期"
)

type Date struct {
	Date      string
	DayOfWeek string
	IconPath  string
	Lunar     string
	Jieqi     string
	Festivals string
	Holiday   string
}

func dateKey(mode string) string {
	return fmt.Sprintf("./icon/date_%s.png", mode)
}

func lunarKey(mode string) string {
	return fmt.Sprintf("./icon/calendarfull_%s.png", mode)
}

func weekKey(weekday string, mode string) string {
	switch weekday {
	case "周一":
		return fmt.Sprintf("./icon/monday_%s.png", mode)
	case "周二":
		return fmt.Sprintf("./icon/tuesday_%s.png", mode)
	case "周三":
		return fmt.Sprintf("./icon/wednesday_%s.png", mode)
	case "周四":
		return fmt.Sprintf("./icon/thursday_%s.png", mode)
	case "周五":
		return fmt.Sprintf("./icon/friday_%s.png", mode)
	case "周六":
		return fmt.Sprintf("./icon/saturday_%s.png", mode)
	case "周日":
		return fmt.Sprintf("./icon/sunday_%s.png", mode)
	}
	return ""
}

func holidayKey(holiday string, mode string) string {
	if strings.Contains(holiday, "班") {
		return fmt.Sprintf("./icon/overtime_%s.png", mode)
	}
	return fmt.Sprintf("./icon/holiday_%s.png", mode)
}

func festivalKey(mode string) string {
	return fmt.Sprintf("./icon/temple_%s.png", mode)
}

func jieqiKey(mode string) string {
	return fmt.Sprintf("./icon/blossom_%s.png", mode)
}

func (d Date) DetailFilter() *alfred.Items {
	items := alfred.NewItems()
	items.Append(&alfred.Item{
		Title:    d.Date,
		SubTitle: DateTitle,
		Arg:      d.Date,
		Icon:     &alfred.Icon{Path: dateKey(MODE)},
	})
	items.Append(&alfred.Item{
		Title:    d.Lunar,
		SubTitle: LunarTitle,
		Arg:      d.Lunar,
		Icon:     &alfred.Icon{Path: lunarKey(MODE)},
	})
	items.Append(&alfred.Item{
		Title:    d.DayOfWeek,
		SubTitle: WeekTitle,
		Arg:      d.DayOfWeek,
		Icon:     &alfred.Icon{Path: weekKey(d.DayOfWeek, MODE)},
	})
	if d.Holiday != "" {
		items.Append(&alfred.Item{
			Title:    d.Holiday,
			SubTitle: HolidayTitle,
			Arg:      d.Holiday,
			Icon:     &alfred.Icon{Path: holidayKey(d.Holiday, MODE)},
		})
	}
	if d.Festivals != "" {
		items.Append(&alfred.Item{
			Title:    d.Festivals,
			SubTitle: FestivalTitle,
			Arg:      d.Festivals,
			Icon:     &alfred.Icon{Path: festivalKey(MODE)},
		})
	}
	if d.Jieqi != "" {
		items.Append(&alfred.Item{
			Title:    d.Jieqi,
			SubTitle: JieqiTitle,
			Arg:      d.Jieqi,
			Icon:     &alfred.Icon{Path: jieqiKey(MODE)},
		})
	}
	return items
}

type Items struct {
	Items     []*Item `json:"items"`
	Preselect string  `json:"preselect"`
}

type Item struct {
	alfred.Item
	Uid string `json:"uid"`
}

func (d Date) ToAlfredElem() *Item {
	return &Item{
		Item: alfred.Item{
			Title: func() string {
				if d.Holiday != "" {
					return d.Holiday
				}
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
				builder.WriteString(fmt.Sprintf("%s:%s %s:%s", DateTitle, d.Date, LunarTitle, d.Lunar))
				if d.Holiday != "" {
					builder.WriteString(fmt.Sprintf(" %s:%s", HolidayTitle, d.Holiday))
				}
				if d.Festivals != "" {
					builder.WriteString(fmt.Sprintf(" %s:%s", FestivalTitle, d.Festivals))
				}
				if d.Jieqi != "" {
					builder.WriteString(fmt.Sprintf(" %s:%s", JieqiTitle, d.Jieqi))
				}
				return builder.String()
			}(),
			Arg:  d.Date,
			Icon: &alfred.Icon{Path: d.IconPath},
		},
		Uid: d.Date,
	}
}

func iconPath(date time.Time, mode string) string {
	_, _, day := date.Date()
	return fmt.Sprintf("icon/date_%d_%s.png", day, mode)
}

func currentMode(date time.Time, now time.Time, mode string) string {
	var rmode string
	if mode == "black" {
		rmode = "white"
	} else {
		rmode = "black"
	}

	if date.Month() != now.Month() {
		return rmode
	}
	if date.Day() == time.Now().Day() {
		return "color"
	}
	return mode
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

func Holiday(date time.Time) string {
	holiday := HolidayUtil.GetHoliday(date.Format("2006-01-02"))
	if holiday == nil {
		return ""
	}

	var result strings.Builder
	result.WriteString(holiday.GetName())
	if holiday.IsWork() {
		result.WriteString("-班")
	} else {
		result.WriteString("-休")
	}
	return result.String()
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
			IconPath:  iconPath(date, currentMode(date, now, MODE)),
			Lunar:     Lunar(date),
			Festivals: Festivals(date),
			Jieqi:     JieQi(date),
			Holiday:   Holiday(date),
		})
	}
	return dates
}

func Detail(date string) Date {
	ts, _ := time.Parse("2006-01-02", date)
	return Date{
		Date:      date,
		DayOfWeek: ChineseDayOfWeek(ts),
		Lunar:     Lunar(ts),
		Festivals: Festivals(ts),
		Jieqi:     JieQi(ts),
		Holiday:   Holiday(ts),
	}
}
