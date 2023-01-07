import requests
from outputs.output_interface import OutputInterface


class DiscordWebhook(OutputInterface):
    def __init__(self, webhook_url: str, username: str = "aya"):
        self.webhook_url = webhook_url
        self.username = username
        self.buff = ""

    def print(self, text: str):
        self.buff = self.buff + text

    def flush(self):
        payload = {'content': self.buff,
                   'username': self.username}
        files = [

        ]
        headers = {}
        response = requests.request("POST", self.webhook_url, headers=headers, data=payload, files=files)
        self.buff = ""
