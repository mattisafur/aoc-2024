from pathlib import Path


if __name__ == "__main__":
    list1: list[int] = []
    list2: list[int] = []

    # read values
    with open(f"{Path(__file__).parent}/../input.txt", "r") as input_file:
        for line in input_file.readlines():
            values: list[str] = line.split()
            list1.append(int(values[0]))
            list2.append(int(values[1]))

    similarity_score: int = 0
    for value in list1:
        similarity_score += list2.count(value) * value

    print(similarity_score)
