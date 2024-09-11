package alignment

import "testing"

func TestNeedlemanWunsch(t *testing.T) {
	seq1 := "ACGT"
	seq2 := "ACT"
	match := 1
	mismatch := -1
	gap := -2

	_, align1, align2 := NeedlemanWunsch(seq1, seq2, match, mismatch, gap)
	t.Logf("Alignement 1: %s", align1)
	t.Logf("Alignement 2: %s", align2)
}
