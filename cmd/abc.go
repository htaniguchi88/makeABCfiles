package cmd

import (
	"atc/abcmap"
	"bufio"
	"fmt"
	"log"
	"os"
)

var resourcePath string = "/Users/htaniguchi/code/AtCoder/ABC/ABC000/main.cpp"

// MakeABCFiles ABCコンテストのファイルを生成します
func MakeABCFiles(contestnum string, alphabet string, contestname string) {

	inputFp := getResourceFile(resourcePath)
	defer inputFp.Close()

	ret := getStringArrayFromResourceFile(inputFp)

	generate(contestnum, alphabet, ret, contestname)

}

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

func generate(contestnum string, alphabet string, ret []string, contestname string) {

	alphabets := []string{"a", "b", "c", "d", "e", "f"}

	for i := 0; i < abcmap.Abcmap[alphabet]; i++ {
		outputFp, err := os.Create(contestname + contestnum + alphabets[i] + ".cpp")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer outputFp.Close()

		writeFile(outputFp, ret)
	}
}

func writeFile(fp *os.File, ret []string) {
	writer := bufio.NewWriter(fp)

	for idx := range ret {
		writer.WriteString(ret[idx] + "\n")
	}
	writer.Flush()
}
