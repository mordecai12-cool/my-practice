# The count variable is assigned in the if statement, and given the value 5:
numbers = [1, 2, 3, 4, 5]

if (count := len(numbers)) > 3:
    print(f"List has {count} elements")

# Assign the value "WEEKEND!" if the number is higher than 5, otherwise "Workday":
num = 6

x = "WEEKEND!" if num > 5 else "WORKDAY"

print(x)

# Assign:
# - "Fri" if num is 5
# - "Sat" if num is 6
# - "Sun" if num is 7
# - otherwise assign "weekday":
num = 5

x = "Fri" if num == 5 else "Sat" if num == 6 else "Sun" if num == 7 else "Weekday"

print(x)