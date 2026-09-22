# def ticket_total(price, quantity):
#     total = price * quantity
#     print(total)
# amount = ticket_total("7", 3)
# print(amount)

# ticket_total()

# def ticket_total(price, quantity):
#     total = price * quantity
#     return total

# amount = ticket_total(7, 3)
# print(amount)

# def passing_scores(scores):
#     passed = []
#     for index in range(len(scores) - 1):
#         if scores[index] > 50:
#             passed.append(scores[index])
#     return passed
# print(passing_scores([49, 50, 80, 65]))

# def passing_scores(scores):
#     passed = []
#     for score in scores:
#         if score >= 50:
#             passed.append(score)
#     return passed

# print(passing_scores([49, 50, 80, 65]))

def add_tag(profile, tag):
    updated = profile.copy()
    updated["tags"].append(tag)
    return updated
original = {"name": "Ada", "tags": ["python"]}
changed = add_tag(original, "testing")
print(original["tags"])
print(changed is original)
print(changed["tags"] is original["tags"])