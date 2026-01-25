// Levenshtein Distance (Edit Distance) — Wagner-Fischer algorithm
// Measures minimum edits to transform one string into another
// Operations: insert, delete, replace (each costs 1)
// Time: O(n × m) where n, m are string lengths
// Space: O(n × m) for the DP (Dynamic Programming) matrix

package main

import "fmt"

// Levenshtein returns the minimum edit distance between two strings
// Range: 0 (identical) to max(len(a), len(b)) (no shared characters)
func Levenshtein(a, b string) int {
	if a == b {
		return 0
	}
	if len(a) == 0 {
		return len(b)
	}
	if len(b) == 0 {
		return len(a)
	}

	n, m := len(a), len(b)

	// DP matrix: dp[i][j] = min edits to transform a[0..i] into b[0..j]
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	// Base cases: transforming empty string
	for i := 0; i <= n; i++ {
		dp[i][0] = i // delete all characters from a
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = j // insert all characters into a
	}

	// Fill the matrix
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] // characters match, no operation needed
			} else {
				dp[i][j] = 1 + min(
					dp[i-1][j],   // delete from a
					dp[i][j-1],   // insert into a
					dp[i-1][j-1], // replace in a
				)
			}
		}
	}
	return dp[n][m]
}

func main() {
	a := "master"
	b := "main"
	// Edit distance = 4
	fmt.Println(Levenshtein(a, b))
}

/*
How Levenshtein Distance works (Wagner-Fischer algorithm):

DP = Dynamic Programming — technique that solves problems by breaking them into
overlapping subproblems and storing results to avoid recomputation.

Formula: dp[i][j] = minimum edits to transform a[0..i] into b[0..j]

Operations (each costs 1):
    INSERT  — add a character to a
    DELETE  — remove a character from a
    REPLACE — change a character in a

Example: a = "master", b = "main"

STEP 1 - Initialize DP matrix:
    Base cases represent transforming to/from empty string

        ""  m  a  i  n
    ""   0  1  2  3  4
    m    1  .  .  .  .
    a    2  .  .  .  .
    s    3  .  .  .  .
    t    4  .  .  .  .
    e    5  .  .  .  .
    r    6  .  .  .  .

STEP 2 - Fill matrix cell by cell (first row example):

    Computing dp[1][1] — transform "m" into "m":
        'm' == 'm' → characters match!
        dp[1][1] = dp[0][0] = 0  (no edit needed)

        ""  m  a  i  n
    ""   0  1  2  3  4
    m    1 [0] .  .  .    ← took diagonal (match)

    Computing dp[1][2] — transform "m" into "ma":
        'm' != 'a' → need operation
        dp[1][2] = 1 + min(
            dp[0][2]=2,  ← delete 'm', then "" → "ma" costs 2 → total 3
            dp[1][1]=0,  ← "m" → "m" costs 0, then insert 'a' → total 1 ✓
            dp[0][1]=1   ← replace 'm' with 'm', then "" → "a" → total 2
        ) = 1

        ""  m  a  i  n
    ""   0  1  2  3  4
    m    1  0 [1] .  .    ← took left + insert

    Computing dp[2][2] — transform "ma" into "ma":
        'a' == 'a' → characters match!
        dp[2][2] = dp[1][1] = 0  (no edit needed)

        ""  m  a  i  n
    ""   0  1  2  3  4
    m    1  0  1  2  3
    a    2  1 [0] .  .    ← took diagonal (match)

    Computing dp[3][3] — transform "mas" into "mai":
        's' != 'i' → need operation
        dp[3][3] = 1 + min(
            dp[2][3]=1,  ← delete 's' from "mas", "ma"→"mai" costs 1 → total 2
            dp[3][2]=1,  ← "mas"→"ma" costs 1, insert 'i' → total 2
            dp[2][2]=0   ← "ma"→"ma" costs 0, replace 's'→'i' → total 1 ✓
        ) = 1

        ""  m  a  i  n
    ""   0  1  2  3  4
    m    1  0  1  2  3
    a    2  1  0  1  2
    s    3  2  1 [1] .    ← took diagonal + replace

STEP 3 - Completed matrix:

        ""  m  a  i  n
    ""   0  1  2  3  4
    m    1  0  1  2  3
    a    2  1  0  1  2
    s    3  2  1  1  2
    t    4  3  2  2  2
    e    5  4  3  3  3
    r    6  5  4  4  4

    Answer: dp[6][4] = 4 ✓

STEP 4 - Trace back the operations:
    master → maiter (replace s with i)
    maiter → maiter (replace t with n)
    ... or simpler path:
    master → maser  (delete t)
    maser  → maer   (delete s)
    maer   → mair   (replace e with i)
    mair   → main   (replace r with n)

    Total: 4 operations ✓

Key points:
- Classic dynamic programming approach
- Each cell depends on 3 neighbors (left, top, diagonal)
- Diagonal = match/replace, Left = insert, Top = delete
- Used in: spell checkers, DNA sequencing, diff tools, fuzzy search
*/
