package data_source

import (
	"aya-money-go/internal/contracts"
	"aya-money-go/internal/entities"
	"aya-money-go/internal/utils"
	"fmt"
	"strconv"
	"time"
)

type ds struct{}

type bareksaMfResp struct {
	Status bool `json:"status"`
	Data   struct {
		Auth        bool   `json:"auth"`
		RedirectUrl string `json:"redirect_url"`
		Startdate   string `json:"startdate"`
		Enddate     string `json:"enddate"`
		Subtitle    string `json:"subtitle"`
		UnitY       string `json:"unitY"`
		Datas       []struct {
			Pid   string `json:"pid"`
			Ptype string `json:"ptype"`
			Idate string `json:"idate"`
			Inav  string `json:"inav"`
			Pname string `json:"pname"`
			Nav   []struct {
				Id    string `json:"id"`
				Date  string `json:"date"`
				Value string `json:"value"`
			} `json:"nav"`
		} `json:"datas"`
	} `json:"data"`
}

func (d ds) GetMutualFundData(productId string) (*entities.MutualFund, error) {
	resp := new(bareksaMfResp)
	url := fmt.Sprintf("https://www.bareksa.com/ajax/mutualfund/nav/product1/?id=%s&cperiod=%s", productId, "mtd")
	err := utils.SendJsonRequest("GET", url, nil, resp)
	if err != nil {
		return nil, err
	}
	// do formatting stuff
	out := new(entities.MutualFund)
	if len(resp.Data.Datas) < 1 || resp.Status == false {
		return nil, fmt.Errorf("bareksa datasource return status non true")
	}
	out.Name = resp.Data.Datas[0].Pname
	out.LatestDate = resp.Data.Enddate
	out.NavHistory = make([]entities.NavLog, len(resp.Data.Datas[0].Nav))
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, fmt.Errorf("error can't load location Asia/Jakarta for parsing bareksa datasource")
	}
	for i, nav := range resp.Data.Datas[0].Nav {
		parsedVal, err := strconv.ParseFloat(nav.Value, 64)
		if err != nil {
			return nil, err
		}
		parsedDate, err := time.ParseInLocation("2006-01-02", nav.Date, location)
		if err != nil {
			return nil, err
		}

		out.NavHistory[i].Value = float32(parsedVal)
		out.NavHistory[i].Date = parsedDate
	}

	return out, nil
}

func NewBareksaMutualFundDataSource() contracts.MutualFundDataSource {
	return ds{}
}
