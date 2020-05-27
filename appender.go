package jsonappender

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"unicode"
)

const (
	tailCheckLen = 16
)

var (
	arrayEndsObject = regexp.MustCompile("(\\[\\s*)?](\\s*}\\s*)$")
	justArray       = regexp.MustCompile("(\\[\\s*)?](\\s*)$")
)

type jsonAppender struct {
	f               *os.File
	strippedBracket bool
	needsComma      bool
	tail            []byte
}

func (a jsonAppender) Write(b []byte) (int, error) {
	trimmed := 0
	if !a.strippedBracket {
		t := bytes.TrimLeftFunc(b, unicode.IsSpace)
		if len(t) == 0 {
			return len(b), nil
		}
		if t[0] != '[' {
			return 0, errors.New("not appending array: " + string(t))
		}
		trimmed = len(b) - len(t) + 1
		b = t[1:]
		a.strippedBracket = true
	}
	if a.needsComma {
		a.needsComma = false
		n, err := a.f.Write([]byte(", "))
		if err != nil {
			return n, err
		}
	}
	n, err := a.f.Write(b)
	return trimmed + n, err
}

func (a jsonAppender) Close() error {
	if _, err := a.f.Write(a.tail); err != nil {
		defer a.f.Close()
		return err
	}
	return a.f.Close()
}

// JSONAppender Helper function to append json array
func JSONAppender(file string) (io.WriteCloser, error) {
	err := checkFile(file)
	if err != nil {
		return nil, err
	}
	
	f, err := os.OpenFile(file, os.O_RDWR, 0664)
	if err != nil {
		return nil, err
	}

	pos, err := f.Seek(0, io.SeekEnd)
	if err != nil {
		return nil, err
	}

	if pos < tailCheckLen {
		pos = 0
	} else {
		pos -= tailCheckLen
	}
	_, err = f.Seek(pos, io.SeekStart)
	if err != nil {
		return nil, err
	}

	tail, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	hasElements := false

	if len(tail) == 0 {
		_, err = f.Write([]byte("["))
		if err != nil {
			return nil, err
		}
	} else {
		var g [][]byte
		if g = arrayEndsObject.FindSubmatch(tail); g != nil {
		} else if g = justArray.FindSubmatch(tail); g != nil {
		} else {
			return nil, errors.New("does not end with array")
		}

		hasElements = len(g[1]) == 0
		_, err = f.Seek(-int64(len(g[2])+1), io.SeekEnd) // 1 for ]
		if err != nil {
			return nil, err
		}
		tail = g[2]
	}

	return jsonAppender{f: f, needsComma: hasElements, tail: tail}, nil
}

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}