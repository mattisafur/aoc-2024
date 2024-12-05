import itertools
from pathlib import Path


def is_m_or_s(char: str) -> bool:
    return char == "M" or char == "S"


if __name__ == "__main__":
    # read word search
    word_search: list[list[str]] = [
        list(line)
        for line in open(f"{Path(__file__).parent}/../input.txt")
        .read()
        .strip()
        .splitlines()
    ]

    # count occurences
    occurences: int = 0
    for idv, idh in itertools.product(
        range(1, len(word_search) - 1), range(1, len(word_search[0]) - 1)
    ):
        if word_search[idv][idh] != "A":
            continue

        all_m_or_s: bool = True
        for dirv, dirh in itertools.product((-1, 1), repeat=2):
            if (
                word_search[idv + dirv][idh + dirh] != "M"
                and word_search[idv + dirv][idh + dirh] != "S"
            ):
                all_m_or_s = False
                break
        if not all_m_or_s:
            continue

        if word_search[idv - 1][idh - 1] == word_search[idv + 1][idh + 1]:
            continue
        if word_search[idv - 1][idh + 1] == word_search[idv + 1][idh - 1]:
            continue

        occurences += 1

    print(occurences)
