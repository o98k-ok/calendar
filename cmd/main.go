package main

import (
	"fmt"
	"os"

	"github.com/o98k-ok/calendar/internal/date"
	"github.com/o98k-ok/lazy/v2/alfred"
)

func main() {
	app := alfred.NewApp("日历助手")
	app.Bind("all", func(s []string) {
		dates := date.GetDates()
		items := alfred.NewItems()
		for _, date := range dates {
			items.Append(date.ToAlfredElem())
		}
		fmt.Println(items.Encode())
	})

	app.Bind("detail", func(s []string) {
		detail := date.Detail(s[0])
		fmt.Println(detail.DetailFilter().Encode())
	})

	app.Run(os.Args)
}
