import (
    "fmt"
)

func minimumRecolors(blocks string, k int) int {
    looping := len(blocks) - k + 1
    minimumW := len(blocks)
    for i:=0; i< looping; i++ {
        countW:=strings.Count(blocks[i:i+k], "W")
        if countW < minimumW {
            minimumW = countW
        }
    }
    return minimumW
}