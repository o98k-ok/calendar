package main

import (
	"fmt"

	"github.com/o98k-ok/calendar/internal/date"
	"github.com/o98k-ok/lazy/v2/alfred"
)

func main() {
	dates := date.GetDates()
	items := alfred.NewItems()
	for _, date := range dates {
		items.Append(date.ToAlfredElem())
	}
	fmt.Println(items.Encode())
}
