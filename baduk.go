package main

import (
    "fmt" 
    "strconv"
    "os"
    "bufio"
)

//Array containment check
func contains(list []int, elem int) bool { 
    for _, t := range list { 
        if t == elem { 
            return true 
        } 
    } 
    return false 
} 

//Function to convert lists to strings
func listToString(list []int) string {
    str := ""
    for _, v := range(list) {
        str += fmt.Sprintf("%d", v)
    }
    return str
}

//Board struct
type Board struct {
    size int
    cells []int
    newGroupID int
    cellToGroupList []map[int]int
    groupToCells map[int][]int
    groupToLiberties map[int][]int
}

//Board constructor
func createBoard(n int) *Board {
    b := &Board{size:n}
    b.clear()
    return b
}

func (b *Board) clear() {
    n := b.size
    b.cells = make([]int, n*n)
    groupMapW:= make(map[int]int)
    groupMapB:= make(map[int]int)
    b.cellToGroupList = []map[int]int{groupMapW, groupMapB}
    b.groupToCells = make(map[int][]int)
    b.groupToLiberties = make(map[int][]int)
}

func (b *Board) copy() *Board {
	z := new(Board)
	z.size = b.size
	z.clear()
	copy(z.cells, b.cells)
	z.newGroupID = b.newGroupID
	for k, v := range b.cellToGroupList[0] {
		z.cellToGroupList[0][k] = v
	}
	for k, v := range b.cellToGroupList[1] {
		z.cellToGroupList[1][k] = v
	}
	for k, v := range b.groupToCells {
		copy(z.groupToCells[k], v)
	}
	for k, v := range b.groupToLiberties {
		copy(z.groupToLiberties[k], v)
	}
	return z
}

func (b *Board) prettyPrint() {
    for i := range b.cells {
        if(i % b.size == 0) {
			s := ""
			if((i/b.size)<9) {
				s = " "
			}
			fmt.Printf("\n%s%d | ", s, (i/b.size)+1)
        }
        fmt.Printf("%d",b.cells[i])
    }
    fmt.Println()
}

func (b *Board) putCellContents(cell, colour int) {
    b.cells[cell] = colour + 1
    b.cellToGroupList[colour][cell] = b.newGroupID
    b.groupToCells[b.newGroupID] = []int{cell}
	b.groupToLiberties[b.newGroupID] = make([]int, 0, b.size * b.size)
    nbrList := b.getNbrs(cell)
    matchingNbrlist := make([]int, 1, 5)
    matchingNbrlist[0] = b.newGroupID
    for _, nbr := range(nbrList) {
        if (b.getCellContents(nbr) == colour+1) {
            elem := b.cellToGroupList[colour][nbr]
            if (! contains(matchingNbrlist, elem)) {
                matchingNbrlist = append(matchingNbrlist, elem)
            }
        }
    }
    b.newGroupID ++
    if(len(matchingNbrlist) > 1) {
        b.mergeGroups(matchingNbrlist, cell)
    }
    b.updateLiberties()
}

func (b *Board) mergeGroups(list []int, bridgeCell int) {
    cellList := make([]int, 0, b.size * b.size)
    colour := b.getCellContents(bridgeCell)-1
    for _, id := range(list) {
        for _, cell := range(b.groupToCells[id]) {
            cellList = append(cellList, cell)
            b.cellToGroupList[colour][cell] = b.newGroupID
        }
        delete(b.groupToCells, id)
		delete(b.groupToLiberties, id)
    }
    b.groupToCells[b.newGroupID] = cellList
	b.groupToLiberties[b.newGroupID] = make([]int, 0, b.size * b.size)
    b.newGroupID ++
	b.updateLiberties()
}

func (b *Board) getNbrs(cell int) []int {
    nbrList := make([]int, 0, 4)
    //right
    if ((cell + 1)%b.size != 0) {
        nbrList = append(nbrList, cell+1)
    }
    //left
    if (cell%b.size != 0) {
        nbrList = append(nbrList, cell-1)
    }
    //up
    if (cell  >= b.size) {
        nbrList = append(nbrList, cell - b.size)
    }
    //down
    if (cell < ((b.size * b.size) - b.size)) {
        nbrList = append(nbrList, cell + b.size)
    }
    return nbrList
}

func (b *Board) getCellContents(cell int) int {
    return b.cells[cell]
}

func (b *Board) updateLiberties() {
	b.groupToLiberties = make(map[int][]int)
	for cell, colour := range(b.cells) {
		if(colour == 0) {
			for _, nbr := range(b.getNbrs(cell)) {
				nbrColour := b.cells[nbr]
				if(nbrColour != 0 && !contains(b.groupToLiberties[b.cellToGroupList[nbrColour-1][nbr]], cell)) {
					b.groupToLiberties[b.cellToGroupList[nbrColour-1][nbr]] = append(b.groupToLiberties[b.cellToGroupList[nbrColour-1][nbr]], cell)
				}
			}
		}
	}
}

func (b *Board) destroyGroup(id int) {
    colour := b.getCellContents(b.groupToCells[id][0])-1
    for _, v := range(b.groupToCells[id]) {
        b.cells[v] = 0
        delete(b.cellToGroupList[colour], v)
    }
    delete(b.groupToCells, id)
    delete(b.groupToLiberties, id)
    b.updateLiberties()
}

func (b *Board) filePrint() {
    output, outputError := os.OpenFile("board.txt", os.O_WRONLY|os.O_CREATE, 0666)
    if outputError != nil {
        fmt.Printf("File couldn't be opened ...\n")
        return
    }
    defer output.Close()
    outputWriter := bufio.NewWriter(output)
    for i := range b.cells {
        if i != 0 {
            if(i % b.size == 0) {
                outputWriter.WriteString("\n")
            }
        }
        outputWriter.WriteString(strconv.Itoa(b.cells[i]))
        outputWriter.WriteString(" ")
    }
    outputWriter.WriteString("\n")
    outputWriter.Flush()
    output.Close()
}


//Game struct
type Game struct {
    board *Board
    history map[string]bool
    whoseTurn int
    aiColour int
}

//Game constructor
func createGame(n, aiColour int) *Game {
    g := &Game{whoseTurn:1}
    g.board = createBoard(n)
    g.history = make(map[string]bool)
    g.history[listToString(g.board.cells)] = true
    g.aiColour = aiColour
    return g
}

func (g *Game) run() {
    g.board.prettyPrint()
    passCounter := 0
    for {
        if(g.aiColour != g.whoseTurn) {
            if(g.doPlayerMove()) {
                passCounter = 0
            } else {
				fmt.Printf("Player %d passes\n", g.whoseTurn)
                passCounter++
            }
        } else {
			if(g.doAiMove()) {
                passCounter = 0
            } else {
				fmt.Println("Computer player passes")
                passCounter++
            }
		}
        if (passCounter == 2) {
			fmt.Println("Both players have passed. Game is over.")
            break
        }
        g.whoseTurn = 3 - g.whoseTurn
    }
}

func (g *Game) doPlayerMove() bool {
    var x,y int
    var d string    
    for {
        fmt.Printf("Player %d: Do you want to place a counter? ", g.whoseTurn)
        fmt.Printf("Enter y or n: ")
        fmt.Scan(&d)
        if(d == "y"){
            fmt.Printf("Enter a row number from 1 to %d: ", g.board.size)
            fmt.Scan(&y)
            fmt.Printf("Enter a column number from 1 to %d: ", g.board.size)
            fmt.Scan(&x)
            cell := (x-1)+((y-1)*g.board.size)
            if(g.isValidMove(cell)) {
				g.board.prettyPrint()
				g.board.filePrint()
                return true
            }
        } else if(d == "n"){
            return false
        } else {
            fmt.Println("Invalid input, please try again.")
        }
    }
    return false
}

func (g *Game) isValidMove(cell int) bool {
    if(g.board.cells[cell] != 0) {
		fmt.Println("Illegal move, stone already in this position.")
		 return false
	}
	boardCopy := g.board.copy()
	boardCopy.putCellContents(cell, g.whoseTurn-1)
	for _, nbr := range(boardCopy.getNbrs(cell)) {
		if(boardCopy.cells[nbr] == 3-g.whoseTurn) {
			if(len(boardCopy.groupToLiberties[boardCopy.cellToGroupList[(3-g.whoseTurn)-1][nbr]]) == 0) {
				boardCopy.destroyGroup(boardCopy.cellToGroupList[(3-g.whoseTurn)-1][nbr])
			}
		}
	}
	if(len(boardCopy.groupToLiberties[boardCopy.cellToGroupList[g.whoseTurn-1][cell]]) == 0) {
		fmt.Println("Illegal move, placed stone has no liberties.")
		return false
	}
	_, ok := g.history[listToString(boardCopy.cells)]
	if(ok) {
		fmt.Println("Illegal move, violates ko rule.")
		return false
	}
	g.board = boardCopy
	g.history[listToString(g.board.cells)] = true
    return true
}

func (g *Game) doAiMove() bool {
	return false
}

func main() {
    fmt.Println("Program compiled. Executing...")
    var b = createBoard(9)
    b.putCellContents(1, 0)
    b.putCellContents(9, 0)
    b.putCellContents(19, 0)
	b.putCellContents(11, 0)
    b.putCellContents(2, 1)
    b.putCellContents(12, 1)
    b.putCellContents(20, 1)
	b.putCellContents(10, 1)
    b.prettyPrint()
    var g = createGame(9, -1)
    g.run()
}