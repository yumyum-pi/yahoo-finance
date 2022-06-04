package yahooFinace

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// CSS selectors for Market Price
const MarketPriceSelector = "#quote-header-info fin-streamer"
const PriceSelector = "regularMarketPrice"
const ChangeSelector = "regularMarketChange"
const PercentageSelector = "regularMarketChangePercent"

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

	doc, err := FetchDoc(URL)
	if err != nil {
		err = fmt.Errorf("GoQuery: %w", err)
		return
	}
	// Find the element
	q := doc.Find(MarketPriceSelector)

	// Check the no. elements
	if q.Size() == 0 {
		err = fmt.Errorf(
			"Doc.Find: No element found with the selector: %s",
			MarketPriceSelector,
		)
		return
	}

	var f string // filed
	var e bool   // exist
	var v string // value
	q.Each(func(_ int, s *goquery.Selection) {
		// get the data fields
		f, e = s.Attr("data-field")
		if !e {
			//throw error for attribute does not exist
		}

		// get the value
		v, e = s.Attr("value")
		if !e {
			//throw error for attribute does not exist
		}

		switch f {
		case PriceSelector:
			v = strings.ReplaceAll(v, ",", "")
			cm.Price, _ = strconv.ParseFloat(v, 8)
			break
		case ChangeSelector:
			v = strings.ReplaceAll(v, ",", "")
			cm.Change, _ = strconv.ParseFloat(v, 8)
			break
		case PercentageSelector:
			v = strings.ReplaceAll(v, ",", "")
			cm.ChangePercentage, _ = strconv.ParseFloat(v, 64)
			break
		default:
		}
	})

	// error for string to float conversion
	if err != nil {
		err = fmt.Errorf("String to Float: %w", err)
		return
	}
	return
}
