package influxdb

const (
	userName string = "admin"
	password string = "FlbY3CD6mcFUfZvb"
	Test string = "test"
	Caicloud string = "caicloud"
)

type serie struct {
	Name   string `json:"name,omitempty"`
	Column []string `json:"column,omitempty"`
	Values []([]interface{}) `json:"values,omitempty"`
}
type result struct {
	series []serie `json:"series,omitempty"`
}
type results []result
