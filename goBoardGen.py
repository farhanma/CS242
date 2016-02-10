import random

n = random.randint(3, 25)
mat = [[random.randint(0, 2) for i in range(n)] for j in range(n)]

g = open("board.txt", 'w')

for row in mat:
    string = ""
    for item in row:
        string += str(item) + ' '
    g.write(string + '\n')

g.close()
