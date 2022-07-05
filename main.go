package main

import (
	//  "github.com/sblackstone/go-chess/boardstate"
	//"github.com/sblackstone/go-chess/bitopts"
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewScanner(os.Stdin)
	for {
		buf.Scan()
		fmt.Println(buf.Text())
	}

	// b := boardstate.Initial();
	// b.Print(65)
	// fmt.Println()
	// //bitopts.Print(18446462598732840960, 0);
	// info := unsafe.Sizeof(b)
	// fmt.Printf("%v", info)

}
