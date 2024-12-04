def is_safe(report: list[int]) -> bool:
    return all(
        1 <= report[i + 1] - report[i] <= 3 for i in range(len(report) - 1)
    ) or all(-3 <= report[i + 1] - report[i] <= -1 for i in range(len(report) - 1))


if __name__ == "__main__":
    # read reports
    data: list[list[int]] = []
    with open("input.txt", "r") as file:
        for report in file.readlines():
            data.append([int(level) for level in report.split()])

    safe_reports_count: int = 0
    for report in data:
        if is_safe(report):
            safe_reports_count += 1

    print(safe_reports_count)
