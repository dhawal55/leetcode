import "fmt"

type Node struct {
	Val   string
	Level int
}

// Time Complexity: O(M^2xN), where M is the length of each word and N is the total number of words in the input word list. For each word in the word list, we iterate over its length to find all the intermediate words corresponding to it. Since the length of each word is M and we have N words, the total number of iterations the algorithm takes to create all_combo_dict is M×N. Breadth first search in the worst case might go to each of the N words. For each word, we need to examine M possible intermediate words/combinations. Notice, we have used the substring operation to find each of the combination. Thus, M combinations take O(M^2) time. As a result, the time complexity of BFS traversal would also be O(M^2xN)
// Space Complexity: O(M^2xN). Each word in the word list would have M intermediate combinations. To create the all_combo_dict dictionary we save an intermediate word as the key and its corresponding original words as the value. Note, for each of M intermediate words we save the original word of length M. This simply means, for every word we would need a space of M^2 to save all the transformations corresponding to it. Thus, all_combo_dict would need a total space of O(M^2xN). Visited dictionary would need a space of O(M×N) as each word is of length M. Queue for BFS in worst case would need a space for all O(N) words and this would also result in a space complexity of 	O(M×N).
func ladderLengthSimple(beginWord string, endWord string, wordList []string) int {
	wordTransformations := make(map[string][]string)
	for i := 0; i < len(wordList); i++ {
		wordArr := []byte(wordList[i])

		for j := 0; j < len(wordArr); j++ {
			newWord := fmt.Sprintf("%s%c%s", wordArr[0:j], '*', wordArr[j+1:])
			wordTransformations[newWord] = append(wordTransformations[newWord], wordList[i])
		}
	}

	queue := []Node{
		Node{
			Val:   beginWord,
			Level: 1,
		},
	}

	visited := make(map[string]bool)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		wordArr := []byte(node.Val)
		for i := 0; i < len(wordArr); i++ {
			newWord := fmt.Sprintf("%s%c%s", wordArr[0:i], '*', wordArr[i+1:])

			transformations := wordTransformations[newWord]
			for _, t := range transformations {
				if visited[t] {
					continue
				}

				if t == endWord {
					return node.Level + 1
				}

				visited[t] = true
				queue = append(queue, Node{Val: t, Level: node.Level + 1})
			}
		}
	}

	return 0
}

func addWordTransformations(word string, wordTransformations map[string][]string) {
	arr := []byte(word)

	for i := 0; i < len(arr); i++ {
		newWord := fmt.Sprintf("%s%s%s", arr[0:i], '*', arr[i+1:])
		wordTransformations[newWord] = append(wordTransformations[newWord], word)
	}
}

// Time Complexity: O(M^2xN), where M is the length of words and N is the total number of words in the input word list. Similar to one directional, bidirectional also takes O(M^2 x N) time for finding out all the transformations. But the search time reduces to half, since the two parallel searches meet somewhere in the middle.
// Space Complexity: O(M^2×N), to store all M transformations for each of the N words in the dictionary, same as one directional. But bidirectional reduces the search space. It narrows down because of meeting in the middle.
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordTransformations := make(map[string][]string)
	endWordFound := false
	for i := 0; i < len(wordList); i++ {
		if wordList[i] == endWord {
			endWordFound = true
		}

		wordArr := []byte(wordList[i])

		for j := 0; j < len(wordArr); j++ {
			newWord := fmt.Sprintf("%s%c%s", wordArr[0:j], '*', wordArr[j+1:])
			wordTransformations[newWord] = append(wordTransformations[newWord], wordList[i])
		}
	}

	if !endWordFound {
		return 0
	}

	queueBegin := []Node{
		Node{
			Val:   beginWord,
			Level: 1,
		},
	}
	queueEnd := []Node{
		Node{
			Val:   endWord,
			Level: 1,
		},
	}

	visitedBegin := make(map[string]int)
	visitedEnd := make(map[string]int)
	visitedBegin[beginWord] = 1
	visitedEnd[endWord] = 1

	var ans int
	for len(queueBegin) > 0 && len(queueEnd) > 0 {
		ans, queueBegin = visitWordNode(queueBegin, visitedBegin, visitedEnd, wordTransformations)
		if ans > 0 {
			return ans
		}

		ans, queueEnd = visitWordNode(queueEnd, visitedEnd, visitedBegin, wordTransformations)
		if ans > 0 {
			return ans
		}
	}

	return 0
}

func visitWordNode(queue []Node, visited, visitedOther map[string]int, wordTransformations map[string][]string) (int, []Node) {
	node := queue[0]
	queue = queue[1:]

	wordArr := []byte(node.Val)
	for i := 0; i < len(wordArr); i++ {
		newWord := fmt.Sprintf("%s%c%s", wordArr[0:i], '*', wordArr[i+1:])

		transformations := wordTransformations[newWord]
		for _, t := range transformations {
			if val, ok := visitedOther[t]; ok {
				return node.Level + val, queue
			}

			if _, ok := visited[t]; !ok {
				visited[t] = node.Level + 1
				queue = append(queue, Node{Val: t, Level: node.Level + 1})
			}
		}
	}

	return -1, queue
}