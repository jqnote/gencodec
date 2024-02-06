package immutable

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	data := `{
    "name": "nnnnn",
    "js_age": "aaaaaa",
    "embed": {
        "name": "yyyname",
        "age": "yage"
    },
    "ptr": {
        "name": "yyyname",
        "age": "yage"
    }
}`
	var x = X{}
	err := json.Unmarshal([]byte(data), &x)
	fmt.Println(err)

	fmt.Printf("%v \n%v \n%v", x, x.embed, x.ptr)
}
