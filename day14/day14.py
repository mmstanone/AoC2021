def parse_input(file):
    with open(file, 'r') as f:
        lines = list(map(lambda x: x.strip(), f.readlines()))
        polymer = lines[0]
        rules = lines[2:]
    
    actual_rules = {}
    for rule in rules:
        rule = rule.replace(' ', '').replace('\n', '').split('->')
        actual_rules[rule[0]] = rule[1]
        
    return list(polymer), actual_rules

def actual_grow(polymer, rules, iterations):
    pairs = {}
    for k in rules:
        pairs[k] = 0
    
    for i in range(len(polymer) - 1):
        pairs[polymer[i] + polymer[i + 1]] += 1
    
    old = pairs.copy()
    for _ in range(iterations):
        new = old.copy()
        for pair in old:
            if old[pair] == 0:
                continue
            first_pair = pair[0] + rules[pair]
            second_pair = rules[pair] + pair[1]
            new[first_pair] += old[pair]
            new[second_pair] += old[pair]
            new[pair] -= old[pair]
            
        old = new.copy()
        
    elems = {}
    for k in old:
        elems[k[0]] = elems.get(k[0], 0) + old[k]
        elems[k[1]] = elems.get(k[1], 0) + old[k]
    print(max(elems.values()) - min(elems.values()), (max(elems.values()) - min(elems.values())) / 2)
            

def main():
    polymer, rules = parse_input('puzzle14.txt')
    actual_grow(polymer, rules, 10)
    actual_grow(polymer, rules, 40)
    
main()