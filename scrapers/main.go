package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/url"

	"github.com/knq/chromedp"
	"github.com/leonb/brewfun"
)

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

	// retrieve javascript
	byt, err := ioutil.ReadFile("hop-substitutes-brew365.js")
	if err != nil {
		log.Fatal(err)
	}
	expr := string(byt)

	url, err := url.Parse("http://www.brew365.com/hop_substitution_chart.php")
	if err != nil {
		log.Fatal(err)
	}

	chart, err := retrieveHopSubstitutes(ctxt, c, *url, expr)
	if err != nil {
		log.Fatal(err)
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

	log.Printf("%+v", chart)
}

func retrieveHopSubstitutes(ctxt context.Context, client *chromedp.CDP, url url.URL, expression string) (brewfun.HopSubstitutionChart, error) {
	// run expression
	var res brewfun.HopSubstitutionChart
	err := client.Run(ctxt, chromedp.Tasks{
		chromedp.Navigate("http://www.brew365.com/hop_substitution_chart.php"),
		chromedp.WaitVisible(`table tr`, chromedp.ByQuery),
		chromedp.Evaluate(expression, &res),
	})
	return res, err
}
