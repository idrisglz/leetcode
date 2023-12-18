def main():
    print(climbStairs(5))


def climbStairs(n):
    prev = 1
    prev2 = 0

    for i in range(1, n+1):
        curr = prev + prev2
        prev2 = prev
        prev = curr
    return prev


main()
