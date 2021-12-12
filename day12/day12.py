def parse_input(file):
    paths = dict()

    with open(file, "r") as f:
        for line in f:
            line = line.strip().split("-")
            if line[0] not in paths:
                paths[line[0]] = []
            if line[1] not in paths:
                paths[line[1]] = []
            paths[line[0]].append(line[1])
            paths[line[1]].append(line[0])

    return paths


def find_paths(paths, curr: str, visited, completed, full, possible_visit):
    if curr == "end":
        path = "-".join(completed)
        full.add(path)
        return
    if curr.islower() and curr in visited and not possible_visit:
        return

    if curr.islower() and curr in visited and possible_visit:
        possible_visit = False

    for next_node in paths[curr]:
        if next_node == 'start':
            continue
        find_paths(
            paths,
            next_node,
            visited | {curr},
            completed + [next_node],
            full,
            possible_visit
        )


def main():
    paths = parse_input("puzzle12.txt")

    visited = set()
    full = set()
    find_paths(paths, "start", visited, ["start"], full, True)
    print(len(full))


main()
