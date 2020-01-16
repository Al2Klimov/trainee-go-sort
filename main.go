package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	numericParameter := flag.Bool("n", false, "Sort numbers from lowest to highest")
	flag.Parse()

	var data []byte

	if len(flag.Args()) < 1 {
		content, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err)
			return
		}
		data = content
	} else {
		fileContent, readFileErr := ioutil.ReadFile(flag.Args()[0])
		if readFileErr != nil {
			fmt.Println(readFileErr)
			return
		}
		data = fileContent
	}

	if string(data) == "" {
		return
	}

	dataSplit := bytes.Split(data, []byte("\n"))
	lengthOfData := len(dataSplit)

	for i := 0; i < lengthOfData; i++ {
		j := i
		for j > 0 {
			if !*numericParameter {
				if bytes.Compare(dataSplit[j], dataSplit[j-1]) < 0 {
					dataSplit[j], dataSplit[j-1] = dataSplit[j-1], dataSplit[j]
					j -= 1
				} else {
					break
				}
			} else {
				comparer1, convertErr1 := strconv.Atoi(string(dataSplit[j]))
				if convertErr1 != nil {
					fmt.Println(convertErr1)
					return
				}
				comparer2, convertErr2 := strconv.Atoi(string(dataSplit[j-1]))
				if convertErr2 != nil {
					fmt.Println(convertErr2)
					return
				}
				if comparer1 < comparer2 {
					dataSplit[j], dataSplit[j-1] = dataSplit[j-1], dataSplit[j]
					j -= 1
				} else {
					break
				}
			}
		}
	}
	dataJoin := bytes.Join(dataSplit, []byte("\n"))
	fmt.Printf("%s\n", strings.TrimPrefix(string(dataJoin), "\n"))
}
