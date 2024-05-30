#!/bin/python3

"""
File Detection Using Python. 
"""
import os

path = "/home"

if os.path.exists(path):
    print("That location exists.")

    if os.path.isfile(path):
        print("That is a file.")
    else:
        print("That is a directory.")
else:
    print("That location does not exist.")


