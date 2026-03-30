
dir = [
    [-1, 0],[1, 0], [0, -1],[0, 1]
]

def main():
    board = []
    n: int
    max_val :int = -1

    n = int(input())

    for i in range(0, n):
        # line = input().split()
        line = list(input())
        board.append(line)

    for i in range(0, n):
        for j in range(0, n):
            for d in dir:
                dsti = i+d[0]
                dstj = j+d[1]

                if dsti < 0 or dsti >= n or dstj < 0 or dstj >=n:
                    continue

                source = board[i][j]
                dst = board[dsti][dstj]

                board[i][j] = dst
                board[dsti][dstj] = source

                cnt = check_board(board, i, j)
                if max_val < cnt:
                    max_val = cnt

                cnt2 = check_board(board, dsti, dstj)
                if max_val < cnt2:
                    max_val = cnt2
                
                board[dsti][dstj] = dst
                board[i][j] = source
    
    print(max_val)
                

def check_board(board:[], i:int, j:int) -> int:
    max_val = -1                                                 
    count = 1
    for col in range(0, len(board)-1):                           
        if board[i][col] == board[i][col+1]:
            count+=1
        else:                                                    
            count = 1
        if max_val < count:                                      
            max_val = count

    count = 1  # 초기화
    for row in range(0, len(board)-1):
        if board[row][j] == board[row+1][j]:                     
            count+=1
        else:                                                    
            count = 1

        if max_val < count:
            max_val = count
                                                                
    return max_val
# def check_board(board:[], i:int, j:int) -> int:
#     max_val = -1
#     count = 1
#     for col in range(0, len(board)-1):
#         if board[i][col] == board[i][col+1]:
#             count+=1
#         else:
#             count = 1
#         if max_val < count:
#             max_val = count
    
#     for row in range(0, len(board)-1):
#         if board[row][j] == board[row+1][j]:
#             count+=1
#         else:
#             count = 1
        
#         if max_val < count:
#             max_val = count
    
#     return max_val
        

 
if __name__ == "__main__":
    main()