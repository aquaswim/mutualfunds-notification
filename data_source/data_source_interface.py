from entities.mutual_fund import MutualFund


class DataSourceInterface:
    def get_mutual_funds_data(self, product_id: str) -> MutualFund:
        """GET MUTUAL FUND DATA"""
        pass
