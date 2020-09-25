import keyword
import  random

print("hello python")
print("hello world")
# 23123123
print(">>>>>>")
var = 1.01 ** 365
x = ">" * 50
re = "加油" * 50
count = 7200 * 12
av = 100000 / 12
print(x)
print(re)
print(count)
print(av)
# price
price = 3.5
num = 6

sum = price * num
print(sum)
var = type(sum)
print(var)
# p = input("please input\n")
# print(p)
print(int("123"))
print(float("123"))
name = "张杨杨"
print("我的名字是%s" % name)
stud_no = 100123456
print("我的学号是%010d" % stud_no)
money_float = 0.0677
print("money: %.2f  %.2f%%" % (money_float, money_float * 100))

print(keyword.kwlist)
if money_float >= 1:
    print(" money >1")
    print(keyword)
else:
    print(money_float)

print("====" * 10)

print("===for====")

for le in re:
    print(" letter:", le)

#
run = random.randint(1, 100)
print(run)



i =1
sum = 0
while i <=100:
    # print(i)
    sum += i
    i = i+1
print(sum, end= "")