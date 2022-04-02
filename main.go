package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/go-co-op/gocron"
	"golang.org/x/net/html"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func ParseHtml() []string {
	file, err := os.Open("test.html")
	check(err)
	doc, err := html.Parse(file)
	check(err)
	elements := make([]string, 0)
	var f func(*html.Node)
	f = func(n *html.Node) {
		nextNode := n
		if nextNode.Type == html.ElementNode {
			for c := nextNode.FirstChild; c != nil; c = c.NextSibling {
				if c.Data == "a" && c.Parent.Data == "td" {
					for _, a := range c.Parent.Attr {
						if a.Key == "class" && a.Val == "colorMyGrade" {
							elements = append(elements, c.FirstChild.Data)
						}
					}
				}
			}
		}
		if nextNode.FirstChild != nil {
			for c := nextNode.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
	}
	f(doc)
	return elements
}

func writeResponse(w http.ResponseWriter, r *http.Request) {
	bytes, err := os.ReadFile("outFile.html")
	check(err)
	w.Write(bytes)
}

var gradeOrder map[string]int = map[string]int{
	"A":  1,
	"A-": 2,
	"B+": 3,
	"B":  4,
	"B-": 5,
	"C+": 6,
	"C-": 7,
	"D+": 8,
	"D":  9,
	"D-": 10,
	"E":  11,
}

func main() {
	sched := gocron.NewScheduler(time.Now().Location())
	f, err := os.OpenFile("grades.txt", os.O_RDWR|os.O_TRUNC, 0755)
	check(err)
	g := ParseHtml()
	fmt.Print(g)
	for _, element := range g {
		_, err := f.WriteString(element + " ")
		check(err)

	}
	sched.Every(3600).Seconds().Do(func() {
		fmt.Print("executing...")
		grades := ParseHtml()
		grades = grades[:36]

		tmpl := template.Must(template.ParseFiles("gradeView.html"))
		check(err)
		st, err := os.ReadFile("grades.txt")
		check(err)
		str := string(st)
		fmt.Print(str)
		check(err)
		gradesAsStr := ""
		for _, element := range grades {
			gradesAsStr += element + " "
		}
		gradesAsStr = strings.ReplaceAll(gradesAsStr, "[ i ]", "")
		grades = strings.Fields(gradesAsStr)
		str = strings.ReplaceAll(str, "[ i ]", "")
		strlist := strings.Fields(str)

		fmt.Printf("\n\n%v\n\n", strlist)
		fmt.Printf("\n\n%v\n\n", grades)
		for index := range strlist {
			print(index)
			if strlist[index] == "Nothing here yet" {
				strlist[index] = "B"
			}
			if grades[index] == "[ i ]" {
				grades[index] = "B"
			}
			switch {
			case gradeOrder[strlist[index]] < gradeOrder[grades[index]]:
				err = beeep.Notify("Your grade has gone down", "from a/an "+strlist[index]+"to a/an"+grades[index], "assets/information.png")
				check(err)
				fmt.Println(grades[index])
				fmt.Println(strlist[index])
				time.Sleep(100000000)
			case gradeOrder[strlist[index]] > gradeOrder[grades[index]]:
				err = beeep.Notify("Your grade has gone up", "from a/an "+strlist[index]+"to a/an"+grades[index], "assets/information.png")
				check(err)
				time.Sleep(100000000)
			default:
				err = beeep.Notify("All quiet on the western front", "o7", "assets/information.png")
				check(err)
				time.Sleep(1000000)
			}
		}

		f, err = os.OpenFile("gradeview.html", os.O_RDWR|os.O_TRUNC, 0755)
		check(err)
		tmpl.Execute(f, grades)
		f, err = os.OpenFile("grades.txt", os.O_RDWR|os.O_TRUNC, 0755)
		check(err)
		for _, element := range grades {
			_, err := f.WriteString(element + " ")
			check(err)
		}
	})
	sched.StartBlocking()

}
