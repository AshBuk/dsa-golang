// Fibonacci Number Calculation (with memoization)
// Time: O(n) - each number calculated once and cached
// Space: O(n) - memo map stores up to n values

package main

import "fmt"

func fibonacci(n int) int {
	memo := make(map[int]int)
	var fib func(int) int
	fib = func(n int) int {
		if n <= 1 {
			return n
		}
		if v, ok := memo[n]; ok {
			return v
		}
		memo[n] = fib(n-1) + fib(n-2)
		return memo[n]
	}
	return fib(n)
}

func main() {
	fmt.Println(fibonacci(8))
}

// Naive recursive implementation (without memoization)
// Time: O(2^n) - exponential, recalculates same values many times
// Space: O(n) - recursion call stack depth

// func fibonacci(n int) int {
// 	if n <= 1 {
// 		return n
// 	}
// 	return fibonacci(n-1) + fibonacci(n-2)
// }

/*
Execution trace for memoized Fibonacci:

fib(5)
├─ memo[5] not found → calculate
├─ fib(4)
│  ├─ memo[4] not found → calculate
│  ├─ fib(3)
│  │  ├─ memo[3] not found → calculate
│  │  ├─ fib(2)
│  │  │  ├─ memo[2] not found → calculate
│  │  │  ├─ fib(1) → 1 (base case)
│  │  │  ├─ fib(0) → 0 (base case)
│  │  │  └─ memo[2] = 1 + 0 = 1 ✓
│  │  ├─ fib(1) → 1 (base case)
│  │  └─ memo[3] = 1 + 1 = 2 ✓
│  ├─ fib(2) → 1 (found in memo!) ✅
│  └─ memo[4] = 2 + 1 = 3 ✓
├─ fib(3) → 2 (found in memo!) ✅
└─ memo[5] = 3 + 2 = 5 ✓

So memo isn’t copied with each call, it lives on the heap,
and all recursive calls point to (share) the same memory location.

Result: 5
Memo state: {0:0, 1:1, 2:1, 3:2, 4:3, 5:5}

Without memoization (naive recursion):
    fib(5) would make 15 function calls (many duplicate calculations)
With memoization:
    Only 9 calls - each fib(n) calculated exactly once
*/
