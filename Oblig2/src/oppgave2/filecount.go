package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	//Leser filen
	fileInfo := os.Args[1]
	scanLines(fileInfo)
	temp := check(fileInfo)
	freqRunes(temp)

}
//Scanner antall linjer
func scanLines(fileInfo string) int {
	file, err := os.Open(fileInfo)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	//Skriver ut informasjon om fil og antall linjer
	fmt.Println("**********************************")
	fmt.Println("Information about:", fileInfo)
	fmt.Println("")
	fmt.Printf("Number of lines in file: %d\n", count)

	return count
}
//Finner og teller runene
func checkRune(fileInfo string, search string) int {
	file, err := os.Open(fileInfo)
	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 512*1024)
	counter := 0
	searchFor := []byte(search)
	c, _ := file.Read(buffer)
	counter += bytes.Count(buffer[:c], searchFor)
	return counter
}

func check(fileInfo string) map[int]string {
	m := make(map[int]string)

	for i := 0x20; i <= 0x7F; i++ {
		count := checkRune(fileInfo, string(i))
		runer := string(i)
		m[count] = runer
	}
	fmt.Println()
	return m
}
//Sorterer runene
func freqRunes(m map[int]string) {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// Printer de 5 runes som forekommer mest:
	value1 := keys[len(keys)-1]
	value2 := keys[len(keys)-2]
	value3 := keys[len(keys)-3]
	value4 := keys[len(keys)-4]
	value5 := keys[len(keys)-5]

	fmt.Println("Most common runes:")
	fmt.Println("")
	fmt.Println("1. Rune: ", m[value1], "Counts: ", value1)
	fmt.Println("2. Rune: ", m[value2], "Counts: ", value2)
	fmt.Println("3. Rune: ", m[value3], "Counts: ", value3)
	fmt.Println("4. Rune: ", m[value4], "Counts: ", value4)
	fmt.Println("5. Rune: ", m[value5], "Counts: ", value5)
	fmt.Println("**********************************")
}
