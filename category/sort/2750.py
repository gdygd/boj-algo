import sys

def main():
    n = int(sys.stdin.readline())
    numlist = []

    for i in range(0, n):
        numlist.append(int(sys.stdin.readline()))

    numlist.sort()

    for n in numlist:
        print(n)

    
    
if __name__ == "__main__":
    main()