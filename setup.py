from distutils.core import setup

setup(name="hm_message",
      version="1.0",  # 版本
      description="消息发送接收模拟",
      long_description="消息发送接收模拟",
      author="zhangyangyang",
      author_email="254856090@qq.com",
      url="www.baidu.com",
      py_modules=[
          "tools.send",
          "tools.recieve"
      ]


      )