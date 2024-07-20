#!/bin/python3
# GIVES A RECURSION ERROR WHEN EXECUTED ???

board = {
    "top_left": " ", "top_mid": " ", "top_right": " ",
    "mid_left": " ", "mid_mid": " ", "mid_right": " ",
    "low_left": " ", "low_mid": " ", "low_right": " "
}

def printBoard(board):
    print(board["top_left"] + "|" + board["top_mid"] + "|" + board["top_right"])

    print("-+-+-")

    print(board["mid_left"] + "|" + board["mid_mid"] + "|" + board["mid_right"])

    print("-+-+-")

    print(board["low_left"] + "|" + board["low_mid"] + "|" + board["low_right"])

    turn = "X"

    for move in range(9):
        printBoard(board)
        print("Turn for " + turn + ". Which move modern man?")
        move = input()

        board[move] = turn

        if turn == "X":
            turn = "O"
        else:
            turn = "X"

printBoard(board)


