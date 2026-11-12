package actioninfo

import (
	"errors"
	"testing"
)

type failure_parser struct {
	parse_calls int
	info_calls  int
}

func (p *failure_parser) Parse(string) error {
	p.parse_calls++
	if p.parse_calls == 1 {
		return errors.New("parse failure")
	}
	return nil
}

func (p *failure_parser) ActionInfo() (string, error) {
	p.info_calls++
	if p.info_calls == 1 {
		return "", errors.New("info failure")
	}
	return "ok", nil
}

func TestInfoContinuesAfterErrors(t *testing.T) {
	parser := &failure_parser{}
	Info([]string{"bad parse", "bad activity", "valid"}, parser)
	if parser.parse_calls != 3 || parser.info_calls != 2 {
		t.Fatalf("calls: Parse=%d, ActionInfo=%d", parser.parse_calls, parser.info_calls)
	}
}
