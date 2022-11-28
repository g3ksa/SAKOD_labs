package main

import (
	"bufio"
	"fmt"
	binaryTree "lab4_SAKOD/BinaryTree"
	"os"
	"strconv"
	"strings"
)

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		temp, err := sliceAtoi(strings.Split(scanner.Text(), ","))
		if err != nil {
			panic(err.Error())
		}

		lines = append(lines, temp...)

	}
	return lines, scanner.Err()

}

func main() {
	tree := &binaryTree.BinaryTree{}
	tree = binaryTree.BuildBalancedTree([]int{1, 5, 3, 8, 2, 7}, 0, 6)
	binaryTree.Print(os.Stdout, tree.Root, 0, 'M')

	lines, err := readLines("test.txt")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(lines)
}
