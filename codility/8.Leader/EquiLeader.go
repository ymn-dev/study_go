// https://app.codility.com/programmers/lessons/8-leader/
package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// Implement your solution here
	checker := map[int]int{}
	var leader int
	for i := 0; i < len(A); i++ {
		_, exist := checker[A[i]]
		if exist {
			checker[A[i]] += 1
			if checker[A[i]] > len(A)/2 {
				leader = A[i]
			}
		} else {
			checker[A[i]] = 1
		}
	}
	leftLeader := 0
	rightLeader := checker[leader]
	leftCount := 0
	rightCount := len(A)
	equiLeader := 0
	for i := 0; i < len(A); i++ {
		leftCount++
		rightCount--
		if A[i] == leader {
			leftLeader++
			rightLeader--
		}
		if leftLeader > leftCount/2 {
			if rightLeader > rightCount/2 {
				equiLeader++
			}
		}
	}
	return equiLeader
}
