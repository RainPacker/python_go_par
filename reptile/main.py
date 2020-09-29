# -*- coding=utf8
from bs4 import BeautifulSoup
import re
import urllib.request
import urllib.response
import urllib.error

import ssl

context = ssl._create_unverified_context()

SRC_REG = re.compile(r'<div class="post">.*<a href="(.*?)".*>', re.S)
# IMG_REG = re.compile(r'<div class="post">\n<a href=".*".*>\n<img.*src="(.*?)"', re.S)
IMG_REG = re.compile(r'<div class="post">\n<a href=".*" target="_blank">.<img.*src="(.*?)"', re.S)
NAME_REG = re.compile(r'<div class="title">\n<a.*target="_blank">(.*?)</a>', re.S)


def main():
    print("==run==")
    # url
    base_url = "https://www.douban.com/doulist/968362/?start="
    ask_url(base_url)
    # 获取 网页数据
    data_list = get_data(base_url)
    # 解析数据
    # 保存数据
    pass


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
    for i in range(0, 1):
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
            info = sp.find_all("div", class_="title")[0]
            name = info.text
            imgs = re.findall(IMG_REG, item_str)
            # name = re.findall(NAME_REG,item_str)[0]
            name = str(name).replace("\n", " ").strip()
            print("movieName", name)
            image_url = ""
            if len(imgs) > 0:
                image_url = imgs[0]

            info = {
                "movie_name": name,
                "img": image_url,
                "url": movie_url
            }
            print("imageUrl", image_url)
            data_list.append(info)
    print(len(data_list))
    print(data_list)


if __name__ == '__main__':
    main()
