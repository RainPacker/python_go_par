from bs4 import BeautifulSoup


file = open("../pay.html", "rb")
html = file.read()

bs = BeautifulSoup(html, "html.parser")
# print(bs.div)
# print(bs.title.string)
# print(bs.input)

# print(bs.head)
# print(bs.find_all("div"))
t_list = bs.select("#goods_info")
for item in t_list:
    print(item)
    print(type(item))




file.close()
