package parser

import (
	"testing"
)

func TestReadFasta(t *testing.T) {
	sequences, err := ReadFasta("exemple.fasta")
	if err != nil {
		t.Errorf("Erreur lors de la lecture du fichier: %v", err)
	}

	if len(sequences) == 0 {
		t.Errorf("Aucune séquence trouvée dans le fichier.")
	}

	for _, seq := range sequences {
		t.Logf("ID: %s, Seq: %s", seq.ID, seq.Seq)
	}
}
