package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/knq/chromedp"
	"github.com/leonb/brewfun"
)

type scraper struct {
	jsFile   string
	url      string
	selector string
	match    float64
}

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	scrapers := []scraper{
		scraper{
			jsFile:   "hop-substitutes-brew365.js",
			url:      "http://www.brew365.com/hop_substitution_chart.php",
			selector: "table tr",
			match:    0.8,
		},
		// scraper{
		// 	jsFile:   "hop-substitutes-aha.js",
		// 	url:      "https://www.homebrewersassociation.org/how-to-brew/hop-substitutions/",
		// 	selector: "table tr",
		// 	match:    0.5,
		// },
		// scraper{
		// 	jsFile:   "hop-substitutes-homebrewstuff.js",
		// 	url:      "http://www.homebrewstuff.com/hop-profiles",
		// 	selector: ".std",
		// 	match:    0.5,
		// },
	}

	substitutes := brewfun.HopSubstitutes{}
	for _, scraper := range scrapers {
		subs, err := retrieveHopSubstitutes(ctxt, c, scraper)
		if err != nil {
			log.Fatal(err)
		}

		substitutes = append(substitutes, subs...)
	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}

	// @TODO: aggregate results?
	// @TODO: remove double entries?

	b, _ := json.MarshalIndent(substitutes, "", "   ")
	log.Println(string(b))
}

func retrieveHopSubstitutes(ctxt context.Context, client *chromedp.CDP, scraper scraper) (brewfun.HopSubstitutes, error) {
	// run expression
	substitutes := brewfun.HopSubstitutes{}

	// retrieve javascript
	byt, err := ioutil.ReadFile(scraper.jsFile)
	if err != nil {
		return substitutes, err
	}
	expr := string(byt)

	err = client.Run(ctxt, chromedp.Tasks{
		chromedp.Navigate(scraper.url),
		chromedp.WaitReady(scraper.selector, chromedp.ByQuery),
		chromedp.Evaluate(fmt.Sprintf("var source = '%s';", scraper.url), &[]byte{}),
		chromedp.Evaluate(fmt.Sprintf("var selector = '%s';", scraper.selector), &[]byte{}),
		chromedp.Evaluate(fmt.Sprintf("var defaultMatch = %f;", scraper.match), &[]byte{}),
		chromedp.Evaluate(expr, &substitutes),
	})
	return substitutes, err
}
