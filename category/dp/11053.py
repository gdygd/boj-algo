import sys

def main():
    n = int(input())
    dp = [0] * 1001
    # numlist = []

    # arr = list(map(int, sys.stdin.readline().split()))
    # numlist = list(map(int, sys.stdin.readline().split()))
    # arr = list(map(int, input().split()))
    numlist = list(map(int, input().split()))

    ans = 0

    for i in range(1, n):
        num1 = numlist[i]
        for j in range(i-1, -1, -1):
            num2 = numlist[j]

            if num2 < num1:
                if dp[i] > dp[j]:
                    continue
                dp[i] = dp[j] + 1
                if ans < dp[i]:
                    ans=dp[i]


    print(ans+1)


if __name__ == "__main__":
    main()