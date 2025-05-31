import os

class InterfaceQuestions:
    """
    Talks to the user
    1. Put files into input_directory/
    2. Loop:
        Here is the first row...
        Which columns containes {date|time|high|low|volume}?
    """
    def __init__(
            self, 
            input_dir: str = "input_dir", 
            output_dir: str = "output_files"
    ) -> None:

        self._input_dir = input_dir
        self._output_dir = output_dir

        # Create directories
        if not os.path.exists(self._input_dir):
            os.mkdir(self._input_dir)

        if not os.path.exists(self._output_dir):
            os.mkdir(self._output_dir)

    def introduction(self) -> None:
        if not os.path.exists(self._input_dir):
            os.mkdir(self._input_dir)

        msg = f"Put files that you want to convert into '{self._input_dir}' directory. Press ENTER"
        input(msg)

    def map_responses(self, row: list[str], column_names: list[str]) -> dict[str, int]:
        result = dict()
        print(f"Here are all the rows:")

        for index, column_elem in enumerate(row):
            print(f"{index+1} {column_elem}) ", end="\t")

        for column_name in column_names:
            column_index = int(input(f"Which column is {column_name}? Enter a number")) - 1
            result[column_name] = column_index

        return result
