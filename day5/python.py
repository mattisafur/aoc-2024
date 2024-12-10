from pathlib import Path


class Solution:
    def __init__(self, rules: dict[int, set[int]], updates: list[list[int]]) -> None:
        self.rules: dict[int, set[int]] = rules
        self.updates: list[list[int]] = updates

    def get_first_invalid_index(
        self, page: int, check_against: list[int]
    ) -> int | None:
        page_rules: set[int] | None = self.rules.get(page)
        if not page_rules:
            return None
        for idx, previous_page in enumerate(check_against):
            if previous_page in page_rules:
                return idx
        return None

    def is_page_valid(self, page: int, check_against: list[int]) -> bool:
        page_rules: set[int] | None = self.rules.get(page)
        if not page_rules:
            return True
        for previous_page in check_against:
            if previous_page in page_rules:
                return False
        return True

    def is_update_valid(self, update: list[int]) -> bool:
        for idx, page in enumerate(update):
            if not self.is_page_valid(page, update[:idx]):
                return False
        return True

    def part1(self) -> int:
        sum_of_middles: int = 0
        for update in self.updates:
            if not self.is_update_valid(update):
                continue
            sum_of_middles += update[len(update) // 2]
        return sum_of_middles

    def part2(self) -> int:
        sum_of_middles: int = 0
        for update in self.updates:
            if self.is_update_valid(update):
                continue

            idx: int = 1
            while idx < len(update):
                page: int = update[idx]
                invalid_idx: int | None = self.get_first_invalid_index(
                    page, update[:idx]
                )
                if invalid_idx is None:
                    idx += 1
                    continue
                update[idx], update[invalid_idx] = update[invalid_idx], update[idx]
                idx = invalid_idx

            sum_of_middles += update[len(update) // 2]
        return sum_of_middles


if __name__ == "__main__":
    # read and parse input
    rules: dict[int, set[int]] = {}
    updates: list[list[int]] = []
    rules_and_updates: list[str] = (
        open(f"{Path(__file__).parent}/input.txt").read().strip().split("\n\n")
    )
    unparsed_rules: list[str] = rules_and_updates[0].split()
    unparsed_updates: list[str] = rules_and_updates[1].split()
    for unparsed_rule in unparsed_rules:
        rule_values: list[int] = [int(page) for page in unparsed_rule.split("|")]
        rules.setdefault(rule_values[0], set()).add(rule_values[1])
    for unparsed_update in unparsed_updates:
        updates.append([int(page) for page in unparsed_update.split(",")])

    # solve
    solution = Solution(rules, updates)
    print(f"part1: {solution.part1()}")
    print(f"part2: {solution.part2()}")
