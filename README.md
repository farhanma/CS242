
# GO Game

This project is an implementation of the Chinese board game GO.
This was a final project for the graduate-level course CS242: Programming Languages, which was given at KAUST in Spring semester 2014.
The code is written by Google GO Programming Language.

**Development Team:**

* Liam Mencel (liam.mencel@kaust.edu.sa)
* Mohammed Al Farhan (mohammed.farhan@kaust.edu.sa)

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

* Two players take alternate turns in placing a stone on the board. One player has black stones, the other has white stones
* The main objective is to secure territory. If a player manages to surround a section of the board with his stones, 
then that section effectively belongs to him
* Stones of the same colour which are connected vertically or horizontally form a 'group'. 
This is like a formation of soldiers; a group is always stronger than its individual parts
* But if an opponent surrounds a group with stones of the opposite colour, the inner group dies and vanishes from the board
* Instead of placing a stone, a player has the option to ‘pass’. When both players pass the game ends. 
This would happen when neither of them think they can make a move in their favour
* The scoring system is complex, and in general does not follow an easy formula. 
This is because players usually end the game early when they think one player is a clear winner, and do not play the game till the end. 
The program would not easily be able to evaluate this as well as humans do, 
hence we have decided not to implement an AI-based scoring system

## Implementation Details

* To take advantage of the GO language's strong support for Interfaces, we use an object-oriented approach for handling the game. 
The entire state of the board is stored in a struct called **Board**
* The most important Board variable is the board layout, stored in a Slice named **cell**, which is a unique feature of GO. 
It operates in the same way as an array, holding integers to represent which coloured stone sits in the corresponding position. 
Cell contents can be accessed in constant amortised time, through the Board member functions
* The Board struct holds various 'maps' implemented as hash tables. These maps keep track of the groups of stones and their locations. 
The reason for this is that it is better to store the stones as groups rather than individual coordinates. 
This makes it easier to destroy a group as a whole when it is captured
* Board includes various functions which control and update the maps each time a stone is placed
* Board has a **filePrint** function, which stores the current state of the board in a text file. 
We used an external Python library called PyGame to copy the data and render a graphical display. It updates live as the user plays the game
* In addition to Board, we implement a Game interface to control the flow of the game. It has a pointer to the Board variable
* Game has a function called `doPlayerMove()`. This function asks the user to input a move. 
When a move is entered, further functions simulate the move on a dummy board to test whether it would break any rules. If not, then it allows the move
* Game's main loop function is `run()`. It repeatedly calls `doPlayerMove()` until it receives two successive passes, 
indicating that both players have opted to end the game

## Installation Requirements

* GO compiler installed on your machine (https://golang.org/doc/install)
* Python PyGame tool (http://www.pygame.org/download.shtml)

## Execution

* Run **baduk.go** using the terminal: `go run baduk.go` 
* Run **goGraphics.py** using different terminal window: `python2 goGraphics.py` 

## How to Play the Game?

* Use the terminal window for GO to play the game following the instructions give at each step on the screen 
* Use the pyGame tool to visualize your game movements on the X window screen 

## References

* Balbaert, Ivo, *The Way to Go: A Thorough Introduction To The Go Programming Language*, iUniverse, 2012, pp. 1-35
* *The GO Programming Language* (http://www.golang.org)
