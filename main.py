import os

from data_source.bareksa import BareksaDataSource
from outputs.discord_webhook import DiscordWebhook
from use_cases.mf_daily_report import MutualFundDailyReporter


def main():
    data_source = BareksaDataSource()
    output = DiscordWebhook(
        webhook_url=os.environ.get('DISCORD_WEBHOOK_URL')
    )
    mf_daily_reporter = MutualFundDailyReporter(
        data_source=data_source,
        output=output
    )
    # list_mf_id = [
    #     "316",
    #     "136",
    #     "4447",
    #     "339",
    #     "2106",
    #     "2452",
    #     "3397",
    #     "2548",
    #     "4100",
    #     "1878",
    #     "1742"
    # ]

    list_mf_id = os.environ.get("MF_PRODUCT_IDS").split(",")
    if len(list_mf_id) < 1:
        print("Need atlest list of mutual fund id in env \"MF_PRODUCT_IDS\"")
        exit(-1)
    print("will report mutual fund ids: ", list_mf_id)
    mf_daily_reporter.run(list_mf_id)
    print("Done")


# Press the green button in the gutter to run the script.
if __name__ == '__main__':
    main()
