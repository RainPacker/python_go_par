def demo1():
    a_int = int(input("请输入整数"))


def demo2():
    return demo1()


def input_password(pwd):
    # 1
    if len(pwd) > 8:
        return pwd
    else:
        print("exception....")
        raise  Exception("密码长度不够。。。")


try:
    demo2()
except Exception as result:
    print(result)


try:
    input_password(input("请输入密码"))
except Exception as result:
    print(result)
