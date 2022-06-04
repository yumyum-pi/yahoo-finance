# Go module to download stock data Yahoo! Finance

<table border=1 cellpadding=10><tr><td>

#### \*\*\* IMPORTANT LEGAL DISCLAIMER \*\*\*

---

**Yahoo!, Y!Finance, and Yahoo! finance are registered trademarks of
Yahoo, Inc.**

yfinance is **not** affiliated, endorsed, or vetted by Yahoo, Inc. It's
an open-source tool that uses Yahoo's publicly available APIs, and is
intended for research and educational purposes.

**You should refer to Yahoo!'s terms of use**
([here](https://policies.yahoo.com/us/en/yahoo/terms/product-atos/apiforydn/index.htm),
[here](https://legal.yahoo.com/us/en/yahoo/terms/otos/index.html), and
[here](https://policies.yahoo.com/us/en/yahoo/terms/index.htm)) **for
details on your rights to use the actual data downloaded. Remember - the
Yahoo! finance API is intended for personal use only.**

</td></tr></table>


Go module to scrape stock data from Yahoo Finance

## Installation
yahooFinance requires Go1.1+ and is tested on Go1.8+.
Make sure go modules are working in your project.

    $ go get github.com/yumyum-pi/yahoo-finance
    
## Examples
```Go
package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/PuerkitoBio/goquery"
)

package test

import (
	"fmt"

	yahooFinace "github.com/yumyum-pi/yahooFinance"
)

func main() {
	ticket := yahooFinace.NewTicker("TANLA.NS")
	currentMarket, err := ticket.GetPrice()
	if err != nil {
			log.Fatalf("Unexpected Error: %s", err)
	}
	
  fmt.Println(currentMarket.Price)
  fmt.Println(currentMarket.Change)
  fmt.Println(currentMarket.ChangePercentage)
}

```
