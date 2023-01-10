from .output_interface import OutputInterface


class ConsoleOutput(OutputInterface):
    def print(self, text: str):
        print(text)
