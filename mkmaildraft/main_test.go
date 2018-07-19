package maildraft

import (
	"os"
	"strings"
	"testing"
)

func TestMake(t *testing.T) {
	Make(map[string][]string{
		"Subject": {"日本語メールテスト"},
		"To":      {"はやまかおる <hymkor@nyaos.org>"},
	}, strings.NewReader("本文でーす"), os.Stdout)
}
