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

func (student Student) minScore() (string,int) {
	minScore := student.score[0]
	minScoreName := student.name[0]
	for i := range student.score {
		currentScore := student.score[i]

		if minScore > currentScore {
			minScore = currentScore
			minScoreName = student.name[i]
		}
	}
	return minScoreName, minScore
}

func (student Student) maxScore() (string,int) {
	maxScore := student.score[0]
	maxScoreName := student.name[0]
	for i := range student.score {
		currentScore := student.score[i]

		if maxScore < currentScore {
			maxScore = currentScore
			maxScoreName = student.name[i]
		}
	}
	return maxScoreName, maxScore
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
	minimalScoreName, minimalScore := students.minScore()
	maximalScoreName, maximalScore := students.maxScore()

	fmt.Println()
	fmt.Println("Average Score:", avgScore)
	fmt.Println("Min Score of Students :", minimalScoreName, minimalScore)
	fmt.Println("Max Score of Students :", maximalScoreName, maximalScore)
}