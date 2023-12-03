package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

// func getFollowers(url string) {
// 	c := colly.NewCollector()

// 	url = "https://www.instagram.com/al__pha____jt/followers/"

// 	firstElement := "9f619 xjbqb8w x78zum5 x168nmei x13lgxp2 x5pf9jr xo71vjh x1n2onr6 x1plvlek xryxfnj x1c4vz4f x2lah0s x1q0g3np xqjyukv x6s0dn4 x1oa3qoh x1nhvcw1"

// 	// Set up callbacks to handle scraping events
// 	c.OnHTML(firstElement, func(e *colly.HTMLElement) {

// 		followers := e.ChildTexts("span._ap3a _aaco _aacw _aacx _aad7 _aade")

// 		fmt.Printf("Followers: %s\n", followers)
// 	})

// 	// Visit the URL and start scraping
// 	err := c.Visit(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.instagram.com", "www.facebook.com", "www.settrade.com"),
	)

	url := "https://www.instagram.com/al__pha____jt"
	// url := "https://www.settrade.com/th/equities/quote/COM7/overview"

	c.OnHTML("body", func(e *colly.HTMLElement) {
		fmt.Println(e)
	})

	// c.OnHTML("div.row div.col-12.col-md-6.d-md-block", func(e *colly.HTMLElement) {
	// 	// selection := e.DOM
	// 	fmt.Println("==================>", e)
	// 	// fmt.Println(selection.Find("div.justify-content-lg-start"))
	// })

	// Visit the URL and start scraping
	err := c.Visit(url)
	if err != nil {
		log.Println("error: and errro is ====>:", err)
		log.Fatal(err)
	}

}
