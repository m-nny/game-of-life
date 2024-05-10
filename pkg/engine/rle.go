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
			res := data[start : i+1]
			if i == start {
				res = append([]byte{'1'}, res...)
			}
			return i + 1, res, nil
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
	cur_idx := 0
	for s.Scan() {
		token := s.Text()
		fmt.Printf("token: {%s}\n", token)
		var n int
		var r rune
		if _, err := fmt.Sscanf(token, "%d%c", &n, &r); err != nil {
			return nil, err
		}
		fmt.Printf("n: %d r: %c\n", n, r)
		fmt.Println()
		if r == 'b' {
		} else if r == 'o' {
			for i := 0; i < n; i++ {
				g.cells[cur_idx+i] = true
			}
		} else {
			return nil, fmt.Errorf("unknown rune: %c", r)
		}
		cur_idx += n
	}
	fmt.Printf("scan done\n")
	if err := s.Err(); err != nil {
		return nil, err
	}
	return g, nil
}
