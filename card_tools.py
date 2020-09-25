# 名片列表
card_list = []


def show_menu():
    print("*" * 50)
    print(" 欢迎使用名片管理系统1.0")
    print(" 1. 新增名片")
    print(" 2. 显示全部 ")
    print(" 3. 查询名片")
    print(" 0. 退出")
    print("*" * 50)


def new_card():
    """新增名片"""
    print("-" * 50)
    print("新增名片")
    name_str = input("请输入姓名：")
    phone_str = input("请输入电话：")
    qq_str = input("请输入qq：")
    email_str = input("请输入email：")
    card_dict = {
        "name": name_str,
        "phone": phone_str,
        "qq": qq_str,
        "email": email_str
    }

    print(card_dict)
    card_list.append(card_dict)
    print("添加%s的名片成功" % name_str)


def show_all():
    """显示全部"""
    print("-" * 50)
    print("显示全部")

    # 判断是否有名片
    if len(card_list) == 0:
        print("当前没有名片...")
        return

    # 打印表头
    for name in ["姓名", "电话", "qq", "email"]:
        print(name, end="\t\t")
    print("")
    print("=" * 50)
    # 遍历所有列表
    for card_dic in card_list:
        print("%s\t\t%s\t\t%s\t\t%s" % (card_dic["name"], card_dic["phone"], card_dic["qq"], card_dic["email"]))


def search_card():
    """查询"""
    print("-" * 50)
    print("查询名片")
    find_name = input("查找姓名：")
    for card_dic in card_list:
        if card_dic["name"] == find_name:
            for name in ["姓名", "电话", "qq", "email"]:
                print(name, end="\t\t")
            print("")
            print("=" * 50)
            print("%s\t\t%s\t\t%s\t\t%s" % (card_dic["name"], card_dic["phone"], card_dic["qq"], card_dic["email"]))

            #  修改，删除 dict
            deal(card_dic)
            pass
            break
    else:
        print("没有找到")


def deal(find_dict):

    action_str = input("请选择要执行的菜单，【1】： 修改 【2】：删除 【0】：返回上一级菜单")

    if action_str == "1":

        find_dict["name"] = card_input_info(find_dict["name"], "姓名:")
        find_dict["phone"] = card_input_info(find_dict["phone"], "电话:")
        find_dict["qq"] = card_input_info(find_dict["qq"], "qq:")
        find_dict["email"] = card_input_info(find_dict["email"], "email:")
        print("修改成功")
    elif action_str == "2":
        card_list.remove(find_dict)
        print("删除")


def card_input_info(dict_value, tips_message):
    """

    :param dict_value:
    :param tips_message: 提示
    :return:
    """
    input_str = input(tips_message)
    if len(input_str) > 0:
        return input_str
    else:
        return dict_value
