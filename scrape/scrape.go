package scrape

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"

	// "github.com/kr/pretty"
	"github.com/pdunnavant/modem-scraper/config"
)

// Scrape scrapes data from the modem.
func Scrape(config config.Configuration) (*ModemInformation, error) {
	doc, err := getDocumentFromURL(config.IP + "/cmconnectionstatus.html")
	if err != nil {
		return nil, err
	}
	connectionStatus := scrapeConnectionStatus(doc)

	doc, err = getDocumentFromURL(config.IP + "/cmswinfo.html")
	if err != nil {
		return nil, err
	}
	softwareInformation := scrapeSoftwareInformation(doc)

	doc, err = getDocumentFromURL(config.IP + "/cmeventlog.html")
	if err != nil {
		return nil, err
	}
	eventLog := scrapeEventLogs(doc)

	modemInformation := ModemInformation{
		ConnectionStatus:    *connectionStatus,
		SoftwareInformation: *softwareInformation,
		EventLog:            eventLog,
	}
	// fmt.Printf("%# v", pretty.Formatter(modemInformation))

	return &modemInformation, nil
}

func getDocumentFromURL(url string) (*goquery.Document, error) {
	fmt.Printf("Grabbing [%s]...\n", url)
	defer timeTrack(time.Now(), fmt.Sprintf("Got [%s].", url))

	// TODO: add a timeout here (10s?)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s (Took %s.)\n", name, elapsed)
}
