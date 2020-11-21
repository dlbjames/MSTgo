// Inorder to view the results of all the algorithms
// run the following command in the terminal:
// 		go run main.go fileName.txt
//
// Also be sure to be in the same directory as the main.go file

package main

import (
	"bufio"
	"fmt"
	algo "graphs/algorithms"
	edge "graphs/edge"
	"os"
	"strconv"
	"time"
)

// AverageExecutionTime - The average execution time for each algorithm
type AverageExecutionTime struct {
	primET      time.Duration
	kruskalET   time.Duration
	boruvkaET   time.Duration
	totalGraphs int64
}

var aet AverageExecutionTime = AverageExecutionTime{primET: 0, kruskalET: 0, boruvkaET: 0}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	// Close the file at the end
	defer file.Close()

	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanWords) // use scanwords

	for sc.Scan() {
		numberOfGraphs, _ := strconv.Atoi(sc.Text())
		currentGraphIteration := 1
		aet.totalGraphs = int64(numberOfGraphs)
		sc.Scan()

		for numberOfGraphs > 0 {
			numOfVerticies, _ := strconv.Atoi(sc.Text())
			graph := make([][]int, numOfVerticies)
			seen := make([]bool, numOfVerticies)
			index := 0
			sc.Scan()

			for i := 0; i < numOfVerticies; i++ {
				edgeConnectsTo := make([]int, numOfVerticies)
				for j := 0; j < numOfVerticies; j++ {
					cost, _ := strconv.Atoi(sc.Text())
					edgeConnectsTo[j] = cost
					sc.Scan()
				}

				graph[index] = edgeConnectsTo
				index++
			}

			printAlgorithms(numOfVerticies, currentGraphIteration, seen, graph)

			// Run tests on the next graph
			currentGraphIteration++
			numberOfGraphs--
		}

		// This is used in the case where not all of the test cases
		// in a file need to be read in
		if numberOfGraphs == 0 {
			break
		}
	}

	fmt.Println("Algorithm | Average Execution Time (in µs)")
	fmt.Printf("Prim's    | %d\n", int64(aet.primET.Microseconds())/aet.totalGraphs)
	fmt.Printf("Kruskal's | %d\n", int64(aet.kruskalET.Microseconds())/aet.totalGraphs)
	fmt.Printf("Boruvka's | %d\n", int64(aet.boruvkaET.Microseconds())/aet.totalGraphs)
	fmt.Println()

	if err := sc.Err(); err != nil {
		fmt.Println(err)
	}
}

// printAlgorithms - prints the algorithms to the console
//	Arguments: the number of verticies, the current graph iteration, a visited array, and the graph
func printAlgorithms(numOfVerticies, count int, seen []bool, graph [][]int) {
	// The list of edges for Kruskal's and Boruvka's algorithms
	edges := edge.MakeEdges(graph, numOfVerticies)

	// Set up and run the different algorithms
	fmt.Printf("Graph #%d\n\n", count)
	start := time.Now()
	prims := algo.Prims{Verticies: numOfVerticies, Seen: seen, Graph: graph}

	fmt.Printf("Prim's\n")
	fmt.Println("Edges | Cost")
	prims.Construct()

	elapsed := time.Since(start)
	aet.primET += elapsed
	printExecutionTimes(elapsed)

	start = time.Now()
	kruskals := algo.Kruskals{Verticies: numOfVerticies, Edges: edges}

	fmt.Printf("Kruskal's\n")
	fmt.Println("Edges | Cost")
	kruskals.Construct()

	elapsed = time.Since(start)
	aet.kruskalET += elapsed
	printExecutionTimes(elapsed)

	start = time.Now()
	boruvkas := algo.Boruvkas{Verticies: numOfVerticies, Edges: edges}

	fmt.Printf("Boruvka's\n")
	fmt.Println("Edges | Cost")
	boruvkas.Construct()

	elapsed = time.Since(start)
	aet.boruvkaET += elapsed
	printExecutionTimes(elapsed)
}

// printExecutionTimes - prints the execution time to complete an algorithm on a graph
//	Arguments: the duration it took an algorithm to run and complete
func printExecutionTimes(elapsedTime time.Duration) {
	fmt.Printf("Execution Time: %s\n", elapsedTime)
	fmt.Println()
}
