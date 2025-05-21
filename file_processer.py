import csv

class FileProcesser:
    def __init__(self, file_path: str, skip_first_line: bool = False) -> None:
        self.file_path = file_path
        self.skip_first_line = skip_first_line

    def get_first_line(self) -> list[str]:
        with open(self.file_path) as file:
            reader = csv.reader(file)
            for row in reader:
                return row
            return []

