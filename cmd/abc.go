package cmd

import (
	"atc/abcmap"
	"bufio"
	"fmt"
	"log"
	"os"
)

func getResourceFile(fn string) *os.File {
	fp, err := os.OpenFile(fn, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return fp
}

func getStringArrayFromResourceFile(fp *os.File) []string {
	var ret []string
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		ret = append(ret, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return ret
}

func MakeFiles(contestnum string, alphabet string) {

	fp := getResourceFile("/Users/htaniguchi/code/AtCoder/ABC/ABC000/main.cpp")
	defer fp.Close()

	ret := getStringArrayFromResourceFile(fp)

	alphabets := []string{"a", "b", "c", "d", "e", "f"}

	for i := 0; i < abcmap.Abcmap[alphabet]; i++ {

		fp, err := os.Create("abc" + contestnum + alphabets[i] + ".cpp")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer fp.Close()

		writer := bufio.NewWriter(fp)

		for idx := range ret {
			writer.WriteString(ret[idx] + "\n")
		}
		writer.Flush()

	}

}
