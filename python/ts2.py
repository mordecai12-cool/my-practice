# def ticket_total(price, quantity):
#     total = price * quantity
#     print(total)
# amount = ticket_total("7", 3)
# print(amount)

# def ticket_total(price, quantity):
#     total = price * quantity
#     return total  

# amount = ticket_total(7, 3) 
# print(amount)

def passing_scores(scores):
    passed = []
    for index in range(len(scores)):
        if scores[index] >= 50:
            passed.append(scores[index])
    return passed
print(passing_scores([49, 50, 80, 65]))

# 1. Pass boundary 50 (Checks that exactly 50 is included)
assert passing_scores([49, 50, 51]) == [50, 51]

# 2. Single passing score (Checks that a lone passing score is found)
assert passing_scores([85]) == [85]

# 3. Empty list (Checks that an empty list returns an empty list)
assert passing_scores([]) == []

def add_tag(profile, tag):
    updated = profile.copy()
    updated["tags"].append(tag)
    return updated
original = {"name": "Ada", "tags": ["python"]}
changed = add_tag(original, "testing")
print(original["tags"])
print(changed is original)
print(changed["tags"] is original["tags"])

def add_tag(profile, tag):
    updated = profile.copy()
    updated["tags"] = profile["tags"] + [tag]
    return updated
def add_tag(profile, tag):
    updated = profile.copy()
    updated["tags"] = profile["tags"].copy()
    updated["tags"].append(tag)
    return updated

# def summarise_amounts(raw_values):
#     total = 0
#     for raw in raw_values:
#         try:
#             total += int(raw)
#         except:
#             pass
#     return {"total": total, "rejected": 0}
