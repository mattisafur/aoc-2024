from pathlib import Path


if __name__ == "__main__":
    # read and parse input
    rules: dict[int, set[int]] = {}
    updates: list[list[int]] = []
    rules_and_updates: list[str] = (
        open(f"{Path(__file__).parent}/../input.txt").read().strip().split("\n\n")
    )
    unparsed_rules: list[str] = rules_and_updates[0].split()
    unparsed_updates: list[str] = rules_and_updates[1].split()
    for unparsed_rule in unparsed_rules:
        rule_values: list[int] = [int(page) for page in unparsed_rule.split("|")]
        rules.setdefault(rule_values[0], set()).add(rule_values[1])
    for unparsed_update in unparsed_updates:
        updates.append([int(page) for page in unparsed_update.split(",")])

    sum_of_middle_page_numbers: int = 0
    for update in updates:
        update_valid: bool = True
        for page_idx, page in enumerate(update):
            if page not in rules:
                continue
            for page_to_check in rules[page]:
                if page_to_check in update[:page_idx]:
                    update_valid = False
                    break
            if not update_valid:
                break
        if update_valid:
            sum_of_middle_page_numbers += update[len(update) // 2]

    print(sum_of_middle_page_numbers)
