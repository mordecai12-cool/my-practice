x = "awesome"

# def myfunc():
#     print("python is " + x)

# myfunc()   

#If you use the global keyword, the variable belongs to the global scope:
def myfunc():
    global x
    x = "fantastic"

myfunc()
print("python is " + x)    