package utils

import "bufio"

type LineIterator struct {
	scanner *bufio.Scanner
	current string
	done    bool
}

func NewLineIterator(scanner *bufio.Scanner) *LineIterator {
	scanner.Split(bufio.ScanLines)
	return &LineIterator{scanner: scanner}
}

func (it *LineIterator) Next() bool {
	if it.done {
		return false
	}
	if it.scanner.Scan() {
		it.current = it.scanner.Text()
		return true
	}
	it.done = true
	return false
}

func (it *LineIterator) Value() string {
	return it.current
}