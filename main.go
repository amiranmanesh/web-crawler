package main

import (
	"fmt"
	"github.com/amiranmanesh/people-crawler/database"
	colly "github.com/gocolly/colly/v2"
	"log"
	"strings"
	"sync"
)

var c *colly.Collector
var wait sync.WaitGroup

func getUrl(page int) string {
	return fmt.Sprintf("https://www.locatefamily.com/Street-Lists/Iran/index-%d.html", page)
}

func main() {
	initMain()
	start := 1
	//start := 390
	end := 2527

	wait.Add(end - start + 1)

	for i := start; i <= end; i++ {
		makeRequest(getUrl(i))
		wait.Done()
	}
	wait.Wait()

}

func initMain() {
	c = colly.NewCollector()
	database.DB.Initialize()
}

func makeRequest(url string) {

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println(fmt.Sprintf("Something went wrong(url:%s): ", url), err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML("div[itemtype=\"https://schema.org/Person\"]", func(e *colly.HTMLElement) {

		person := database.Person{}

		line := e.ChildText("li[class=linenumber]")
		person.Address = e.ChildText("span[itemprop=streetAddress]")

		if add := e.ChildAttr("a[href]", "href"); strings.Contains(add, "https://maps.google.com/") {
			person.AddressGoogle = add
		}
		person.AddressLocality = e.ChildText("span[itemprop=addressLocality]")
		person.Phone = e.ChildText("a[class=phone-link]")
		person.FName = e.ChildText("span[itemprop=givenName]")
		person.LName = e.ChildText("span[itemprop=familyName]")

		err := person.Save()
		if err != nil {
			fmt.Println(fmt.Sprintf("Error in save person at line: %s at url: %s", line, url))
			println("fuck here")
		}

	})

	_ = c.Visit(url)

	c.Wait()
}
