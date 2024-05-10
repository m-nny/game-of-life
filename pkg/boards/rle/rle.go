package rle

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"minmax.uk/game-of-life/pkg/boards"
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

func Parse(rd io.Reader) (boards.BoardSpec, error) {
	var board boards.BoardSpec
	if _, err := fmt.Fscanf(rd, "x = %d, y = %d\n", &board.Cols, &board.Rows); err != nil {
		return board, err
	}
	scanner := bufio.NewScanner(rd)
	scanner.Split(rleSplitter)
	str := ""
	for scanner.Scan() {
		token := scanner.Text()
		var n int
		var r rune
		if _, err := fmt.Sscanf(token, "%d%c", &n, &r); err != nil {
			return board, err
		}
		if r == 'b' {
			//empty
			str += strings.Repeat(".", n)
		} else if r == 'o' {
			str += strings.Repeat("O", n)
		} else {
			return board, fmt.Errorf("unknown rune: %c", r)
		}
	}
	board.Str = str
	if err := scanner.Err(); err != nil {
		return board, err
	}
	return board, nil
}
