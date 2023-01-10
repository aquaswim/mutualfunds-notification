from .data_source_interface import DataSourceInterface


class PasardanaDataSource(DataSourceInterface):
    def __init__(self):
        pass

    def get_mutual_funds_data(self, product_id: str) -> MutualFund:
        pass