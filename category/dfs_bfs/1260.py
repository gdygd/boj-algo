
class Que:
    def __init__(self):
        self.arr = []

    def push(self, n):
        self.arr.append(n)

    def pop(self):
        if len(self.arr) < 1:
            return None, False
        
        f = self.arr.pop(0)
        return f, True
    
    def size(self) -> int:
        return len(self.arr)
    

graph = [[0]*1001 for _ in range(1001)]
visit = [0]*1001

def dfs(v):
    global visit
    print(v, end=" ")
    for v2 in range(0, 1001):
        if graph[v][v2] == 1:
            if visit[v2] == 0:
                visit[v2] = 1

                dfs(v2)

def bfs(start):
    q = Que()
    q.push(start)

    global visit

    visit = [0] * 1001
    visit[start] = 1

    while q.size() > 0:
        v, ok = q.pop()
        if not ok:
            break

        print(v, end=" ")

        for v2 in range(1, 1001):
            if graph[v][v2] == 1:
                if visit[v2]==0:
                    visit[v2] = 1
                    q.push(v2)

def main():
    vcnt, lcnt, start = map(int, input().split())

    for i in range(0, lcnt):
        v1, v2 = map(int, input().split())

        graph[v1][v2] = 1
        graph[v2][v1] = 1

    visit[start] = 1

    visit[start] = 1
    dfs(start)
    print()

    bfs(start)
    print()
    

if __name__ == "__main__":
    main()