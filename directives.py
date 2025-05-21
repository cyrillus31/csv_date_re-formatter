from dataclasses import dataclass


@dataclass
class Directives:
    date_row: str
    time_row: str
    open_row: str
    high_row: str
    low_row: str
    close_row: str
    volume_row: str
    order: list[str] = ["date", "time", "open", "high", "low", "close", "volume"]
