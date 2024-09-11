package api

import (
	"bioinfo-aligner/internal/alignment"
	"bioinfo-aligner/internal/parser"
	"encoding/json"
	"net/http"
)

type AlignResponse struct {
	Alignment1 string `json:"alignment1"`
	Alignment2 string `json:"alignment2"`
	Seq1       string `json:"seq1"`
	Seq2       string `json:"seq2"`
}

func AlignFromFastaHandler(w http.ResponseWriter, r *http.Request) {
	fastaFile := "./data/example.fasta"
	sequences, err := parser.ReadFasta(fastaFile)
	if err != nil {
		http.Error(w, "Erreur lors de la lecture du fichier FASTA", http.StatusInternalServerError)
		return
	}

	if len(sequences) < 2 {
		http.Error(w, "Pas assez de séquences dans le fichier FASTA", http.StatusBadRequest)
		return
	}

	seq1 := sequences[0].Seq
	seq2 := sequences[1].Seq

	_, align1, align2 := alignment.NeedlemanWunsch(seq1, seq2, 1, -1, -2)

	response := AlignResponse{
		Alignment1: align1,
		Alignment2: align2,
		Seq1:       seq1,
		Seq2:       seq2,
	}

	json.NewEncoder(w).Encode(response)
}

func RunServer() {
	http.HandleFunc("/align-from-fasta", AlignFromFastaHandler)
	http.ListenAndServe(":8080", nil)
}
