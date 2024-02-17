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
		fmt.Printf("\nVertex %v : %v", v.key, v.value)
		for _, v := range v.adjacent {
			fmt.Printf("\nСоседи %v ", v.key)
		}
	}
	fmt.Println()
}

func main() {
	test := &Graph{}
	for i := 0; i < 5; i++ {
		test.AddVertex(i, i+10)
	}
	//test.AddVertex(2, 23)
	test.AddEdge(1, 2)
	test.Print()
}
