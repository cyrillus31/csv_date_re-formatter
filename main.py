import csv
from os import path
import os
from typing import Generator

def convert_csv(input_file, output_file):
    # Open the input CSV for reading.
    with open(input_file, mode='r', newline='', encoding='utf-8') as infile:
        reader = csv.DictReader(infile)
        
        # Define the output CSV header.
        output_fieldnames = ["data", "time", "open", "high", "loi", "close", "volume"]
        
        # Open the output CSV for writing.
        with open(output_file, mode='w', newline='', encoding='utf-8') as outfile:
            writer = csv.DictWriter(outfile, fieldnames=output_fieldnames)
            # writer.writeheader()
            
            for row in reader:
                # Extract original time value.
                time_field = row["<TIME>"] if "<TIME>" in row else row["TIME"]
                # Convert time to 4-character format (assumes original format is HHMMSS).
                time_4char = time_field[:-2] if len(time_field) >= 6 else time_field

                # Extract values from the row according to the mapping.
                new_row = {
                    "data": row["<DATE>"] if "<DATE>" in row else row["DATE"],
                    "time": time_4char,
                    "open": row["<OPEN>"] if "<OPEN>" in row else row["OPEN"],
                    "high": row["<HIGH>"] if "<HIGH>" in row else row["HIGH"],
                    "loi": row["<LOW>"] if "<LOW>" in row else row["LOW"],
                    "close": row["<CLOSE>"] if "<CLOSE>" in row else row["CLOSE"],
                    "volume": row["<VOL>"] if "<VOL>" in row else row["VOL"]
                }
                writer.writerow(new_row)


def files_generator(in_dir: str, out_dir) -> Generator:
    for root, dirs, files in os.walk(in_dir):
        for file in files:
            yield path.join(root, file), path.join(out_dir, file)

def create_dirs(in_dir: str, out_dir: str) -> None:
    try:
        os.makedirs(input_dir)
    except FileExistsError:
        pass
    try:
        os.makedirs(output_dir)
    except FileExistsError:
        pass


if __name__ == "__main__":
    input_dir = path.abspath("input_files")
    output_dir = path.abspath("output_files")
    create_dirs(input_dir, output_dir)
    input(f"Put all files you want to convert into:\n{input_dir}\n\nPress ENTER when ready.")

    for in_filepath, out_filepath in files_generator(input_dir, output_dir):
        try:
            convert_csv(in_filepath, out_filepath)
            print(f"Conversion complete. Output saved to {output_dir}")
        except Exception as e:
            print(f"Couldn't process file {in_filepath} due to: {e}.")


