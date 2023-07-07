class Solution:
    def findWords(self, words: List[str]) -> List[str]:
        row1 = 'qwertyuiop'
        row2 = 'asdfghjkl'
        row3 = 'zxcvbnm'
        hasil = []
        for kata in words:
            count_row1, count_row2, count_row3 = 0, 0, 0
            kata_lower = kata.lower()
            
            for i in range(len(kata)):
                if kata_lower[i] in row1:
                    count_row1 += 1
                if kata_lower[i] in row2:
                    count_row2 += 1
                if kata_lower[i] in row3:
                    count_row3 += 1
                
            if count_row1 == len(kata) or count_row2 == len(kata) or count_row3 == len(kata):
                hasil.append(kata)
        return hasil