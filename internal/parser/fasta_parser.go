package parser

import (
	"bufio"
	"os"
	"strings"
)

type Sequence struct {
	ID  string
	Seq string
}

func ReadFasta(filename string) ([]Sequence, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var sequences []Sequence
	var currentSeq Sequence
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, ">") {
			if currentSeq.ID != "" {
				sequences = append(sequences, currentSeq)
			}
			currentSeq = Sequence{
				ID:  strings.TrimPrefix(line, ">"),
				Seq: "",
			}
		} else {
			currentSeq.Seq += strings.TrimSpace(line)
		}
	}
	if currentSeq.ID != "" {
		sequences = append(sequences, currentSeq)
	}
	return sequences, scanner.Err()
}
