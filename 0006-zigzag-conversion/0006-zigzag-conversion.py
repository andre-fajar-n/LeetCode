class Solution:
    def convert(self, s: str, numRows: int) -> str:
        temp = [[] for i in range(numRows)]
        position = 0
        arrow = "up"
        for i, v in enumerate(s):
            print("CEK", i, v, position, numRows)
            temp[position].append(v)
            if numRows > 1:
                if arrow == "up":
                    if position == numRows - 1:
                        position -= 1
                        arrow = "down"
                    else:
                        position += 1
                elif arrow == "down":
                    if position == 0:
                        position += 1
                        arrow = "up"
                    else:
                        position -= 1
            
        output = ""
        for v in temp:
            output += "".join(v)
        
        return output