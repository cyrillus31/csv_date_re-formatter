import os


class Interface:
    def __init__(self, input_dir) -> None:
        self.input_dir = input_dir
        self._file = ""

    def pick_file(self) -> None:

        root, dirs, files = next(os.walk(self.input_dir))

        for i, file in enumerate(files):
            print(f"{i+1}) {file}")

        self._file = files[int(input("Which file do you want to convert?")) - 1]

    def start(self) -> None:
        print("Pick letters that correspond to rows you want to have:")

