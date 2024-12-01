if __name__ == "__main__":
    list1: list[int] = []
    list2: list[int] = []

    # read values
    with open("../input.txt", "r") as input_file:
        for line in input_file.readlines():
            values: list[str] = line.split()
            list1.append(int(values[0]))
            list2.append(int(values[1]))

    # sort lists
    list1.sort()
    list2.sort()

    # sum differences between elements (list2 - list1)
    diff_sum: int = 0
    if len(list1) == len(list2):
        for value1, value2 in zip(list1, list2):
            diff_sum += abs(value2 - value1)
    else:
        raise ValueError("list lengths do not match")

    print(diff_sum)
