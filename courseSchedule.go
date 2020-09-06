package main

import "fmt"

func main() {
	numCourses := 2
	prerequisites := [][]int{
		[]int{1, 0},
		[]int{0, 1},
	}

	fmt.Println(canFinish(numCourses, prerequisites))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	adjacentList := make(map[int][]int)
	for i := 0; i < len(prerequisites); i++ {
		course := prerequisites[i][0]
		dependency := prerequisites[i][1]

		list := adjacentList[course]
		adjacentList[course] = append(list, dependency)

	}

	visited := make([]bool, numCourses)
	stackRec := make([]bool, numCourses)

	for i := 0; i < len(prerequisites); i++ {
		if isCyclic(prerequisites[i][0], visited, stackRec, adjacentList) {
			return false
		}
	}

	return true
}

func isCyclic(i int, visited, stackRec []bool, adjacentList map[int][]int) bool {
	if stackRec[i] {
		return true
	}

	if visited[i] {
		return false
	}

	visited[i] = true
	stackRec[i] = true

	for _, v := range adjacentList[i] {
		if isCyclic(v, visited, stackRec, adjacentList) {
			return true
		}
	}

	stackRec[i] = false
	return false
}
