import pygame
from plan_sprites import *

pygame.init()
screen = pygame.display.set_mode((480, 700))
# 绘制背景 1>load >blit >update
bg = pygame.image.load("../images/background.png")
screen.blit(bg, (0, 0))
# pygame.display.update()

# 绘制飞机
hero = pygame.image.load("../images/me1.png")
screen.blit(hero, (150, 300))
# 统一更新
pygame.display.update()

# 创建时钟对象
clock = pygame.time.Clock()

# 飞机初始位置
hero_rect = pygame.Rect(150, 300, 102, 126)
# 创建敌机
enemy = GameSprite("../images/enemy1.png")
enemy_group = pygame.sprite.Group(enemy)

while True:
    clock.tick(60)

    # 事件
    event_list = pygame.event.get()

    for event in event_list:
        if event.type == pygame.QUIT:
            print("退出游戏...")
            pygame.quit()
            exit()

    hero_rect.y -= 1
    # 判定飞机位置
    if hero_rect.bottom <= 0:
        hero_rect.y = 700

    screen.blit(bg, (0, 0))
    screen.blit(hero, hero_rect)

    enemy_group.update()
    enemy_group.draw(screen)

    pygame.display.update()
