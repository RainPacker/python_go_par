import urllib.request
import urllib.parse

#请求一个 get 请求
response = urllib.request.urlopen("http://www.baidu.com")
# print(response.read().decode("utf-8"))
# 请求使用 post
data = bytes(urllib.parse.urlencode({"hello": "world"}), "utf-8")
header = {
    "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.89 Safari/537.36"
}
req = urllib.request.Request(url="http://httpbin.org/post",headers=header, data=data)
rep = urllib.request.urlopen(req)
print(rep.read().decode("utf-8"))