package cmd

import (
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

func MakeFiles() {

	fp := getResourceFile("/Users/htaniguchi/code/AtCoder/ABC/ABC000/main.cpp")
	defer fp.Close()

	ret := getStringArrayFromResourceFile(fp)

	fp, err := os.Create("abc156a.cpp")
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
