package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	lines, err := readLines("/Users/Assad.Ginem/list_of_calories.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	var elves []int
	var sum = 0
	for _, line := range lines {
		if line != "" {
			var n, err = strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
			}
			sum += n
		} else {
			elves = append(elves, sum)
			sum = 0
			continue
		}
	}
	max := elves[0]

	for i := 1; i < len(elves); i++ {
		if max < elves[i] {
			max = elves[i]
		}
	}

	sort.Ints(elves)

	var sum_of_three = 0
	for _, elf := range elves[len(elves)-3:] {
		sum_of_three += elf
	}

	fmt.Println("Number of elves: ", len(elves))
	fmt.Println("Elf with the larges amount of food calories: ", max)
	fmt.Println("Sum of elves with top most calories: ", sum_of_three)
}
