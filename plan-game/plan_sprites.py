import random
import pygame

SCREEN_RECT = pygame.Rect(0, 0, 480, 700)
FRAME_PER_SEC = 60
CREATE_ENEMY_EVENT = pygame.USEREVENT
HERO_FIER_EVENT = pygame.USEREVENT + 1


class GameSprite(pygame.sprite.Sprite):
    def __init__(self, image_name, speed=1):
        super().__init__()
        self.image = pygame.image.load(image_name)
        self.rect = self.image.get_rect()
        self.speed = speed

    def update(self):
        # 在屏幕的垂直方向上移动
        self.rect.y += self.speed


class Background(GameSprite):
    def __init__(self, is_alt=False):
        super().__init__("../images/background.png")
        if is_alt:
            self.rect.y = -self.rect.height

    def update(self):
        #  1.调用父类
        super().update()
        # 判定是否移出屏幕，如果超出将图象设置到屏幕上方
        if self.rect.y >= SCREEN_RECT.height:
            self.rect.y = -self.rect.height


class Enemy(GameSprite):
    def __init__(self):
        super().__init__("../images/enemy1.png")
        self.speed = random.randint(1, 3)
        self.rect.bottom = 0
        max_x = SCREEN_RECT.width - self.rect.width
        self.rect.x = random.randint(0, max_x)

    def update(self):
        # 调用父类
        super().update()
        # 2飞出屏幕 需要删除
        if self.rect.y >= SCREEN_RECT.height:
            print("飞出屏幕 需要删除 ")
            self.kill()

    def __del__(self):
        print(self.rect)


class Hero(GameSprite):
    def __init__(self):
        super().__init__("../images/me1.png", 0)
        # 初始化英雄位置
        self.rect.centerx = SCREEN_RECT.centerx
        self.rect.bottom = SCREEN_RECT.bottom - 20
        # 创建子弹的精灵组
        self.bullets = pygame.sprite.Group()

    def update(self):
        # 英雄在水平方向移动
        self.rect.x = self.rect.x + self.speed
        if self.rect.x < 0:
            self.rect.x = 0
        elif self.rect.right > SCREEN_RECT.right:
            self.rect.right = SCREEN_RECT.right

    def fire(self):
        print("fire ...")

        for i in (0, 1, 2):
            # 创建子弹
            bullet = Bullet()
            # 设置子弹的位置
            bullet.rect.bottom = self.rect.y - i * 20
            bullet.rect.centerx = self.rect.centerx
            # 添加精灵到 精灵组
            self.bullets.add(bullet)



class Bullet(GameSprite):
    """子弹精灵"""

    def __init__(self):
        super().__init__("../images/bullet1.png", -4)


    def update(self):
        super().update()
        # 判定子弹是否飞出屏幕
        if self.rect.bottom < 0:
            self.kill()



    def __del__(self):
        print("bullet del...")
