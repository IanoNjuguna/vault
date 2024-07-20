#!/bin/python3

#pprint module is for pretty printing (cleaner diplay)
import pprint

count = {}

print("This program counts the number of occurrences of each letter in a string.")

print("To use, input a string")
str = input()

for char in str:
    count.setdefault(char, 0)
    count[char] = count[char] + 1

pprint.pprint(count)
