package srt

import (
	"fmt"
	"strings"
	"testing"
)

func TestScan(t *testing.T) {
	srt := `
1
00:00:13,920 --> 00:00:14,003
AAA
BBB

2
00:00:14,003 --> 00:00:18,382
CCC
DDD
`

	scanner := NewScanner(strings.NewReader(srt))
	actual := "\n"
	for scanner.Scan() {
		actual += fmt.Sprintf("%+v\n", *scanner.Record())
	}
	expected := `
{Seq:1 From:13.92s To:14.003s Text:AAA BBB}
{Seq:2 From:14.003s To:18.382s Text:CCC DDD}
`
	if actual != expected {
		t.Fatalf("expect\n%s\ngot\n%s", expected, actual)
	}
}
