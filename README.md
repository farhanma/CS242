
# GO Game

**An implementation of the Chinese board game GO.**

**This is a final project for Programming Languages course (CS242) given at KAUST in Spring 2014.**

**Implemented using: Google GO Programming Language**

**Development Team:**

*Liam Mencel (liam.mencel@kaust.edu.sa)
*Mohammed Al Farhan (mohammed.farhan@kaust.edu.sa)

## A brief outline of GO

The board game Go is known to have originated in China at least 2500 years ago. 
Since then it has been a popular game worldwide, played by many philosophers, scientists, academics, hobbyists, and even warlords. 
It is a game of territory and strategy, akin to a real military battle.

Why did we choose to implement this game? For one thing, the rules of the game are very simple. 
It requires a lot less attention to complex detail than that of say Chess, and this makes for an elegant programmatic solution. 
Despite its simplicity, the game derives quite complex gameplay, which is algorithmic in nature, 
and through good use of algorithms and data structures it makes for a very interesting project.

### Rules of GO

The general rules of the game can be found in the Wikipedia page: https://en.wikipedia.org/wiki/Go_(game).

Our implementation is based upon the following rules:

*Two players take alternate turns in placing a stone on the board. One player has black stones, the other has white stones
*The main objective is to secure territory. If a player manages to surround a section of the board with his stones, 
then that section effectively belongs to him
*Stones of the same colour which are connected vertically or horizontally form a 'group'. 
This is like a formation of soldiers; a group is always stronger than its individual parts
*But if an opponent surrounds a group with stones of the opposite colour, the inner group dies and vanishes from the board
*Instead of placing a stone, a player has the option to ‘pass’. When both players pass the game ends. 
This would happen when neither of them think they can make a move in their favour
*The scoring system is complex, and in general does not follow an easy formula. 
This is because players usually end the game early when they think one player is a clear winner, and do not play the game till the end. 
The program would not easily be able to evaluate this as well as humans do, 
hence we have decided not to implement an AI-based scoring system

## Installation

You need to have a GO compiler installed on your machine (https://golang.org/doc/install), and the pyGame tool (http://www.pygame.org/download.shtml).

## Usage

First, you need to run the **baduk.go** via typing `go run baduk.go` in your terminal. Then, run **goGraphics.py** via `python2 goGraphics.py`. 
You play through your terminal, following the instructions and the pyGame will simulate your movement on the X window screen.   
