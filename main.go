package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Slice filter and find the index
func findIndexFilter(vs [][]string, f func([]string) bool) []int {
	vsf := make([]int, 0)
	for i, v := range vs {
		if f(v) {
			vsf = append(vsf, i)
		}
	}
	return vsf
}

func checkPartOfIsland(i, j int, islands *[][]string) {
	ii := strconv.Itoa(i)
	jj := strconv.Itoa(j)
	coddin := (ii + "" + jj)
	coddinInt, err := strconv.Atoi(coddin)
	if err != nil {
		panic(err)
	}
	if len(*islands) != 0 {
		filtered := findIndexFilter(*islands, func(s []string) bool {
			co := s
			result := false
			for _, c := range co {
				d, err := strconv.Atoi(c)
				if err != nil {
					panic(err)
				}
				compare := (coddinInt - d)
				if compare == 10 || compare == 01 {
					result = true
					break
				}
			}
			return result
		})
		if len(filtered) != 0 {

			in0 := filtered[0]

			// If filtered index more then 1
			if len(filtered) > 1 {
				in1 := filtered[1]

				//Merge the two index to the first index
				(*islands)[in1] = append((*islands)[in1], codi)
				(*islands)[in0] = append((*islands)[in0], (*islands)[in1]...)

				//Remove index position
				copy((*islands)[in1:], (*islands)[in1+1:])
				(*islands)[len((*islands))-1] = []string{""}
				(*islands) = (*islands)[:len((*islands))-1]

			} else {
				(*islands)[in0] = append((*islands)[in0], coddin)
			}
		} else {
			*islands = append(*islands, []string{coddin})
		}
	} else {
		*islands = append(*islands, []string{coddin})
	}
}
func main() {
	var input [][]string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the area data: ")
	for scanner.Scan() {
		splitString := strings.Split(scanner.Text(), "")
		input = append(input, strings.Fields(strings.Join(splitString, " ")))
		if scanner.Text() == "" {
			break
		}
	}

	//TODO - Will come from console input
	// input := [][]int{
	// 	{1, 0, 0, 0, 1},
	// 	{1, 1, 0, 0, 0},
	// 	{1, 0, 1, 1, 0},
	// 	{0, 0, 0, 0, 0},
	// }

	var islands [][]string
	for i, res := range input {
		for j, res1 := range res {
			if res1 == "1" {
				checkPartOfIsland(i, j, &islands)
			}
		}
	}
	fmt.Println(len(islands))
}
