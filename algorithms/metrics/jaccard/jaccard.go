// Jaccard Similarity Coefficient
// Measures similarity between two finite sets
// Formula: J(A,B) = |A ∩ B| / |A ∪ B|
// Time: O(n + m) where n, m are set sizes
// Space: O(n + m) for set conversion

package main

import "fmt"

// Jaccard returns the similarity coefficient between two sets
// Range: 0.0 (disjoint) to 1.0 (identical)
func Jaccard[T comparable](a, b []T) float64 {
	if len(a) == 0 && len(b) == 0 {
		return 1.0
	}
	setA := toSet(a)
	setB := toSet(b)
	intersection := 0
	for elem := range setA {
		if _, ok := setB[elem]; ok {
			intersection++
		}
	}
	// Union formula: |A ∪ B| = |A| + |B| - |A ∩ B|
	union := len(setA) + len(setB) - intersection
	return float64(intersection) / float64(union)
}

func toSet[T comparable](slice []T) map[T]struct{} {
	set := make(map[T]struct{}, len(slice))
	for _, v := range slice {
		set[v] = struct{}{}
	}
	return set
}

func main() {
	a := []string{"go", "rust", "zig"}
	b := []string{"go", "rust", "c++"}
	// Jaccard coefficient = 0.50 (half modern, half legacy :D)
	fmt.Println(Jaccard(a, b))
}

/*
How Jaccard Similarity works:

Formula: J(A,B) = |A ∩ B| / |A ∪ B|

Symbols:
    A, B    — two sets to compare
    ∩       — intersection (elements in BOTH sets)
    ∪       — union (ALL unique elements from both)
    |...|   — cardinality (count of elements)

Example: A = {go, rust, zig}, B = {go, rust, c++}

STEP 1 - Convert to sets (remove duplicates):
    setA = {go, rust, zig}
    setB = {go, rust, c++}

STEP 2 - Find intersection (elements in both):
    go   → in setA? ✓  in setB? ✓  → count
    rust → in setA? ✓  in setB? ✓  → count
    zig  → in setA? ✓  in setB? ✗
    c++  → in setA? ✗  in setB? ✓

    intersection = {go, rust} → |A ∩ B| = 2

STEP 3 - Calculate union:
    |A ∪ B| = |A| + |B| - |A ∩ B|
            = 3 + 3 - 2 = 4

    union = {go, rust, zig, c++} → 4 elements

STEP 4 - Calculate similarity:
    J(A,B) = |A ∩ B| / |A ∪ B|
           = 2 / 4
           = 0.50 ✓

Result interpretation:
    1.0  — sets are identical
    0.5  — half of elements are shared
    0.0  — sets are completely disjoint (nothing in common)

Key points:
- Measures similarity as ratio of shared to total elements
- Order of elements does not matter (set property)
- Duplicates are ignored (set property)
- Used in: plagiarism detection, recommendation systems, NLP
*/
