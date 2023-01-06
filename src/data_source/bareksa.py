import requests
from src.data_source.data_source_interface import DataSourceInterface
from src.entities.mutual_fund import MutualFund


def _create_mf_entity_from_response(resp) -> MutualFund:
    reksa_dana_data = resp["data"]["datas"][0]
    return MutualFund(
        name=reksa_dana_data["pname"],
        latest_nav_date=resp["data"]["enddate"],
        nav_logs=list(map(lambda nav: float(nav['value']), reksa_dana_data["nav"]))
    )


class BareksaDataSource(DataSourceInterface):
    def __init__(self):
        self.c_period = "mtd"

    def get_mutual_funds_data(self, product_id: str) -> MutualFund:
        url = "https://www.bareksa.com/ajax/mutualfund/nav/product1/?id={}&cperiod={}".format(product_id, self.c_period)

        payload = {}
        headers = {
            'X-Requested-With': 'XMLHttpRequest',
        }

        response = requests.request("GET", url, headers=headers, data=payload)
        # todo need to handle error
        return _create_mf_entity_from_response(response.json())
