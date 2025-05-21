import os
from directives import Directives
from file_processer import FileProcesser


class Interface:
    def __init__(self, input_dir: str) -> None:
        self.input_dir: str = input_dir
        self._file: str = ""
        self._fp: FileProcesser | None = None
        self.default_order = ["date", "time", "open", "high", "low", "close", "volume"]
        # self.desired_order: dict[str, int] = {

    def pick_file(self) -> None:

        root, dirs, files = next(os.walk(self.input_dir))

        for i, file in enumerate(files):
            print(f"{i+1}) {file}")
        print()

        self._file = files[int(input("Which file do you want to convert? Enter number: ")) - 1]

        self._fp = FileProcesser(os.path.join(self.input_dir, self._file))

        print(self._fp.get_first_line())
        print()

        answer = input("Does the first line containe data? y/n: ")
        if "n" in answer.lower():
            self._fp.skip_first_line = True

    def which_column(self) -> None:

    def setup(self) -> None:
        self.pick_file()
        print("Pick numbers that correspond to rows you want to have:")
        
