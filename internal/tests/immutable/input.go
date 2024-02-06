package immutable

//go:generate go run github.com/fjl/gencodec -type X -field-override Xo -formats json,yaml,toml -out x_immutable.go -immutable
//go:generate go run github.com/fjl/gencodec -type Y -formats json,yaml,toml -out y_immutable.go -immutable

type X struct {
	name  string
	age   string `json:"js_age" toml:"tm_age"`
	embed Y
	ptr   *Y
}

type Xo struct {
}

type Y struct {
	name string
	age  string
}
