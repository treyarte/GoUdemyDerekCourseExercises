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


func getSchoolGrade(age int) string {
	if age <= 3 {
		return "You are to young for school"
	}

	if age > 17 {
		return "Go to college"
	}

	grade := age - 5

	gradeMsg := fmt.Sprintf("Go to grade %d\n", grade)

	return gradeMsg
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