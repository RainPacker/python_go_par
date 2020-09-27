import os

file = open("../pay.html")  # 打开文件
# file.write("<--->")
text = file.read()

print(text)

while True:
    text2 = file.readline()
    if not text2:
        break
    print(text2)

# 关闭文件
file.close()

print(os.listdir())
print(os.getcwd())
