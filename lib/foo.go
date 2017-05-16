package libfoo

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

func Foo() (int, error) {
	return *kingpin.Flag("foo", "foo flag").Int(), nil
	return 0, fmt.Errorf("foo")
}
