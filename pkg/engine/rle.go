package engine

import (
	"bufio"
	"fmt"
	"io"
)

var _ bufio.SplitFunc = rleSplitter

func rleSplitter(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if data[0] == '!' {
		return 0, nil, bufio.ErrFinalToken
	}
	start := 0
	for ; start < len(data); start++ {
		if data[start] != '$' && data[start] != '!' {
			break
		}
	}
	for i := start; i < len(data); i++ {
		if '0' <= data[i] && data[i] <= '9' {
			continue
		} else if data[i] == 'b' || data[i] == 'o' {
			return i + 1, data[start : i+1], nil
		} else {
			return 0, nil, fmt.Errorf("idk")
		}
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}

func FromRLE(rd io.Reader) (*Engine, error) {
	var cols, rows int64
	if _, err := fmt.Fscanf(rd, "x = %d, y = %d\n", &cols, &rows); err != nil {
		return nil, err
	}
	fmt.Printf("rows %d cols %d \n", rows, cols)
	g := EmptyGame(rows, cols)
	s := bufio.NewScanner(rd)
	s.Split(rleSplitter)
	for s.Scan() {
		token := s.Text()
		fmt.Printf("token: {%s}\n", token)
		fmt.Println()
	}
	fmt.Printf("scan done\n")
	if err := s.Err(); err != nil {
		return nil, err
	}
	return g, nil
}
