import os, sys
import pygame
import time
from pygame.locals import *

pygame.init()
width = 600
height = 600
screen = pygame.display.set_mode((width, height))
pygame.display.set_caption('Go')

background = pygame.Surface(screen.get_size())
background = background.convert()
background.fill((185,122,87))

screen.blit(background, (0, 0))
pygame.display.flip()

goEmpty = pygame.image.load('goEmpty.bmp')
goBlack = pygame.image.load('goBlack.bmp')
goWhite = pygame.image.load('goWhite.bmp')

oldMat = []

clock = pygame.time.Clock()

done = False
while(not done):
    time.sleep(1)

    for event in pygame.event.get():
        if event.type == QUIT:
            done = True
        if event.type == KEYDOWN and event.key == K_ESCAPE:
            done = True
            
    f = open("board.txt", 'r')
    mat = []
    line = f.readline()
    while line:
        mat.append([int(x) for x in line.split()])
        line = f.readline()
    f.close()

    if mat == oldMat:
        continue

    oldMat = mat
    n = len(mat[0])

    goEmpty1 = pygame.transform.smoothscale(goEmpty, (width/n, height/n))
    goBlack1 = pygame.transform.smoothscale(goBlack, (width/n, height/n))
    goWhite1 = pygame.transform.smoothscale(goWhite, (width/n, height/n))

    screen.blit(background, (0, 0))
    
    for i in range(n):
        for j in range(n):
            tile = goEmpty1
            if mat[i][j] == 1:
                tile = goBlack1
            elif mat[i][j] == 2:
                tile = goWhite1
            screen.blit(tile, (j*width/n, i*height/n))

    pygame.display.flip()
