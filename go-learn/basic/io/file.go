page main
import (
"bufio"
"fmt"
"io"
"os"
)

func main() {
	inputFile, err := os.Open("input.dat")
	if err != nil {
	    fmt.Printf("open file err: %v\n", err)
	    return
	}

	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	for {
	   inputString reader
	}
}