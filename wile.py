import function

function.mutil_table()
function.say_hello()

list_a = ["xxx", function]
print(list_a)
list_a.append("jjj")
list_a[1] = "2222"

list_a.remove("xxx")
list_a.sort()
print(list_a.count("jjj"))

list_a.extend(["j", "k"])

print("x" in list_a)
print("k" in list_a)
print(list_a)

list_tuple = tuple(list_a)
print(type(list_tuple))

xm = {"name": "11"}
print(xm['name'])
xm["age"] = 18
xm["sex"] = "男"
for key in xm:
    print("%s-%s" % (key, xm[key]))
xm.pop("sex")

print(xm)
print(max(xm))
print(min(xm))
h = "hello .."
print(h.find('e'))
print(h.startswith('h'))
print(h.endswith('e'))
print(len(h))
print(h[:])
