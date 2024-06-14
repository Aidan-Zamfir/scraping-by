package scraper

import (
	"github.com/gocolly/colly"
	"log"
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

//for i in site list -> go routine for each site scraping data

func ScrapeData() {
	log.Println("scraping...")

	scrapeURL := "https://www.idealista.com/en/alquiler-viviendas/malaga-malaga/"

	c := colly.NewCollector()

	c.OnHTML("div.item-info-container ", func(e *colly.HTMLElement) {

		log.Println(e.ChildText("a.title"))
		//apartment := Rental{
		//	Neighborhood:
		//	Price: e.ChildText("div.property-card-bottom-price"),
		//	PriceSQM:
		//	Kitchens:
		//	LivingR:
		//	Bathrooms:
		//	LastUpdated:
		//	PropertyCode:
		//	Link:
		//}
	})

	c.Visit(scrapeURL)

}
