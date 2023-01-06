import array


class MutualFund:
    def __init__(self, name=str, latest_nav_date=str, nav_logs=array):
        self.name = name
        self.latest_nav_date = latest_nav_date
        self.nav_logs = nav_logs

    def print(self) -> str:
        return "hello world, {}".format(self.name)
