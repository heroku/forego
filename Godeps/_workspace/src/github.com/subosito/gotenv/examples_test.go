package gotenv_test

import (
	"fmt"
	"strings"

	"github.com/heroku/forego/Godeps/_workspace/src/github.com/subosito/gotenv"
)

func ExampleParse() {
	pairs := gotenv.Parse(strings.NewReader("FOO=test\nBAR=$FOO"))
	fmt.Printf("%+v\n", pairs) // gotenv.Env{"FOO": "test", "BAR": "test"}

	pairs = gotenv.Parse(strings.NewReader(`FOO="bar"`))
	fmt.Printf("%+v\n", pairs) // gotenv.Env{"FOO": "bar"}
}
