from math import prod
from pathlib import Path
import re


if __name__ == "__main__":
    mul_command_regex: str = r"mul\(([0-9]*)\,([0-9]*)\)"
    do_command_regex: str = r"do\(\)"
    dont_command_regex: str = r"don't\(\)"

    # read input
    corrupted_memory: str = open(f"{Path(__file__).parent}/../input.txt", "r").read()

    # for each match of the mul regex, get the start index of the match and the value of the arguments parsed to an int
    mul_command_matches: list[tuple[int, tuple[int, ...]]] = []
    for mul_command_match in re.finditer(mul_command_regex, corrupted_memory):
        mul_command_matches.append(
            (mul_command_match.start(), tuple(map(int, mul_command_match.groups())))
        )

    result: int = 0
    for mul_command_match in mul_command_matches:
        corrupted_memory_up_to_mul: str = corrupted_memory[: mul_command_match[0]]

        do_command_matches: list[re.Match] = list(
            re.finditer(do_command_regex, corrupted_memory_up_to_mul)
        )
        dont_command_matches: list[re.Match] = list(
            re.finditer(dont_command_regex, corrupted_memory_up_to_mul)
        )

        if dont_command_matches and (
            (not do_command_matches)
            or dont_command_matches[-1].end() > do_command_matches[-1].end()
        ):
            continue

        result += prod(mul_command_match[1])

    print(result)
