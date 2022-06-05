package yahooFinace

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const BASE_URL_ = "https://query2.finance.yahoo.com"
const SCRAPE_URL = "https://finance.yahoo.com/quote"

// http headers
var headers = map[string][]string{
	"Host":            {"finance.yahoo.com"},
	"User-Agent":      {"Mozilla/5.0 (X11; Linux x86_64; rv:101.0) Gecko/20100101 Firefox/101.0"},
	"Accept":          {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8"},
	"Accept-Language": {"en-US,en;q=0.5"},
	"Cookie":          {"A1=d=AQABBHj_mmICEP7EIaqTykPm4A_oWdBH08UFEgEBAQFQnGKkYgAAAAAA_eMAAA&S=AQAAAlOSWal31ku_ETwq3mFkSGU; A3=d=AQABBHj_mmICEP7EIaqTykPm4A_oWdBH08UFEgEBAQFQnGKkYgAAAAAA_eMAAA&S=AQAAAlOSWal31ku_ETwq3mFkSGU; GUC=AQEBAQFinFBipEIhbwTf; maex=%7B%22v2%22%3A%7B%7D%7D; PRF=t%3DTANLA.NS%252BBTC-USD; A1S=d=AQABBHj_mmICEP7EIaqTykPm4A_oWdBH08UFEgEBAQFQnGKkYgAAAAAA_eMAAA&S=AQAAAlOSWal31ku_ETwq3mFkSGU&j=WORLD"},
}

func String2Float(s string) (f float64, err error) {
	const reg = `[-]?\d[\d,]*[\.]?[\d]*`
	re := regexp.MustCompile(reg)

	matchs := re.FindAllString(s, -1)

	if len(matchs) == 0 {
		err = fmt.Errorf("no match found")
		return
	}

	// remove comas
	matchs[0] = strings.ReplaceAll(matchs[0], ",", "")

	f, err = strconv.ParseFloat(string(matchs[0]), 8)
	return
}

// FetchDoc returns goquery.Selection which can then be used for extracting data
// It take a URL and Css selector.
func FetchDoc(URL, selector string) (*goquery.Selection, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, fmt.Errorf("NewRequest: %w", err)
	}

	// add headers
	req.Header = headers
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client Request: %w", err)
	}
	// TODO status code not working for wrong ticker symbol
	// On web, yahoo servers redirect to lookup from quote when searching for wrong symbol
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code:%d on URL:%s", res.StatusCode, URL)
	}

	defer res.Body.Close()

	// create a new document from the response
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("GoQuery: %w", err)
	}

	// find element in the document
	elm := doc.Find(selector)

	// Check the no. elements
	if elm.Size() == 0 {
		return nil, fmt.Errorf(
			"Doc.Find: No element found with the selector: %s",
			selector,
		)
	}
	return elm, nil
}
