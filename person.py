class Person:
    def __init__(self, name, age):
        self.name = name
        self.age = age

    @staticmethod
    def run():
        print("..run...")
        pass

    def __str__(self):
        return "name:%s age:%d" % (self.name, self.age)


class Dog(object):
    def __init__(self, name):
        self.name = name

    @classmethod
    def play(cls):
        print("..蹦蹦跳跳。。")


class XtDog(Dog):
    def play(self):
        super().play()
        print("天上飞。。。")


xm = Person("xm", 1)
xm.run()
print(xm)
xt = XtDog("xt")
xt.play()
