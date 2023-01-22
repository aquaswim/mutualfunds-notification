package main

import (
	"aya-money-go/internal/data_source"
	"aya-money-go/internal/output"
	"aya-money-go/internal/use_cases"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"strings"
	"time"
)

var dcWebhookUrl string
var reportCron string
var mfProductIds []string

func main() {
	mfds := data_source.NewBareksaMutualFundDataSource()
	out := output.NewOutputDiscordWebhook(dcWebhookUrl)

	reporter := use_cases.NewUseCaseMutualFundReport(mfds, out)

	c := cron.New(cron.WithSeconds())

	c.AddFunc(reportCron, func() {
		err := reporter.Generate(mfProductIds)
		if err != nil {
			fmt.Printf("error generate report at %s err: %s\n", time.Now().Format("2006-01-02 15:04:05"), err)
		}
	})
	c.Run()
}

func init() {
	dcWebhookUrl = os.Getenv("DISCORD_WEBHOOK_URL")
	if dcWebhookUrl == "" {
		log.Fatalln("need to set DISCORD_WEBHOOK_URL env variable")
	}
	reportCron = os.Getenv("REPORT_CRONTAB")
	if reportCron == "" {
		reportCron = "0 17 * * *"
	}
	mfProductIdsRaw := os.Getenv("MF_PRODUCT_IDS")
	mfProductIds = strings.Split(mfProductIdsRaw, ",")
	if len(mfProductIds) < 1 {
		log.Fatalln("need to set MF_PRODUCT_IDS env variable to comma separated value of mutual fund product id in bareksa")
	}
}
