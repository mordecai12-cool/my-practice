#Boolean Values

# a = 200
# b = 330

# if b > a:
#     print("b is greater than a")
# else:
#     print("b is not greater than a")

#When you compare two values, the expression is evaluated and Python returns the Boolean answer:
print(10 > 9)
print(10 == 9)
print(10 < 9)

#Evaluate a string and a number:
class myclass():
    def __len__(self):
        return 0

myobj = myclass()
print(bool(myobj)) # False

#Print "YES!" if the function returns True, otherwise print "NO!":
def myFunction():
    return True

if myFunction():
    print("YES!")
else:
    print("NO!")

#Check if an object is an integer or not:
x = 200
print(isinstance(x, int))