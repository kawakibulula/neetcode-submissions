class Solution:
    def isValid(self, s: str) -> bool:
        isClose = {")": "(", "]": "[", "}": "{"}
        stack = []
        for c in s:
            if c not in isClose:
                stack.append(c)
                continue
            if not stack or stack[-1] != isClose[c]:
                return False
            stack.pop()

        return not stack        