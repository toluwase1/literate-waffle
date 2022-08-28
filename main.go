package main

import (
	"fmt"
	"strconv"
)

func main() {
	//graph := make(map[string][]string)
	//graph["f"] = []string{"g", "i"}
	fmt.Println(encode("aaaabbbcc"))
	//graph := map[string][]string{
	//	"f": []string{"g", "i"},
	//	"g": []string{"h"},
	//	"h": []string{},
	//	"i": []string{"g", "k"},
	//	"j": []string{"i"},
	//	"k": []string{},
	//}
	//fmt.Println(hasPathBFS(graph, "g", "k"))
	//fmt.Println(hasPathDFS(graph, "g", "k"))
}

//BFS you cant use recursion for bfs, use queue and solve iteratively
//https://structy.net/problems/has-path
func hasPathBFS(graph map[string][]string, src string, dst string) bool {
	var queue []string
	queue = append(queue, src)

	for len(queue) > 0 {
		var current = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if current == dst {
			return true
		}
		for _, neighbour := range graph[current] {
			queue = append(queue, neighbour)
		}
	}
	return false
}

//DFS: use stack or recursion
//https://structy.net/problems/has-path
func hasPathDFS(graph map[string][]string, src string, dst string) bool {
	if src == dst {
		return true
	}
	for _, neighbour := range graph[src] {
		if hasPathDFS(graph, neighbour, dst) {
			return true
		}
	}
	return false
}

//undirected graph means there is an inverse connection also i.e forward and backward
//for undirected graphs, always handle cycles [mark n]odes as visited

func encode(input string) string {
	//check for edge cases
	if len(input) == 0 || input == "" {
		return ""
	}
	//counter, empty string
	var counter int32 = 0
	var outputString string
	var prevCharacter int32 = 0

	//[aaaabbbcc]
	//iterate through the string
	for _, j := range input {
		if j == prevCharacter {
			counter++
		} else {
			if prevCharacter != 0 {
				val := strconv.Itoa(int(counter))
				outputString += val
				outputString += string(prevCharacter)
			}
			prevCharacter = j
			counter = 1
		}
	}
	val := strconv.Itoa(int(counter))
	outputString += val
	outputString += string(prevCharacter)
	return outputString
}
