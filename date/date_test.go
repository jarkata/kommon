package date_test

import (
	"fmt"
	"testing"

	"github.com/jarkata/kommon/date"
)

func TestParseDate(t *testing.T) {
	dates := date.ParseIsoDateTime("2023-01-01 10:10:10")
	fmt.Println(dates)
}

func TestParseIsoDateTime2(t *testing.T) {
	t2 := date.ParseIsoDateTime2("2023-01-01'T'10:10:10")
	fmt.Println(t2)
}
