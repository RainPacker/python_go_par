import pygame
from plan_sprites import *


class PlanGame(object):
    """飞机"""

    def __init__(self):
        print("init...")
        pygame.init()
        self.screen = pygame.display.set_mode((SCREEN_RECT.width, SCREEN_RECT.height))
        self.clock = pygame.time.Clock()
        # 创建精灵和精灵组
        self.__create_sprites()
        # 创建定时器
        pygame.time.set_timer(CREATE_ENEMY_EVENT, 1000)
        # 英雄射击定时任务
        pygame.time.set_timer(HERO_FIER_EVENT, 500)
        self.score = 0

    def star_game(self):
        print(" game start")
        while True:
            # 设置 刷新帧率
            self.clock.tick(FRAME_PER_SEC)
            # 2事件监听
            self.__event_handler()
            # 3 碰撞检测
            self.__check_collide()
            # 4 更新绘制精灵组
            self.__update_sprites()
            # 5 更新显示
            pygame.display.update()
            pass

    def __event_handler(self):
        """事件处理"""
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                PlanGame.__game_over()
            elif event.type == CREATE_ENEMY_EVENT:
                print("敌机出场")
                enemy = Enemy()
                self.enemy_groop.add(enemy)
            elif event.type == HERO_FIER_EVENT:
                self.hero.fire()
        keys_pressed = pygame.key.get_pressed()
        if keys_pressed[pygame.K_RIGHT]:
            print("向右移动...")
            self.hero.speed = 2

        elif keys_pressed[pygame.K_LEFT]:
            print("向在移动...")
            self.hero.speed = -2
        else:
            self.hero.speed = 0

    def __check_collide(self):
        """碰撞检测"""
        # 子弹 摧毁 敌机
        enemy_groop = pygame.sprite.groupcollide(self.hero.bullets, self.enemy_groop, True, True)

        self.score += len(enemy_groop) * 10
        # 敌机碰撞英雄
        enemy = pygame.sprite.spritecollide(self.hero, self.enemy_groop, True)
        if len(enemy) > 0:
            self.hero.kill()
            self.__game_over()

    def __update_sprites(self):
        self.back_group.update()
        self.back_group.draw(self.screen)
        self.enemy_groop.update()
        self.enemy_groop.draw(self.screen)
        self.hero_group.update()
        self.hero_group.draw(self.screen)

        # bullet
        self.hero.bullets.update()
        self.hero.bullets.draw(self.screen)

        score_str = "Score:%8d" % self.score
        self.textSurface = self.textFont.render(score_str, True, (255, 255, 255))
        self.screen.blit(self.textSurface, (10, 10))

    @staticmethod
    def __game_over():
        print("游戏结束")
        pygame.quit()
        exit()

    def __create_sprites(self):

        bg1 = Background()
        bg2 = Background(True)
        self.back_group = pygame.sprite.Group(bg1, bg2)
        self.enemy_groop = pygame.sprite.Group()
        # hero
        self.hero = Hero()
        self.hero_group = pygame.sprite.Group(self.hero)
        # 加载字体
        # textFont = pygame.font.Font("../font/font.ttf", 30)
        self.textFont = pygame.font.SysFont('arial', 16)
        self.textSurface = self.textFont.render("Score:00000", True, (255, 255, 255))


if __name__ == '__main__':
    game = PlanGame()
    game.star_game()
