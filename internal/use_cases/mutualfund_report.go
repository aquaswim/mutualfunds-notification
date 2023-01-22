package use_cases

import (
	"aya-money-go/internal/contracts"
	"aya-money-go/internal/entities"
	"fmt"
)

type MutualFundDailyReporter struct {
	ds  contracts.MutualFundDataSource
	out contracts.Output
}

func NewUseCaseMutualFundReport(ds contracts.MutualFundDataSource, out contracts.Output) MutualFundDailyReporter {
	return MutualFundDailyReporter{ds: ds, out: out}
}

func (m *MutualFundDailyReporter) Generate(productIds []string) error {
	err := m.out.Print("Laporan Reksa Dana Hari Ini\n")
	if err != nil {
		return err
	}
	for _, productId := range productIds {
		data, err := m.ds.GetMutualFundData(productId)
		if err != nil {
			fmt.Printf("error getting productId: %s err: %s", productId, err)
			continue
		}
		err = m.out.Print(m.formatMfData(data) + "\n")
		if err != nil {
			fmt.Printf("error print productId: %s err: %s", productId, err)
			continue
		}
	}
	err = m.out.Flush()
	if err != nil {
		return err
	}
	return nil
}

const (
	ic_nab_up   = "🟢"
	ic_nab_down = "🔴"
)

func (m *MutualFundDailyReporter) formatMfData(data *entities.MutualFund) string {
	latestNAB := &data.NavHistory[len(data.NavHistory)-1]
	yesterdayNAB := &data.NavHistory[len(data.NavHistory)-2]
	diff := latestNAB.Value - yesterdayNAB.Value
	icon := ic_nab_up
	if diff <= 0 {
		icon = ic_nab_down
	}
	return fmt.Sprintf("%s:\t\t%.2f (%.2f) %s", data.Name, latestNAB.Value, diff, icon)
}
