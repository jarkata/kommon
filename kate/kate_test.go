package kate_test

import (
	"fmt"
	"testing"

	"github.com/jarkata/kommon/kate"
)

func TestParseDate(t *testing.T) {
	dates := kate.ParseIsoDateTime("2023-01-01 10:10:10")
	fmt.Println(dates)
}

func TestParseIsoDateTime2(t *testing.T) {
	t2 := kate.ParseIsoDateTime2("2023-01-01'T'10:10:10")
	fmt.Println(t2)
}
