class Luhn:
    def __init__(self, card_num: str):
        self.nums = card_num.replace(" ", "")

    def valid(self) -> bool:
        if not self.nums.isdigit() or len(self.nums) < 2:
            return False

        even = len(self.nums) % 2 == 0
        result = 0
        for char in self.nums:
            digit = int(char)
            if even:
                digit *= 2
                if digit > 9:
                    digit -= 9
            even = not even
            result += digit

        return result % 10 == 0
