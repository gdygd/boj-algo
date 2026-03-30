

def main():
    n, m = map(int, input().split())
    cards = []

    cards = list(map(int, input().split()))
    
    gab = 300000
    rst = 0
    for a in range(0, n-2):
        for b in range(a+1, n-1):
            for c in range(b+1, n):
                a1 = cards[a]
                b1 = cards[b]
                c1 = cards[c]

                sum = a1 + b1 + c1
                if m-sum >= 0:
                    if gab > (m - sum):
                        gab = m-sum
                        rst = sum

                    
    print(rst)


if __name__ == "__main__":
    main()