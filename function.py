def mutil_table():
    """打印九九乘法表"""
    row = 1
    while row <= 9:
        col = 1
        while col <= row:
            print("%d*%d=%d" % (col, row, col * row), end="\t")
            col += 1
        print("")
        row += 1


def say_hello():
    print("hello...")


def sum_2_num(num1, num2):
    """
     求和
    :param num1:
    :param num2:
    :return:
    """
    print("sum: %d" % (num1 + num2))
    say_hello()
    return num1 + num2


say_hello()
sum_2_num(12, 13)
mutil_table()
