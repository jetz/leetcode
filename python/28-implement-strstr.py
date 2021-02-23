class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        if needle == '':
            return 0
        for i in range(0, len(haystack) - len(needle) + 1):
            k, j = 0, i
            while k < len(needle):
                if haystack[j] != needle[k]:
                    break
                else:
                    j += 1
                    k += 1
            else:
                return i

        return -1


if __name__ == "__main__":
    s = Solution()
    haystack, needle = "hello", "lo"
    print(s.strStr(haystack, needle))
    haystack, needle = "aaaaa", "bba"
    print(s.strStr(haystack, needle))
