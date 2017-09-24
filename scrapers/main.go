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
		// scraper{
		// 	jsFile:   "hop-substitutes-brew365.js",
		// 	url:      "http://www.brew365.com/hop_substitution_chart.php",
		// 	selector: "table tr",
		// },
		// scraper{
		// 	jsFile:   "hop-substitutes-aha.js",
		// 	url:      "https://www.homebrewersassociation.org/how-to-brew/hop-substitutions/",
		// 	selector: "table tr",
		// },
		scraper{
			jsFile:   "hop-substitutes-homebrewstuff.js",
			url:      "http://www.homebrewstuff.com/hop-profiles",
			selector: ".std",
		},
	}

	charts := []brewfun.HopSubstitutionChart{}
	for _, scraper := range scrapers {
		chart, err := retrieveHopSubstitutes(ctxt, c, scraper)
		if err != nil {
			log.Fatal(err)
		}

		charts = append(charts, chart)
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

	b, _ := json.MarshalIndent(charts, "", "   ")
	log.Println(string(b))
}

func retrieveHopSubstitutes(ctxt context.Context, client *chromedp.CDP, scraper scraper) (brewfun.HopSubstitutionChart, error) {
	// run expression
	var res brewfun.HopSubstitutionChart

	// retrieve javascript
	byt, err := ioutil.ReadFile(scraper.jsFile)
	if err != nil {
		return res, err
	}
	expr := string(byt)

	err = client.Run(ctxt, chromedp.Tasks{
		chromedp.Navigate(scraper.url),
		chromedp.WaitReady(scraper.selector, chromedp.ByQuery),
		chromedp.Evaluate(fmt.Sprintf("var source = '%s';", scraper.url), &[]byte{}),
		chromedp.Evaluate(fmt.Sprintf("var selector = '%s';", scraper.selector), &[]byte{}),
		chromedp.Evaluate(expr, &res),
	})
	return res, err
}
