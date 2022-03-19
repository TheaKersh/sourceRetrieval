package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func m(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("test.html")
	check(err)
	bytes, err := ioutil.ReadAll(f)
	w.Write(bytes)
}

//Function to parse a list of words by
//string params
//Note that it returns a list of indices where it finds
func ParseFor(toParse []string, params []string, escapes []string) ([]int, error) {
	if len(params) != len(escapes) {
		err := *new(error)
		return nil, err
	}
	depth := 0
	retval := make([]int, 0)
	for index, word := range toParse {
		if word == params[depth] {
			fmt.Printf("\n%v\n", depth)
			fmt.Println(word + " \n")
			if depth == len(params)-1 {
				retval = append(retval, index)
			} else {
				depth++
			}
		} else if word == escapes[depth] {
			depth--
		}

	}
	return retval, nil
}

func main() {
	bytes, err := os.ReadFile("test.html")
	check(err)
	s := string(bytes)
	words := strings.Fields(s)
	indices, err := ParseFor(words, []string{"<td", "class=\"colorMyGrade\"><a"}, []string{"</td>", "</a>"})
	check(err)
	f, err := os.Create("links.txt")
	check(err)
	for _, element := range indices {
		fmt.Print(strings.TrimLeft(words[element+1], "href=") + "\n")
		f.Write([]byte(strings.TrimLeft(words[element+1], "href=") + "\n"))
	}
}
