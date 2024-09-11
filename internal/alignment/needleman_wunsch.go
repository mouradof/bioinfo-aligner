package alignment

import "math"

func NeedlemanWunsch(seq1, seq2 string, match, mismatch, gap int) ([][]int, string, string) {
	n := len(seq1)
	m := len(seq2)

	scores := make([][]int, n+1)
	for i := range scores {
		scores[i] = make([]int, m+1)
	}

	for i := 0; i <= n; i++ {
		scores[i][0] = i * gap
	}
	for j := 0; j <= m; j++ {
		scores[0][j] = j * gap
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			matchScore := mismatch
			if seq1[i-1] == seq2[j-1] {
				matchScore = match
			}
			scores[i][j] = int(math.Max(
				math.Max(float64(scores[i-1][j-1]+matchScore), float64(scores[i-1][j]+gap)),
				float64(scores[i][j-1]+gap),
			))
		}
	}

	align1 := ""
	align2 := ""
	i, j := n, m

	for i > 0 && j > 0 {
		current := scores[i][j]
		diagonal := scores[i-1][j-1]
		up := scores[i-1][j]
		if current == diagonal+match && seq1[i-1] == seq2[j-1] {
			align1 = string(seq1[i-1]) + align1
			align2 = string(seq2[j-1]) + align2
			i--
			j--
		} else if current == up+gap {
			align1 = string(seq1[i-1]) + align1
			align2 = "-" + align2
			i--
		} else {
			align1 = "-" + align1
			align2 = string(seq2[j-1]) + align2
			j--
		}
	}

	return scores, align1, align2
}
