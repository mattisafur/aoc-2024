from math import prod
from pathlib import Path
from re import findall


if __name__ == "__main__":
    # read input
    corrupted_memotry: str = open(f"{Path(__file__).parent}/../input.txt", "r").read()

    # extract commands using regex
    extracted_command_arguments: list[tuple[str, ...]] = findall(
        r"mul\(([0-9]*)\,([0-9]*)\)", corrupted_memotry
    )

    # parse extracted strs to ints
    parsed_command_arguments: list[tuple[int, ...]] = [
        tuple(map(int, arg_pair)) for arg_pair in extracted_command_arguments
    ]

    # sum product of int pairs
    result = sum(prod(arg_pair) for arg_pair in parsed_command_arguments)

    print(result)
