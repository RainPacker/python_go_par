import random
print("###game start#######################################")
player = int(input("请验入你要出的的拳头( 1:石头 2:剪刀 3:布):"))

computer = random.randint(1, 3)

print("电脑出的是%d,玩家出的是%d" % (computer, player))
if ((player == 1 and computer == 2)
        or (player == 2 and computer == 3)
        or (player == 3 and computer == 1)):

    print(" 电脑弱爆了，你赢了")
elif player == computer:
    print("真是心有灵犀，再来一盘")
else:
    print("你输了，不服再来！")
print("###game over#######################################")
