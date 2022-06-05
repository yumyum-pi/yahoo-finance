package yahooFinace

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// CurrentMarket is struct to store current market price.
type CurrentMarket struct {
	Price            float64
	Change           float64
	ChangePercentage float64
}

// Bid is a struct to store bids.
// It contains price and quantity of the share
type Bid struct {
	price    float64
	quantity float64
}

// Summery struct contains PreviousClose,Open,Volume,AvgVolume,MarketCap,Beta5YMonthly,
// PE,EPS,EarningDate,ForwardDividend,ExDividendDate,TargetEst1yr,Bid,Ask,
// DayRange,WeekRange52,
type Summery struct {
	PreviousClose float64
	Open          float64
	Volume        float64
	AvgVolume     float64
	MarketCap     float64
	Beta5YMonthly float64
	PE            float64
	EPS           float64

	EarningDate     string
	ForwardDividend string
	ExDividendDate  string
	TargetEst1yr    float64

	Bid Bid
	Ask Bid

	DayRange    [2]float64
	WeekRange52 [2]float64
}

// Ticker struct stores symbol of a share
type Ticker struct {
	symbol string
}

// NewTicker returns a new ticker with the given symbol of a share
func NewTicker(symbol string) (T Ticker) {
	T = Ticker{symbol}
	return T
}

// GetPrice is a method of Ticker. The function returns Current Market Price by scraping
// Yahoo Financial
func (T *Ticker) GetPrice() (cm CurrentMarket, err error) {
	URL := fmt.Sprintf("%s/%s", SCRAPE_URL, T.symbol)

	// get the html element that contains the current market price
	elm, err := FetchDoc(URL, "#quote-header-info fin-streamer")
	if err != nil {
		err = fmt.Errorf("FetchDoc: %s", err)
		return
	}

	// predefined valuables
	var f string // data field
	var e bool   // if exist
	var v string // value of data field

	// #quote-header-info contains multiple fin-streamer.
	// the function will loop over each element
	elm.Each(func(_ int, s *goquery.Selection) {
		// need fin-streamer with specific attribute
		// data-field contains the Key of the data
		// value contains the value of the data

		// get the data fields attribute
		f, e = s.Attr("data-field")
		if !e {
			//throw error for attribute does not exist
			err = fmt.Errorf("attribute not found: %s", "data-field")
			return
		}

		// get the value attribute
		v, e = s.Attr("value")
		if !e {
			//throw error for value does not exist
			err = fmt.Errorf("value not found: %s", "value")
			return
		}

		// switch over different data fields
		switch f {
		case "regularMarketPrice":
			v = strings.ReplaceAll(v, ",", "")
			cm.Price, err = strconv.ParseFloat(v, 8)
			if err != nil {
				return
			}
			break
		case "regularMarketChange":
			v = strings.ReplaceAll(v, ",", "")
			cm.Change, err = strconv.ParseFloat(v, 8)
			if err != nil {
				return
			}
			break
		case "regularMarketChangePercent":
			v = strings.ReplaceAll(v, ",", "")
			cm.ChangePercentage, err = strconv.ParseFloat(v, 64)
			if err != nil {
				return
			}
			break
		default:
		}

		// reset values
		v = ""
		f = ""
		e = false
	})

	return
}

// GetSummery is a method of Ticker. The function returns Summery by scraping Yahoo Financial
func (T *Ticker) GetSummery() (summary Summery, err error) {
	URL := fmt.Sprintf("%s/%s", SCRAPE_URL, T.symbol)

	// get the html element that contains the summary
	doc, err := FetchDoc(URL, "#quote-summary tr")
	if err != nil {
		err = fmt.Errorf("FetchDoc: %s", err)
		return
	}

	// predefined variable
	var value string

	// #quote-summary contains multiple a table. tr is the table element
	// the function will loop over each tr element
	doc.Each(func(i int, s *goquery.Selection) {
		// the tr contains two children. 1st contains key. 2nd contains value
		// get value of of the current tr element
		s.Children().Each(func(j int, c *goquery.Selection) {
			if j == 1 {
				value = c.Text()
			}
		})

		// not capturing the key. Identifying value using the index
		switch i {
		case 0: // Previous Close
			summary.PreviousClose, err = String2Float(value)
			if err != nil {
				return
			}
			break
		case 1: // Open
			summary.Open, err = String2Float(value)
			if err != nil {
				return
			}
			break
		case 2: // Bid
			// value contains price and quantity
			// separating price and quantity
			bid := strings.Split(value, " x ")
			summary.Bid.price, err = String2Float(bid[0])
			if err != nil {
				return
			}
			summary.Bid.quantity, err = String2Float(bid[1])
			if err != nil {
				return
			}
			break
		case 3: // Ask
			// value contains price and quantity
			// separating price and quantity
			ask := strings.Split(value, " x ")
			summary.Ask.price, err = String2Float(ask[0])
			if err != nil {
				return
			}
			summary.Ask.quantity, err = String2Float(ask[1])
			if err != nil {
				return
			}
			break
		case 4: // Day Range
			// value contains low and high
			// separating low and high
			r := strings.Split(value, " - ")
			summary.DayRange[0], err = String2Float(r[0])
			if err != nil {
				return
			}
			summary.DayRange[1], err = String2Float(r[1])
			if err != nil {
				return
			}
			break
		case 5: // 52 week range
			// value contains low and high
			// separating low and high
			r := strings.Split(value, " - ")
			summary.WeekRange52[0], err = String2Float(r[0])
			if err != nil {
				return
			}
			summary.WeekRange52[1], err = String2Float(r[1])
			if err != nil {
				return
			}
			break
		case 6: // Volume
			summary.Volume, err = String2Float(value)
			if err != nil {
				return
			}
			break
		case 7: // Average Volume
			summary.AvgVolume, err = String2Float(value)
			if err != nil {
				return
			}
			break
		case 8: // Market Cap
			summary.MarketCap, err = String2Float(value)
			if err != nil {
				return
			}
			break
		case 9: //Beta 5 year
			summary.Beta5YMonthly, err = String2Float(value)
			if err != nil {
				return
			}
			break
		case 10: // PE ratio
			summary.PE, err = String2Float(value)
			if err != nil {
				return
			}
			break
		case 11: // EPS
			summary.EPS, err = String2Float(value)
			if err != nil {
				return
			}
			break
		case 12: // Earning Date
			// TODO convert string to date
			// also need to take care of N/A value
			summary.EarningDate = value
			break
		case 13: // Forward Dividend
			summary.ForwardDividend = value
			break
		case 14: // Ex-Dividend Date
			// to do convert string to date
			// TODO convert string to date
			// also need to take care of N/A value
			summary.ExDividendDate = value
			break
		case 15: // Target
			summary.TargetEst1yr, err = String2Float(value)
			if err != nil {
				return
			}
			break
		default:
		}
	})

	return
}
