def main():
    solution = Solution()
    print(solution.climbStairs(5))
    print(solution.climbStairs(3))

class Solution:
    cache = {}

    def climbStairs(self, n: int) -> int:
        if n <= 2:
            return n

        if n in self.cache.keys():
            return self.cache[n]

        self.cache[n] = self.climbStairs(n-1) + self.climbStairs(n-2)
        return self.cache[n]



# Different implemnetioan
class Solution2:
    def climbStairs(self, n: int) -> int:
        current, previous = 1, 0

        for i in range(n):
            previous, current = current, previous+current
        return current

main()