from dataclasses import dataclass


@dataclass
class Directives:
    date_column: str
    time_column: str
    open_column: str
    high_column: str
    low_column: str
    close_column: str
    volume_column: str
    order: list[str] = ["date", "time", "open", "high", "low", "close", "volume"]
