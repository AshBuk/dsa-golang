<img src="ALgoDSgo.png" alt="AlgoDSgo Logo" width="250" align="right">

<div align="center">

### AlgoDSgo

#### Algorithms & Data Structures in Go (Golang)

#### Searching, Sorting, Interview Problems, Algorithmic Patterns

#### Educational Computer Science Examples

</div>

>**Educational repository** focused on classic fundamental **algorithms and data structures**, implemented idiomatically in Go.
Includes walkthroughs in complex areas, clean code, semantic naming, and emphasis on building real CS understanding without unnecessary abstraction.

## 📚 Included

<pre>
<strong>algorithms/</strong>
├── <strong>search/</strong>
│   ├── <a href="algorithms/search/binarysearch/binarysearch.go">Binary</a>
│   ├── <a href="algorithms/search/linearsearch/linearsearch.go">Linear</a>
│   ├── <a href="algorithms/search/jumpsearch/jumpsearch.go">Jump</a>
│   ├── <a href="algorithms/search/bfs/filesystem/bfs.go">BFS Filesystem</a> · <a href="algorithms/search/bfs/graph/bfsqueue/bfsqueue.go">BFS Graph Queue</a>
│   └── <a href="algorithms/search/dfs/filesystem/dfs.go">DFS Filesystem</a> · <a href="algorithms/search/dfs/graph/dfsrecursive/dfsrecursive.go">DFS Graph Recursive</a> · <a href="algorithms/search/dfs/graph/dfsstack/dfsstack.go">DFS Graph Stack</a>
└── <strong>sort/</strong>
    ├── <a href="algorithms/sort/bubblesort/bubblesort.go">Bubble</a>
    ├── <a href="algorithms/sort/selectionsort/selectionsort.go">Selection</a>
    ├── <a href="algorithms/sort/insertionsort/insertionsort.go">Insertion</a>
    ├── <a href="algorithms/sort/mergesort/mergesort.go">Merge</a>
    ├── <a href="algorithms/sort/quicksort/quicksort.go">Quick</a>
    ├── <a href="algorithms/sort/quicksortinplace/quicksortinplace.go">Quick In-Place</a>
    └── <a href="algorithms/sort/heapsort/heapsort.go">Heap</a>

<strong>datastructures/</strong>
└── <strong>graph/</strong>
    └── <a href="datastructures/graph/adjacencylist.go">Adjacency List</a>

<strong>leetcode/</strong>
├── <a href="leetcode/twosum/twosum.go">Two Sum #1</a>
├── <a href="leetcode/mergesortedarray/mergesortedarray.go">Merge Sorted Array #88</a>
└── <a href="leetcode/ispalidrome/ispalidrome.go">Valid Palindrome #125</a>

<strong>interview_exercises/</strong>
├── <a href="interview_exercises/reversestring/reversestring.go">Reverse</a>
├── <a href="interview_exercises/removeduplicates/removeduplicates.go">Duplicates</a>
├── <a href="interview_exercises/removewhitespaces/removewhitespaces.go">Whitespaces</a>
└── <a href="interview_exercises/longestword/longestword.go">Longest</a>

<strong>informatica/</strong>
├── <a href="informatica/factorial/factorial.go">Factorial</a>
├── <a href="informatica/fibonacci/fibonacci.go">Fibonacci</a>
├── <a href="informatica/isprime/isprime.go">Prime</a>
└── <a href="informatica/multitable/multitable.go">Table</a>
</pre>

```bash
cd algorithms/sort/bubblesort
go run bubblesort.go
```

This repository provides practical, minimal Go implementations of core Computer Science concepts.  
It includes multiple approaches to BFS/DFS traversal in graphs and filesystems, several sorting strategies including in-place variants,  
basic graph structures, foundational informatics exercises, and structured solutions to common interview and LeetCode tasks.  
The project layout mirrors real code organization, making the repository a compact learning and reference resource.

**🔖[See Article](https://ashbuk.hashnode.dev/go-learning-journey-through-algorithms)**

## MIT [LICENSE](LICENSE)
