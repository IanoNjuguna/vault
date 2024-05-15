def fibo(n):
    a, b = 0, 1

    while 0 < n:
        print(a, end='')
        a, b = b, (a + b)
        
        print()

print("Enter a Number:")
numbr = input()

fibo(numbr)
