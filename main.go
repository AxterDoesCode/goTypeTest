package main

import (
	"bufio"
	"fmt"
	"os"
)

func main () {
    reader := bufio.NewScanner(os.Stdin)
    wordBuf := ""
    for {
        fmt.Print(wordBuf)
        reader.Scan()
        wordBuf += reader.Text()
    }
}
