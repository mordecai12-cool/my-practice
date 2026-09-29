def ticket_total(price, quantity):
    total = price * quantity
    print(total)
amount = ticket_total("7", 3)
print(amount)

#Rewrite the function and its call so the function returns an integer total and does not print internally. Print the returned result once outside the function. You may assume numeric integer inputs in the corrected version.

def ticket_total(price, quantity):
    return price * quantity
amount = ticket_total(7, 3)
print(amount)

#Explain why multiplication behaves this way and why amount does not contain the printed result.

#Return every score greater than or equal to 50, in the original order. Inputs are lists of integers from 0 to 100. An empty list must return an empty list.
def passing_scores(scores):
    passed = []
    for index in range(len(scores) - 1):
        if scores[index] > 50:
            passed.append(scores[index])
    return passed
print(passing_scores([49, 50, 80, 65]))

#Write three executable assertions: one for the pass boundary 50, one for a single passing score, and one for an empty list.
#Correct the function without changing the input list.
#Predict the current printed output. Identify both independent defects and explain which result each defect loses. [5 marks]

def add_tag(profile, tag):
    updated = profile.copy()
    updated["tags"].append(tag)
    return updated
original = {"name": "Ada", "tags": ["python"]}
changed = add_tag(original, "testing")
print(original["tags"])
print(changed is original)
print(changed["tags"] is original["tags"])

# def add_tag(profile, tag):
#     updated = profile.copy()
#     updated["tags"] = profile["tags"] + [tag]
#     return updated

# original = {"name": "Ada", "tags": ["python"]}
# changed = add_tag(original, "testing")

# print(original["tags"])                  
# print(changed is original)              
# print(changed["tags"] is original["tags"]) 

# Initial setup
# original = {"name": "Ada", "tags": ["python"]}
# changed = add_tag(original, "testing")

# # 1. Assert original tags remain unchanged
# assert original["tags"] == ["python"]

# # 2. Assert returned tags contain the new tag
# assert changed["tags"] == ["python", "testing"]

# # Mutation test: append another tag to the returned list
# changed["tags"].append("developer")

# # 3. Assert changed list updated successfully
# assert changed["tags"] == ["python", "testing", "developer"]

# # 4. Assert original list still has only its initial tag (proving total independence)
# assert original["tags"] == ["python"]


def summarise_amounts(raw_values):
    total = 0
    for raw in raw_values:
        try:
            total += int(raw)
        except:
            pass
    return {"total": total, "rejected": 0}
Required example: ["10", " 5 ", "bad", "-3", "0", ""] must return {"total": 15, "rejected": 3}. Inputs are always strings; no other type validation is required.