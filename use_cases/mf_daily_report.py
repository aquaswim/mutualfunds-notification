from array import array
from outputs.output_interface import OutputInterface
from data_source.data_source_interface import DataSourceInterface
from entities.mutual_fund import MutualFund


def _format_mf_data(mf_data: MutualFund) -> str:
    latestNAB = mf_data.nav_logs[-1]
    yesterdayNAB = mf_data.nav_logs[-2]
    diff = latestNAB - yesterdayNAB
    icon = "🟢"
    if (diff <= 0):
        icon = "🔴"
    return "{}:\t\t{:.2f} ({:.2f}) {}".format(mf_data.name, latestNAB, diff, icon)


class MutualFundDailyReporter:
    def __init__(self, data_source: DataSourceInterface, output: OutputInterface):
        self.data_source = data_source
        self.output = output

    def run(self, product_ids: array):
        self.output.print("Laporan Reksa Dana Hari Ini\n")
        for product_id in product_ids:
            mf_data = self.data_source.get_mutual_funds_data(product_id)
            self.output.print("{}\n".format(_format_mf_data(mf_data)))
        self.output.flush()
