package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type Problem struct {
	Question string
	Ans string
}

func main(){
	csvFilename := flag.String("csv","problems.csv", " a csv file in a format of 'question,answer' ")
	flag.Parse()

	file,err := os.Open(*csvFilename)
	if err != nil{
		exit(fmt.Sprintf("File Error %s\n", *csvFilename))

	}
	reader := csv.NewReader(file)
	lines,err :=reader.ReadAll()
	if err != nil{
		exit(fmt.Sprintf("Unable To Parse Error\n"))

	}

	problems := parseLines(lines)
	fmt.Println(problems)

}

func exit (msg string){
	fmt.Println(msg)
	os.Exit(1)
}

func parseLines(lines [][]string) []Problem{
	retValue := make([]Problem,len(lines))

	for i,line := range lines{
		retValue[i] = Problem{
			Question:line[0],
			Ans:      line[1],
		}
	}

	return retValue
}
