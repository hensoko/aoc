package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	flagExpenseFile = flag.String("file", "expenses.txt", "a file with a line-separated list of expenses")
)

func main() {
	flag.Parse()

	exp, err := loadExpenses(*flagExpenseFile)
	if err != nil {
		panic(err)
	}

	Part1(exp)
	Part2(exp)
}

func Part1(exp []int) {
	for _, e1 := range exp {
		for _, e2 := range exp {
			if e1+e2 == 2020 {
				fmt.Printf("Solution is: %d\n", e1*e2)
				return
			}
		}
	}
}

func Part2(exp []int) {
	for _, e1 := range exp {
		for _, e2 := range exp {
			for _, e3 := range exp {
				if e1+e2+e3 == 2020 {
					fmt.Printf("Solution is: %d\n", e1*e2*e3)
					return
				}
			}
		}
	}
}

func loadExpenses(f string) ([]int, error) {
	data, err := ioutil.ReadFile(f)

	// error handling
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("File not found")
	}
	if err != nil {
		return nil, err
	}

	// parsing to int
	lines := strings.Split(string(data), "\n")

	exp := []int{}
	for _, l := range lines {
		if l == "" {
			continue
		}

		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}

		exp = append(exp, n)
	}

	return exp, nil
}
