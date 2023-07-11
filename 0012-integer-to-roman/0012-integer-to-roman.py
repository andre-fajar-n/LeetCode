class Solution:
    def intToRoman(self, num: int) -> str:
        romawi = ['I', 'V', 'X', 'L', 'C', 'D', 'M']
        angka = [1, 5, 10, 50, 100, 500, 1000]
        lst = []
        for i in range(-1, -(len(romawi)+1), -2):
            jml_romawi = num // angka[i]
            if 5 > jml_romawi > 0:
                if jml_romawi == 4:
                    lst.append(romawi[i])
                    lst.append(romawi[i+1])
                else:
                    for k in range(1, jml_romawi+1):
                        lst.append(romawi[i])
            elif jml_romawi == 5:
                lst.append(romawi[i+1])
            elif 10 > jml_romawi >= 5:
                if jml_romawi == 9:
                    lst.append(romawi[i])
                    lst.append(romawi[i+2])
                else:
                    lst.append(romawi[i+1])
                    for k in range(1, (jml_romawi-5)+1):
                        lst.append(romawi[i])
            num -= jml_romawi*angka[i]
        lst = ''.join(lst)
        return lst