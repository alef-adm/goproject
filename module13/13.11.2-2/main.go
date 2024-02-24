package main

import "fmt"
import "sort"

var vertices []string
var adjacencyList [][]string
var visited []bool
var level []int

func insertVertex(vertex string) {

	vertices = append(vertices, vertex)
}

func constructTempList(tmpList *[]string, val string) {

	*tmpList = append(*tmpList, val)
}

func constructAdjacencyList(adjcList *[][]string, tmpList *[]string) {

	*adjcList = append(*adjcList, *tmpList)
	*tmpList = nil
}

func breadthFirstSearch(source string) {

	var queue []int

	sourceIndex := sort.StringSlice(vertices).Search(source)

	visited[sourceIndex] = true
	queue = append(queue, sourceIndex)

	for len(queue) != 0 {

		sourceIndex = queue[0]

		// Slice off the element once it is dequeued.
		queue = queue[1:]

		for j := 0; j < len(adjacencyList[sourceIndex]); j++ {

			str := adjacencyList[sourceIndex][j]

			i := sort.StringSlice(vertices).Search(str)

			if visited[i] == false {

				visited[i] = true
				level[i] = level[sourceIndex] + 1

				queue = append(queue, i)
			}
		}
		fmt.Print(" --> ", vertices[sourceIndex], " of level ", level[sourceIndex])
	}
	fmt.Println()
}

func main() {

	const V = 5

	var tempList []string

	visited = make([]bool, V)

	level = make([]int, V)

	// Adding vertices one by one

	insertVertex("a")
	insertVertex("b")
	insertVertex("c")
	insertVertex("d")
	insertVertex("e")

	// Create Adjacency List

	constructTempList(&tempList, "b")
	constructTempList(&tempList, "c")
	constructTempList(&tempList, "d")

	constructAdjacencyList(&adjacencyList, &tempList)

	constructTempList(&tempList, "a")
	constructTempList(&tempList, "e")

	constructAdjacencyList(&adjacencyList, &tempList)

	constructTempList(&tempList, "a")
	constructTempList(&tempList, "d")
	constructTempList(&tempList, "e")

	constructAdjacencyList(&adjacencyList, &tempList)

	constructTempList(&tempList, "a")
	constructTempList(&tempList, "c")
	constructTempList(&tempList, "e")

	constructAdjacencyList(&adjacencyList, &tempList)

	constructTempList(&tempList, "b")
	constructTempList(&tempList, "c")
	constructTempList(&tempList, "d")

	constructAdjacencyList(&adjacencyList, &tempList)

	breadthFirstSearch("e")

}
