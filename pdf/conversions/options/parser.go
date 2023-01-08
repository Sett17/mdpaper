package options

import (
	"fmt"
	"strconv"
	"strings"
)

// Option represents a single option with a name and value.
type Option struct {
	Name  string
	Value interface{}
}

func (o Option) String() string {
	return fmt.Sprintf("{%s=%v}", o.Name, o.Value)
}

func (o Option) Empty() bool {
	return o.Name == ""
}

// Options represents a collection of options.
type Options struct {
	options []Option
}

func (o *Options) String() string {
	ret := "["
	for i, opt := range o.options {
		ret += opt.String()
		if i < len(o.options)-1 {
			ret += " "
		}
	}
	ret += "]"
	return ret
}

func (o *Options) GetInt(name string) (int, bool) {
	f, ok := o.GetFloat(name)
	if !ok {
		return 0, false
	}
	if f == float64(int(f)) {
		return int(f), ok
	}
	return 0, false
}

func (o *Options) GetFloat(name string) (float64, bool) {
	if o == nil && o.options == nil {
		goto exit
	}
	for _, opt := range o.options {
		if opt.Name == name {
			if value, ok := opt.Value.(float64); ok {
				return value, true
			}
			return 0, false
		}
	}
exit:
	return 0, false
}

func (o *Options) GetBool(name string) (bool, bool) {
	if o == nil && o.options == nil {
		goto exit
	}
	for _, opt := range o.options {
		if opt.Name == name {
			if value, ok := opt.Value.(bool); ok {
				return value, true
			}
			return false, false
		}
	}
exit:
	return false, false
}

func (o *Options) GetString(name string) (string, bool) {
	if o == nil && o.options == nil {
		goto exit
	}
	for _, opt := range o.options {
		if opt.Name == name {
			if value, ok := opt.Value.(string); ok {
				return value, true
			}
			return "", false
		}
	}
exit:
	return "", false
}

// parseOption parses a single option in the form "name=value" and returns an Option structure.
func parseOption(option string) (Option, error) {
	if option == "" {
		return Option{}, nil
	}
	parts := strings.SplitN(option, "=", 2)
	if len(parts) != 2 {
		return Option{}, fmt.Errorf("invalid option: %s", option)
	}
	name := parts[0]
	value := parts[1]
	if strings.HasPrefix(value, "'") || strings.HasPrefix(value, "\"") {
		// Option value is a string
		return Option{Name: name, Value: value[1 : len(value)-1]}, nil
	} else if strings.EqualFold(value, "true") || strings.EqualFold(value, "false") {
		// Option value is a boolean
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return Option{}, err
		}
		return Option{Name: name, Value: boolValue}, nil
	} else {
		// Option value is a number
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return Option{}, err
		}
		return Option{Name: name, Value: floatValue}, nil
	}
}

// Parse parses a string of options in the form "name=value name=value" and returns an Options structure.
func Parse(options string) (*Options, error) {
	options = strings.TrimSpace(options)
	var parsedOptions []Option
	currentOption := ""
	insideQuotes := false
	for _, c := range options {
		if c == '\'' || c == '"' {
			insideQuotes = !insideQuotes
		}
		if c == ' ' && !insideQuotes {
			// End of current option
			opt, err := parseOption(currentOption)
			if opt.Empty() {
				continue
			}
			if err != nil {
				return &Options{}, err
			}
			parsedOptions = append(parsedOptions, opt)
			currentOption = ""
			continue
		}
		currentOption += string(c)
	}
	if currentOption != "" {
		// Add final option
		opt, err := parseOption(currentOption)
		if err != nil {
			return &Options{}, err
		}
		parsedOptions = append(parsedOptions, opt)
	}
	return &Options{options: parsedOptions}, nil
}

func Extract(s string) string {
	i := strings.Index(s, "[")
	if i == -1 {
		return ""
	}

	j := strings.Index(s, "]")
	if j == -1 {
		return ""
	}

	return s[i+1 : j]
}
