
class Que:
    def __init__(self):
        self.arr = []

    def push(self, pos):
        self.arr.append(pos)

    def pop(self):
        if len(self.arr) < 1:
            return None, False
        
        f = self.arr.pop()
        return f, True
    
    def size(self) -> int:
        return len(self.arr)
    
visit = [[0] * 50 for _ in range(50)]

def bfs3(q, maps, n, m):
    dir = [(-1, 0), (1, 0), (0, -1), (0, 1)]

    while q.size() > 0:
        pos, ok = q.pop()
        if not ok:
            break
            
        for d in dir:
            ni = pos[0] + d[0]
            nj = pos[1] + d[1]

            if ni < 0 or ni >= n or nj < 0 or nj >= m:
                continue

            if visit[ni][nj] == 1:
                continue
                
            if maps[ni][nj] == 0:
                continue

            visit[ni][nj] = 1
            q.push([ni, nj])


def main():
    global visit
    casecnt = int(input())
    rst = []
    q = Que()

    for _ in range(casecnt):
        m, n, cnt = map(int, input().split())

        maps = [[0]*50 for _ in range(50)]
        visit = [[0]*50 for _ in range(50)]
        
        warmcnt = 0

        for _ in range(cnt):
            mi, ni = map(int, input().split())
            maps[ni][mi] = 1

        for i in range(n):
            for j in range(m):
                if visit[i][j] == 0:
                    if maps[i][j] == 0:
                        continue

                    visit[i][j] = 1
                    q.push([i, j])

                    bfs3(q, maps, n, m)
                    warmcnt += 1
        
        rst.append(warmcnt)
    
    for r in rst:
        print(r)

if __name__ == "__main__":
    main()