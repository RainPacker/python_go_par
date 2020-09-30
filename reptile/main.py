# -*- coding=utf8
from bs4 import BeautifulSoup
import re
import urllib.request
import urllib.response
import urllib.error
import xlwt
from fileDownload import requestImg

import ssl

context = ssl._create_unverified_context()

SRC_REG = re.compile(r'<div class="post">.*<a href="(.*?)".*>', re.S)
# IMG_REG = re.compile(r'<div class="post">\n<a href=".*".*>\n<img.*src="(.*?)"', re.S)
IMG_REG = re.compile(r'<div class="post">\n<a href=".*" target="_blank">.<img.*src="(.*?)"', re.S)
NAME_REG = re.compile(r'<div class="title">\n<a.*target="_blank">(.*?)</a>', re.S)


def save_data_excel(data_list):
    workbook = xlwt.Workbook(encoding="utf-8")
    sheet = workbook.add_sheet("sheet1")
    col_name = ["电影名称", "电影图象", "电影详情"]
    # sheet.write(0, 0, "ssss")
    for col in range(0, 3):
        print(col)
        sheet.write(0, col, col_name[col])
    row = 1
    for item in data_list:
        sheet.write(row, 0, item["movie_name"])
        sheet.write(row, 1, item["img"])
        requestImg(item["img"], item["movie_name"])
        sheet.write(row, 2, item["url"])
        row += 1


    workbook.save("movies.xls")


def main():
    print("==run==")
    # url
    base_url = "https://www.douban.com/doulist/968362/?start="
    ask_url(base_url)
    # 获取 网页数据   # 解析数据
    data_list = get_data(base_url)
    print(data_list)
    # 保存数据
    save_data_excel(data_list)


def ask_url(url):
    headers = {
        "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36"
    }
    req = urllib.request.Request(url=url, headers=headers)
    html = ""
    try:
        reponse = urllib.request.urlopen(req, context=context)
        html = reponse.read().decode("utf-8")
    except urllib.error.URLError as e:
        if hasattr(e, "code"):
            print(e.code)
        if hasattr(e, "reason"):
            print(e.reason)

    return html


def get_data(baseUr):
    data_list = []
    for i in range(0, 10):
        url = baseUr + str(i * 25)
        html = ask_url(url)
        soup = BeautifulSoup(html, "html.parser")
        for item in soup.find_all("div", class_="doulist-item"):
            item_str = str(item)
            movie_url = ""
            movie_urls = re.findall(SRC_REG, item_str)
            if len(movie_urls) > 0:
                movie_url = movie_urls[0]
            print("moverUrl", movie_url)
            sp = BeautifulSoup(item_str, "html.parser")

            info = sp.find_all("div", class_="title")
            if len(info) > 0:
                info = info[0]

            post_info = sp.find_all("div", class_="post")
            if len(post_info) > 0:
                post_info = post_info[0]
            try:
                image_url = str(post_info.a.img["src"])
            except Exception as result:
                print(result)
            try:
                name = info.text
            except Exception as result:
                name = ""
                print(result)

            name = str(name).replace("\n", " ").strip()
            print("movieName", name)
            # image_url = ""
            # if len(imgs) > 0:
            #     image_url = imgs[0]

            info = {
                "movie_name": name,
                "img": image_url,
                "url": movie_url
            }
            print("imageUrl", image_url)
            data_list.append(info)
    print(len(data_list))
    return data_list


if __name__ == '__main__':
    main()
