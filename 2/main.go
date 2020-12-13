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
	flagPasswordFile = flag.String("file", "passwords.txt", "a file with a line-separated list of passwords and their policies")
)

func main() {
	flag.Parse()
	pws, err := loadPasswords(*flagPasswordFile)
	if err != nil {
		panic(err)
	}

	Part1(pws)
	Part2(pws)
}

func Part1(pws []string) {
	fmt.Printf("Part 1\n")

	valid := 0
	invalid := 0

	for _, l := range pws {
		if l == "" {
			continue
		}

		r, err := parseOldPolicy(l)
		if err != nil {
			panic(err)
		}

		if r {
			valid++
		} else {
			invalid++
		}
	}

	fmt.Printf("Valid passwords: %d\n", valid)
	fmt.Printf("Invalid passwords: %d\n", invalid)
}

func Part2(pws []string) {
	fmt.Printf("Part 2\n")

	valid := 0
	invalid := 0

	for _, l := range pws {
		if l == "" {
			continue
		}

		r, err := parseNewPolicy(l)
		if err != nil {
			panic(err)
		}

		if r {
			valid++
		} else {
			invalid++
		}
	}

	fmt.Printf("Valid passwords: %d\n", valid)
	fmt.Printf("Invalid passwords: %d\n", invalid)
}

func parseNewPolicy(l string) (valid bool, err error) {
	pol, val, pass := parseLine(l)

	pos1, pos2, err := parsePolicy(pol)
	if err != nil {
		return false, nil
	}

	pos1Found := false
	pos2Found := false

	if pass[pos1-1] == val {
		pos1Found = true
	}
	if pass[pos2-1] == val {
		pos2Found = true
	}

	if pos1Found == pos2Found {
		return false, nil
	}

	return true, nil
}

func parseOldPolicy(l string) (valid bool, err error) {
	pol, val, pass := parseLine(l)

	min, max, err := parsePolicy(pol)
	if err != nil {
		return false, err
	}

	c := countValInPass(pass, val)
	if c > max || c < min {
		return false, nil
	}

	return true, nil
}

func parseLine(l string) (policy string, value rune, password []rune) {
	// split by space
	fields := strings.SplitN(l, " ", 3)
	if len(fields) != 3 {
		panic("Failed to parse line")
	}

	pol := fields[0]
	val := []rune(strings.Trim(fields[1], ":"))[0]
	pass := fields[2]

	return pol, val, []rune(pass)
}

func countValInPass(pass []rune, val rune) int {
	c := 0
	for _, r := range pass {
		if r == val {
			c++
		}
	}

	return c
}

func parsePolicy(p string) (min int, max int, err error) {
	fields := strings.Split(p, "-")

	min, err = strconv.Atoi(fields[0])
	if err != nil {
		return -1, -1, err
	}

	max, err = strconv.Atoi(fields[1])
	if err != nil {
		return -1, -1, err
	}

	return min, max, nil
}

func loadPasswords(f string) ([]string, error) {
	data, err := ioutil.ReadFile(f)

	// error handling
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("File not found")
	}
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}
