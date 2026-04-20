
class Que:
    def __init__(self):
        self.arr = []
    
    def push(self, pos):
        self.arr.append(pos)

    def pop(self):
        if len(self.arr)==0:
            return None, False
        
        return self.arr.pop(0), True
    
    def size(self):
        return len(self.arr)

visitcnt = [[0] * 101 for _ in range(101)]
visit = [[0] * 101 for _ in range(101)]


def bfs(q:Que, maps:[], n:int, m:int):
    dir = [
        [-1, 0],[1, 0],[0, -1],[0, 1]
    ]

    while q.size() > 0:
        pos, ok = q.pop()
        if not ok:
            break

        if pos[0] == n and pos[1] == m:
            print(visitcnt[pos[0]][pos[1]])

        for d in dir:
            ni, nj = pos[0] + d[0], pos[1] + d[1]

            if ni < 0 or ni > n or nj < 0 or nj > m:
                continue

            if maps[ni][nj] != 1:
                continue

            if visit[ni][nj] == 1:
                continue

            visit[ni][nj] = 1
            curcnt = visitcnt[pos[0]][pos[1]]
            if curcnt > visitcnt[ni][nj]:
                visitcnt[ni][nj] = curcnt+1
                q.push([ni,nj])

def main():
    q = Que()
    n, m = map(int, input().split())

    maps = []
    for i in range(n):
        # row = list(map(int(input())))
        row = list(map(int, input()))

        maps.append(row)
    
    visit[0][0] = 1
    visitcnt[0][0] = 1
    q.push([0,0])

    bfs(q, maps, n-1, m-1)

if __name__=="__main__":
    main()