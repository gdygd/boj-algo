
class Que:
    def __init__(self):
        self.arr=[]

    def push(self, pos):
        self.arr.append(pos)

    def pop(self):
        if len(self.arr) == 0:
            return None, False
        
        f = self.arr.pop(0)
        return f, True
    
    def size(self):
        return len(self.arr)
    

visit = [[0]*25 for _ in range(25)]

def bfs(q, n, maps):
    dir = [
        [-1,0], [1,0], [0,-1], [0,1]
        ]

    cnt = 0
    while q.size()>0:
        pos, ok = q.pop()
        if not ok:
            break

        for d in dir:
            ni, nj = pos[0] + d[0], pos[1] + d[1]
            if ni < 0 or ni >= n or nj < 0 or nj >= n:
                continue

            if visit[ni][nj] == 0 and maps[ni][nj] == '1':
                visit[ni][nj] = 1

                q.push([ni,nj])
                cnt += 1
    
    return cnt

def main():
    q = Que()
    n = int(input())
    maps =  [[0]*n for _ in range(n)]

    for i in range(n):
        maps[i] = input()

    compcnt = 0
    rst = []
    for i in range(0,n):
        for j in  range(0, n):
            if maps[i][j] == '1' and visit[i][j] == 0:
                q.push([i,j])
                visit[i][j] = 1
                compcnt+=1
                cnt = bfs(q, n, maps)
                rst.append(cnt+1)


    print(compcnt)
    rst.sort()
    for r in rst:
        print(r)



if __name__ == "__main__":
    main()