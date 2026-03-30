import sys
input = sys.stdin.readline

class Queue:
    def __init__(self):
        self.arr = []

    def push(self, x):
        self.arr.append(x)
    
    def pop(self) -> int:
        if not self.arr:
            return -1

        return self.arr.pop(0)

    def size(self) -> int:
        return len(self.arr)
    
    def empty(self) -> int:
        if not self.arr:
            return 1
        return 0

    def front(self) -> int:
        if not self.arr:
            return -1
        return self.arr[0]

    def back(self) -> int:
        if not self.arr:
            return -1
        return self.arr[-1]
        
def main():
    q = Queue()
    n = int(input())

    rst = []

    for _ in range(n):
        cmd = input().split()

        if cmd[0] == "push":
            q.push(int(cmd[1]))
        elif cmd[0] == "pop":
            rst.append(q.pop())
        elif cmd[0] == "size":
            rst.append(q.size())
        elif cmd[0] == "empty":
            rst.append(q.empty())
        elif cmd[0] == "front":
            rst.append(q.front())
        elif cmd[0] == "back":
            rst.append(q.back())
    for r in rst:
        print(str(r))
    
if __name__ == "__main__":
    main()