package main

import (
	"fmt"
)

// Graph structure
type Graph struct {
	vertices []*Vertex
}

// Vertex structure
type Vertex struct {
	key      int
	value    int
	adjacent []*Vertex
}

// Add Vertex
func (g *Graph) AddVertex(k, v int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Вершина %v не добавлена потому, что содержит существующий ключ", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k, value: v})
	}
}

// Add Edge
func (g *Graph) AddEdge(from, to int) {
	//получить вершину
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	//проверить ошибки
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Не верное ребро (%v --> %v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("Существующее ребро (%v --> %v)", from, to)
		fmt.Println(err.Error())
	} else {
		//добавить ребро
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
		toVertex.adjacent = append(toVertex.adjacent, fromVertex)
	}
}

// getVertex возвращает указатель на вершину с числовым ключом
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

// contains
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nВершина %v : %v", v.key, v.value)
		for _, v := range v.adjacent {
			fmt.Printf("\nСосед %v ", v.key)
		}
	}
	fmt.Println()
}

func (g *Graph) bfs(value int) *Vertex {
	for i, v := range g.vertices {
		if v.value == value {
			return g.vertices[i]
		}
	}
	return nil
}

func main() {
	test := &Graph{}
	for i := 0; i < 15; i++ {
		test.AddVertex(i, i+10)
	}
	//test.AddVertex(2, 23)
	test.AddEdge(0, 1)
	test.AddEdge(0, 2)
	test.AddEdge(1, 3)
	test.AddEdge(1, 4)
	test.AddEdge(2, 5)
	test.AddEdge(2, 6)
	test.AddEdge(3, 7)
	test.AddEdge(3, 8)
	test.AddEdge(4, 9)
	test.AddEdge(4, 10)
	test.AddEdge(5, 11)
	test.AddEdge(5, 12)
	test.AddEdge(6, 13)
	test.AddEdge(6, 14)

	test.Print()
	val := 18
	fmt.Printf("\nзначение %v найдено в вершине - %v", test.bfs(val).value, test.bfs(val).key)
}
