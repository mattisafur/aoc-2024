# This code has some edge cases where it dows not work, should work on it at some poiont.. if I'll remember x3
# if __name__ == "__main__":
#     # read report
#     data: list[list[int]] = []
#     with open("input.txt", "r") as file:
#         for report in file.readlines():
#             data.append([int(level) for level in report.split()])

#     # check reports
#     safe_reports_count: int = 0
#     safe_reports = []  # TODO remove
#     for report in data:
#         report_copy = report[:]  # TODO remove
#         report_safe: bool = True
#         dampener_available: bool = True

#         # check expected ascendance/descendance
#         report_ascending: bool = False
#         for diff in (value - report[0] for value in report[1:]):
#             if diff != 0:
#                 report_ascending = diff > 0
#                 break
#             if not dampener_available:
#                 report_safe = False
#             report.pop(0)
#             dampener_available = False
#         else:
#             report_safe = False

#         if not report_safe:
#             continue

#         idx: int = 0
#         while idx < len(report) - 1:
#             diff: int = report[idx + 1] - report[idx]
#             if (report_ascending == (diff > 0)) and (1 <= abs(diff) <= 3):
#                 idx += 1
#                 continue
#             if dampener_available:
#                 report.pop(idx + 1)
#                 dampener_available = False
#                 continue
#             report_safe = False
#             break

#         if report_safe:
#             safe_reports.append(report_copy)  # TODO remove
#             safe_reports_count += 1

#     print(safe_reports_count)
#     print(safe_reports)  # TODO remove


from pathlib import Path


def is_safe(report: list[int]) -> bool:
    return all(
        1 <= report[i + 1] - report[i] <= 3 for i in range(len(report) - 1)
    ) or all(-3 <= report[i + 1] - report[i] <= -1 for i in range(len(report) - 1))


if __name__ == "__main__":
    # read reports
    data: list[list[int]] = []
    with open(f"{Path(__file__).parent}/../input.txt", "r") as file:
        for report in file.readlines():
            data.append([int(level) for level in report.split()])

    safe_reports_count: int = 0
    for report in data:
        # check report
        if is_safe(report) or any(
            is_safe(report[:i] + report[i + 1 :]) for i in range(len(report))
        ):
            safe_reports_count += 1

    print(safe_reports_count)
