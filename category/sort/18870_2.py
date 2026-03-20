import sys
from bisect import bisect_left

def main():
    reader = sys.stdin.readline
    reader()

    numlist = list(map(int, reader().split()))
    compnum = []
    check_n = dict()

    for num in numlist:
        if num not in check_n:
            compnum.append(num)
            check_n[num] = True

    compnum.sort()

    for num in numlist:
        rank = bisect_left(compnum, num)
        print(rank, end=" ")

if __name__ == "__main__":
    main()