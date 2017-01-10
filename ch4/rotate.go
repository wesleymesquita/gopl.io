package main

import ( "fmt"
	 "strings"
	"strconv"
	"os"
	"bufio")

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inp := scanner.Text()
	input := strings.Split(inp, " ")
	iinput := make([]int, len(input))
	for i,v := range input{
		iinput[i], _ = strconv.Atoi(v)
	}
	r := rotate(iinput, 2)
	fmt.Printf("%v\n", r)
}


func rotate(inp []int, i int) []int{
	ret := make([]int,len(inp))
	copy(ret,inp[i:])
	copy(ret[len(inp)-i:], inp[0:i])
	return ret
}

