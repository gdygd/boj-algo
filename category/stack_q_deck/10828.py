import sys

reader  = sys.stdin.readline
write = sys.stdout.write

class Stack:
    def __init__(self):
        self.arr = []
    
    def push(self, x):
        self.arr.append(x)
    
    def pop(self) -> int:
        if not self.arr:
            return -1
        return self.arr.pop()
    
    def size(self) -> int:
        return len(self.arr)
    
    def empty(self) -> int:
        if not self.arr:
            return 1
        else:
            return 0
        
    def top(self) -> int:
        if not self.arr:
            return -1
        return self.arr[-1]

def main():
    st = Stack()

    cmd_cnt = int(reader())
    rst = []

    for _ in range(cmd_cnt):
        cmd = reader().split()
        
        if cmd[0] == "push":
            st.push(int(cmd[1]))
        elif cmd[0] == "pop":
            rst.append((st.pop()))
        elif cmd[0] == "size":
            rst.append(st.size())
        elif cmd[0] == "empty":
            rst.append(st.empty())
        elif cmd[0] =="top":
            rst.append(st.top())

    #  for r in rst:    
    #     write(str(r) + "\n")
    for r in rst:
        write(str(r))
        


if __name__ == "__main__":
    main()
