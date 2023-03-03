// Prioritas 1-2
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Student struct{
	name []string
	score []int
}

func (student Student) averageScore() int {
	totalScore := 0

	for i := range student.score {
		totalScore += student.score[i]
	}

	return totalScore / len(student.score)
}

func (student Student) getScore() (string, string, int, int) {
	minScore := student.score[0]
	maxScore := student.score[0]
	minScoreName := student.name[0]
	maxScoreName := student.name[0]
	
	for i := range student.score {
		currentScore := student.score[i]

		if minScore > currentScore {
			minScore = currentScore
			minScoreName = student.name[i]
		} 

		if maxScore < currentScore {
			maxScore = currentScore
			maxScoreName = student.name[i]
		}
	}
	
	return maxScoreName, minScoreName, maxScore, minScore
}

func main() {
	listStudent := []string{}
	listScore := []int{}

	for i := 1; i <= 5; i++ {
		fmt.Print("Input ", i, " Student's Name ")
		inputName := bufio.NewScanner(os.Stdin)
		inputName.Scan()

		fmt.Print("Input ", i, " Student's Score ")
		inputScore := bufio.NewScanner(os.Stdin)
		inputScore.Scan()
		score,_ := strconv.Atoi(string(inputScore.Text()))

		listStudent = append(listStudent, string(inputName.Text()))
		listScore = append(listScore, score)
	}

	students := Student{
		name: listStudent,
		score: listScore,
	}
	
	avgScore := students.averageScore()
	maximalScoreName, minimalScoreName, maximalScore, minimalScore := students.getScore()

	fmt.Println()
	fmt.Println("Average Score:", avgScore)
	fmt.Println("Min Score of Students :", minimalScoreName, minimalScore)
	fmt.Println("Max Score of Students :", maximalScoreName, maximalScore)
}