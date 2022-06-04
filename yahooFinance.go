package yahooFinace

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const BASE_URL_ = "https://query2.finance.yahoo.com"
const SCRAPE_URL = "https://finance.yahoo.com/quote"

var HEADER = map[string][]string{
	"Host":            {"finance.yahoo.com"},
	"User-Agent":      {"Mozilla/5.0 (X11; Linux x86_64; rv:101.0) Gecko/20100101 Firefox/101.0"},
	"Accept":          {"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8"},
	"Accept-Language": {"en-US,en;q=0.5"},
}

const version = "0.0.1"
const authorName = "Vivek Rawat"
const program_name = "yahooFinance"

func printProgramDetails() {
	fmt.Printf(
		"Program Name: %s\nAuthor Name: %s\nVersion: %s\n",
		program_name,
		authorName,
		version,
	)
}

type CurrentMarket struct {
	Price            float64
	Change           float64
	ChangePercentage float64
}
type Ticker struct {
	symbol        string
	CurrentMarket CurrentMarket
}

func NewTicker(symbol string) (T Ticker) {
	T = Ticker{symbol, CurrentMarket{}}
	return T
}

func (T *Ticker) GetPrice() (cm CurrentMarket, err error) {
	URL := fmt.Sprintf("%s/%s", SCRAPE_URL, T.symbol)
	fmt.Printf("Calling URL: %s\n", URL)

	doc, err := FetchDoc(URL)
	if err != nil {
		err = fmt.Errorf("GoQuery: %w", err)
		return
	}
	// find the element
	selector := "#quote-header-info fin-streamer"
	q := doc.Find(selector)

	// Chick the no. elements
	if q.Size() == 0 {
		err = fmt.Errorf("Doc.Find: No element found with the selector: %s", selector)
		return
	}

	var f string // filed
	var e bool   // exist
	var v string // value
	PriceSelector := "regularMarketPrice"
	ChangeSelector := "regularMarketChange"
	PercentageSelector := "regularMarketChangePercentage"
	q.Each(func(_ int, s *goquery.Selection) {
		f, e = s.Attr("data-field")
		if !e {
			//throw error for attribute does not exist
		}

		v, e = s.Attr("value")
		if !e {
			//throw error for attribute does not exist
		}

		switch f {
		case PriceSelector:
			fmt.Println(f, v)
			v = strings.ReplaceAll(v, ",", "")
			cm.Price, _ = strconv.ParseFloat(v, 8)
			break
		case ChangeSelector:
			fmt.Println(f, v)
			v = strings.ReplaceAll(v, ",", "")
			cm.Change, _ = strconv.ParseFloat(v, 8)
			break
		case PercentageSelector:
			fmt.Println(f, v)
			v = strings.ReplaceAll(v, ",", "")
			cm.ChangePercentage, _ = strconv.ParseFloat(v, 8)
			break
		default:
			fmt.Println(f, v)
		}
	})
	// error for string to float conversion
	if err != nil {
		err = fmt.Errorf("String to Float: %w", err)
		return
	}
	fmt.Println(cm)
	return
}

func FetchDoc(URL string) (*goquery.Document, error) {
	client := http.Client{}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, fmt.Errorf("NewRequest: %w", err)
	}

	// add headers
	req.Header = HEADER
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Client Request: %w", err)
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Unexpected status code:%d on URL:%s", res.StatusCode, URL)
	}
	defer res.Body.Close()
	// create a new document from the response
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("GoQuery: %w", err)
	}
	return doc, nil
}
