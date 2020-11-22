# 343 Programming Project - Minimum Spanning Tree Implementations

<p align="center"><img src="images/updatedRun.png" height="400"></p>

A programming project for an Algorithms class where I decided to implement 3 different Minimum Spanning Tree implementations.
The goal here was to compare the execution times, the graphs at each step, the end result graph, and show which graphs are better in which situtation.

## Version
> **Go 1.14.3**

## How to Run
In the folder with main.go run the command:
> **go run main.go graphs.txt**

Where graphs.txt is the name of the file containing the graphs to run the algorithms on.
> Include the file path if your file is not in the same directory as main.go

## File format
> **n** - number of graphs

> **numOfEdges** - number of edges in each graph

> e0 1 2 3 4 ... edges - first edge with numOfEdges edges where the values are the cost for that connection e0 - e2 costs 2

Ex:
> 1

> 3

> 0 1 2

> 1 0 0

> 2 0 0

### Working on
- [ ] Write the paper for a grade :]
- [ ] Better Displaying System
- [ ] Showing the Graphs at each step
- [X] Adding more test graphs / graph variants (Have 7)
- [X] Showing execution times for each algorithm plus average execution time
- [X] Reading from any file through the command line
- [X] Implement Boruvka's Algorithm
- [X] Implement Kruskal's Algorithm
- [X] Implement a Disjoint Set Union using size / rank
- [X] Implement Prim's Algorithm

### Sources so far
Creating a struct instance using struct literals source: 
    https://www.golangprograms.com/go-language/struct.html

Creating custom sorting for structs source: 
    https://thenotexpert.com/golang-sorting/

How to measure execution time in Golang:
    https://www.admfactory.com/how-to-measure-execution-time-in-golang/

Reading Files: 
    https://golangbot.com/read-files/
    https://golangdocs.com/reading-files-in-golang

Sorting a Collection of Custom Objects:
    https://golangdocs.com/golang-sort-package

Disjoint Set Union:
    https://cp-algorithms.com/data_structures/disjoint_set_union.html

Kruskal's Algorithm:
    https://en.wikipedia.org/wiki/Kruskal%27s_algorithm

Prim's Algorithm:
    https://en.wikipedia.org/wiki/Prim%27s_algorithm
    https://www.tutorialspoint.com/data_structures_algorithms/prims_spanning_tree_algorithm.htm

Boruvka's Algorithm: (The point where the fun stopped)
    https://en.wikipedia.org/wiki/Bor%C5%AFvka%27s_algorithm
    http://www-student.cse.buffalo.edu/~atri/cse331/fall16/recitations/Recitation10.pdf

Why do Kruskal and Prim MST algorithms have different runtimes for sparse and dense graphs?
    https://stackoverflow.com/questions/2011753/why-do-kruskal-and-prim-mst-algorithms-have-different-runtimes-for-sparse-and-de
