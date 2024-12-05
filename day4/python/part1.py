import itertools
from pathlib import Path


if __name__ == "__main__":
    # read word search
    word_search: list[list[str]] = [
        list(line)
        for line in open(f"{Path(__file__).parent}/../input.txt")
        .read()
        .strip()
        .splitlines()
    ]

    word: str = "XMAS"

    occurences: int = 0
    # iterate over each opint in array
    for idv, idh in itertools.product(
        range(len(word_search)), range(len(word_search[0]))
    ):
        # check possible moves from current location
        possible_vertical_moves: list[int] = [-1, 0, 1]
        possible_horizontal_moves: list[int] = [-1, 0, 1]
        if idv < len(word) - 1:
            possible_vertical_moves.remove(-1)
        if idv > len(word_search) - len(word):
            possible_vertical_moves.remove(1)
        if idh < len(word) - 1:
            possible_horizontal_moves.remove(-1)
        if idh > len(word_search[0]) - len(word):
            possible_horizontal_moves.remove(1)

        # run for each possible direction
        for dirv, dirh in itertools.product(
            possible_vertical_moves, possible_horizontal_moves
        ):
            if dirv == 0 and dirh == 0:
                continue

            for idx in range(len(word)):
                if word_search[idv + dirv * idx][idh + dirh * idx] != word[idx]:
                    break
            else:
                occurences += 1

    print(occurences)
