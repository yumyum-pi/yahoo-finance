package test

import (
	"fmt"
	"testing"

	yahooFinace "github.com/yumyum-pi/yahooFinance"
)

func TestGetPrice(t *testing.T) {
	tanla := yahooFinace.NewTicker("TANLA.NS")
	cm, err := tanla.GetPrice()
	if err != nil {
		t.Errorf("Unexpected Error: %s", err)
	} else {
		if cm.Change == 0 {
			t.Errorf("Change should not be zero")
		} else if cm.Price == 0 {
			t.Errorf("Price should not be zero")
		} else if cm.ChangePercentage == 0 {
			t.Errorf("ChangePercentage should not be zero")
		}
	}
}

// TODO make test cases for get price
func TestGetPriceFail(t *testing.T) {
	tanla := yahooFinace.NewTicker("TANLLLA.NS")
	_, err := tanla.GetPrice()
	if err == nil {
		t.Errorf("was expecting an error")
	}

}

// TODO make test cases for get summery
func TestGetSummery(t *testing.T) {
	tanla := yahooFinace.NewTicker("TANLA.NS")
	summery, err := tanla.GetSummery()
	if err != nil {
		t.Errorf("Unexpected Error: %s", err)
	}
	fmt.Println(summery)
}
