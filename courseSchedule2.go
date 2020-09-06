package main

import "fmt"

func main() {
	numCourses := 4
	prerequisites := [][]int{
		[]int{1, 0},
		[]int{2, 0},
		[]int{3, 1},
		[]int{3, 2},
	}

	fmt.Println(findOrder(numCourses, prerequisites))
}

const WHITE = 0
const GREY = 1
const BLACK = 2

func findOrderDFS(numCourses int, prerequisites [][]int) []int {
	adjacentList := make(map[int][]int)
	color := make([]int, numCourses)

	for i := 0; i < len(prerequisites); i++ {
		course := prerequisites[i][0]
		prereq := prerequisites[i][1]

		adjacentList[course] = append(adjacentList[course], prereq)
	}

	var result []int
	var possible bool
	for i := 0; i < numCourses; i++ {
		if color[i] == WHITE {
			result, possible = dfs(i, adjacentList, color, result)
			if !possible {
				return nil
			}
		}
	}

	return result
}

func dfs(i int, adjacent map[int][]int, color []int, result []int) ([]int, bool) {
	color[i] = GREY

	possible := true
	for _, v := range adjacent[i] {
		if color[v] == WHITE {
			result, possible = dfs(v, adjacent, color, result)
			if !possible {
				return result, false
			}
		} else if color[v] == GREY {
			return result, false
		}
	}

	color[i] = BLACK
	result = append(result, i)
	return result, possible
}

// Time Complexity: O(V+E) where V represents the number of vertices and E represents the number of edges.
func findOrder(numCourses int, prerequisites [][]int) []int {
	parentList := make(map[int][]int)
	indegree := make([]int, numCourses)

	for i := 0; i < len(prerequisites); i++ {
		course := prerequisites[i][0]
		dependency := prerequisites[i][1]

		parentList[dependency] = append(parentList[dependency], course)
		indegree[course] += 1
	}

	var queue []int
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	var result []int
	var i int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)
		i++

		if courses, ok := parentList[node]; ok {
			for _, v := range courses {
				indegree[v] = indegree[v] - 1

				if indegree[v] == 0 {
					queue = append(queue, v)
				}
			}
		}
	}

	if len(result) == numCourses {
		return result
	}

	return nil
}
