import urllib.request
from http.client import IncompleteRead, RemoteDisconnected
from urllib.error import HTTPError, URLError
import ssl
import os

context = ssl._create_unverified_context()

def requestImg(url, name, num_retries=3):
    img_src = url
    # print(img_src)
    header = {
        'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; WOW64) \
    AppleWebKit/537.36 (KHTML, like Gecko) \
      Chrome/35.0.1916.114 Safari/537.36',
        'Cookie': 'AspxAutoDetectCookieSupport=1'
    }
    # Request类可以使用给定的header访问URL
    req = urllib.request.Request(url=img_src, headers=header)
    try:
        response = urllib.request.urlopen(req, context=context) # 得到访问的网址
        mkdir("./imgs")
        filename = "imgs/"+ name + '.jpg'
        with open(filename, "wb") as f:
            content = response.read() # 获得图片
            f.write(content) # 保存图片
            response.close()
    except HTTPError as e: # HTTP响应异常处理
        print(e.reason)
    except URLError as e: # 一定要放到HTTPError之后，因为它包含了前者
        print(e.reason)
    except IncompleteRead or RemoteDisconnected as e:
        if num_retries == 0: # 重连机制
            return
        else:
            requestImg(url, name, num_retries-1)


def mkdir(path):
    # 引入模块


    # 去除首位空格
    path=path.strip()
    # 去除尾部 \ 符号
    path=path.rstrip("\\")

    # 判断路径是否存在
    # 存在     True
    # 不存在   False
    isExists=os.path.exists(path)

    # 判断结果
    if not isExists:
        # 如果不存在则创建目录
        # 创建目录操作函数
        os.makedirs(path)

        print( path+' 创建成功')
        return True
    else:
        # 如果目录存在则不创建，并提示目录已存在
        # print (path+' 目录已存在')
        return False



if __name__ == '__main__':
    image_url = "https://img3.doubanio.com/view/photo/s_ratio_poster/public/p480747492.jpg"
    requestImg(image_url, "test")