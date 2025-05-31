import csv
from os import read

class FileProcessor:
    """
    Access file.
    """
    def __init__(self, file_path: str) -> None:
        self.file_path = file_path

    def get_first_line(self) -> list[str]:
        with open(self.file_path) as csvfile:
            reader = csv.reader(csvfile)
            first_row = next(reader)
        return first_row
