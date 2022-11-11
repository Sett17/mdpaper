package globals

import "mdpaper/pdf/spec"

type Config struct {
	FontSize      int
	LineHeight    float64
	Margin        float64
	Columns       int
	ToC           bool
	ToCLineHeight float64
	Title         string
	Authors       []string
}

var DefaultConfig = Config{
	FontSize:      11,
	LineHeight:    1.2,
	Margin:        spec.MmToPt(15),
	Columns:       2,
	ToC:           true,
	ToCLineHeight: 1.3,
	Title:         "Paper",
	Authors:       []string{"Anonymous"},
}

var Cfg Config

var File []byte

func FromMap(m map[string]interface{}) Config {
	c := DefaultConfig
	for k, v := range m {
		switch k {
		case "fontSize":
			switch v.(type) {
			case int:
				c.FontSize = v.(int)
			default:
				panic("invalid type for fontSize")
			}
		case "lineHeight":
			switch v.(type) {
			case float64:
				c.LineHeight = v.(float64)
			case int:
				c.LineHeight = float64(v.(int))
			}
		case "margin":
			switch v.(type) {
			case int:
				c.Margin = spec.MmToPt(float64(v.(int)))
			case float64:
				c.Margin = spec.MmToPt(v.(float64))
			}
		case "columns":
			c.Columns = v.(int)
			if c.Columns < 1 {
				c.Columns = 1
			} else if c.Columns > 2 {
				c.Columns = 2
			}
		case "toc":
			c.ToC = v.(bool)
		case "tocStretch":
			switch v.(type) {
			case int:
				c.ToCLineHeight = float64(v.(int))
			case float64:
				c.ToCLineHeight = v.(float64)
			}
		case "title":
			c.Title = v.(string)
		case "authors":
			c.Authors = []string{}
			switch v.(type) {
			case []string:
				c.Authors = v.([]string)
			case []interface{}:
				for _, a := range v.([]interface{}) {
					c.Authors = append(c.Authors, a.(string))
				}
			case string:
				c.Authors = []string{v.(string)}
			}
		}
	}
	return c
}
