import pandas as pd
import pyarrow


class Tools:
    def ping(self) -> None:
        print("Pong")

    def read_csv(self, filepath: str) -> None:
        self.metadata = pd.read_csv(filepath, low_memory=False)


tools = Tools()
