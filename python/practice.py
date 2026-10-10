# ticket = {
#     "title": "Laptop is not working",
#     "urgency": "high",
#     "status": "open"
# }
# ticket["affected_users"] = 15
# print(ticket["affected_users"])

ticket = {
    "title": "Campus Wi-Fi is down",
    "urgency": "high",
    "status": "open",
    "affected_users": 2
}
ticket["affected_users"] = 15
print(ticket["affected_users"])

tickets = [
    {"title": "Campus Wi-Fi is down", "status": "open"},
    {"title": "Laptop is not working", "status": "open"}
]
print(tickets[1]["title"])

tickets.append({"title": "Software is not opening", "status": "open"})
print(len(tickets))

for i in tickets:
    print(i["title"])