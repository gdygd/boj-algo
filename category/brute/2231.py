
def main():
    strn = input()
    inputn = int(strn)

    offset = len(strn) * 9
    num = int(strn)
    rst = 0

    for i in range(max(1, num-offset), num):
        strnum = str(i)
        sum = i
        for digit in strnum:          
            sum += int(digit)
            
        
        if sum == inputn:
            rst = i
            break

    print(rst)
    

if __name__ == "__main__":
    main()