package data_source

import "testing"

func TestDataSourceBareksa(t *testing.T) {
	ds := ds{}
	// id 316 get from bareksa
	// ex: https://www.bareksa.com/id/data/mutualfund/316/sucorinvest-flexi-fund
	data, err := ds.GetMutualFundData("316")
	if err != nil {
		t.Fatal("error get mutualfund data", err)
	}
	if data.Name == "" {
		t.Fatal("mutual fund data name is empty")
	}
	if data.LatestDate == "" {
		t.Fatal("mutual fund data latestDate is empty")
	}
	if len(data.NavHistory) < 1 {
		t.Fatal("mutual fund data history is less than 1")
	}
}
