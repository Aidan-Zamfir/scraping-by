package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
)

type Rental struct {
	Neighborhood string `json:"neighborhood"`
	Price        string `json:"price"`
	PriceSQM     string `json:"price_sqm"`
	Kitchens     string `json:"kitchens"`
	LivingR      string `json:"living_rooms"`
	Bathrooms    string `json:"bathrooms"`
	LastUpdated  string `json:"last_updated"`
	PropertyCode string `json:"property_code"`
	Link         string `json:"link"`
}

func ScrapeData() {
	c := colly.NewCollector()

	c.OnHTML("div[class=property]", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)

	})
	c.Visit("https://www.indomio.al/en/to-rent/property/tirana-city")

}
