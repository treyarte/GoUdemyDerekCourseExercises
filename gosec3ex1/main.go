package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var pl = fmt.Println

var Grades = map[int]string {
	4: "You are in pre K",
	5: "You are in Kindergarten",
	6: "You are a 1st grader",
	7: "You are a 2nd grader",
	8: "You are a 3rd grader",
	9: "You are a 4th grader",
	10: "You are a 5th grader",
	11: "You are a 6th grader",
	12: "You are a 7th grader",
	13: "You are a 8th grader",
	14: "You are a 9th grader",
	15: "You are a 10th grader",
	16: "You are a 11th grader",
	17: "You are a 12th grader",
}
func getSchoolGrade(age int) string {
	if age <= 3 {
		return "You are to young for school"
	}

	if age > 17 {
		return "You are in college or need to go to college"
	}

	return Grades[age]
}

func main() {
	pl("How old are you?")

	reader := bufio.NewReader(os.Stdin)
	ageStr, err := reader.ReadString('\n');
	
	if err != nil {
		log.Fatal("Failed to get input.")
	}
	
	ageStr = strings.TrimSpace(ageStr)
	age, err := strconv.Atoi(ageStr)
	
	if err != nil || age <= 0 {
		log.Fatal("Please provide a valid age")
	}

	schoolMsg := getSchoolGrade(age)

	pl(schoolMsg)
}