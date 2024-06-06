package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Event struct {
	Date string
	Text string
	Year string
}

func main() {
	c := colly.NewCollector()

	file, err := os.Create("records.csv")
	var events []Event
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(file)
	defer w.Flush()

	c.OnHTML("table.wikitable > tbody", func(h *colly.HTMLElement) {
		// fmt.Printf(h.Text)
		h.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			event := Event{
				el.ChildText("td:nth-child(1)"),
				el.ChildText("td:nth-child(3)"),
				el.ChildText("td:nth-child(2)"),
			}
			events = append(events, event)
		})
	})

	c.Visit("https://en.wikipedia.org/wiki/Data_Encryption_Standard")
	events = events[1:]

	for _, event := range events {
		row := []string{event.Date, event.Year, event.Text}
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
	fmt.Println("Finished writing to a csv records.csv")
}
