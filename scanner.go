package srt

import (
	"bufio"
	"io"
	"strconv"
	"strings"
	"time"
)

type (
	Scanner struct {
		s    *bufio.Scanner
		text string
		rec  Record
	}
	Record struct {
		Seq  int
		From time.Duration
		To   time.Duration
		Text string
	}
)

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{s: bufio.NewScanner(r)}

}

func (s *Scanner) Scan() bool {
	text := ""
	for s.s.Scan() {
		text = strings.TrimSpace(s.s.Text())
		if text != "" {
			break
		}
	}
	if text == "" {
		return false
	}
	seq, _ := strconv.Atoi(text)
	if !s.s.Scan() {
		return false
	}
	fromTo := strings.Split(s.s.Text(), "-->")
	var from, to time.Duration
	if len(fromTo) > 0 {
		from = parseDuration(fromTo[0])
	}
	if len(fromTo) > 1 {
		to = parseDuration(fromTo[1])
	}
	if !s.s.Scan() {
		return false
	}
	s.rec = Record{
		Seq:  seq,
		From: from,
		To:   to,
		Text: scanMultiLine(s.s),
	}
	return true
}

func scanMultiLine(s *bufio.Scanner) string {
	t := s.Text()
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" {
			break
		}
		t += " " + line
	}
	return t
}

var zeroTime = time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)

func parseDuration(s string) time.Duration {
	s = strings.TrimSpace(s)
	s = strings.Replace(s, ",", ".", -1)
	t, err := time.Parse("15:04:05.000", s)
	if err != nil {
		return 0
	}
	return t.Sub(zeroTime)
}

func (s *Scanner) Record() *Record {
	return &s.rec
}
