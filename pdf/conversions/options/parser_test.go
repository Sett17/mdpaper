package options

import (
	"reflect"
	"testing"
)

func TestParseOptions(t *testing.T) {
	type args struct {
		options string
	}
	tests := []struct {
		name    string
		args    args
		want    *Options
		wantErr bool
	}{
		{"Single", args{"id='moin'"}, &Options{options: []Option{{Name: "id", Value: "moin"}}}, false},
		{"Mixed", args{"width=0.5 lineNumbers=false id='moin'"}, &Options{options: []Option{{Name: "width", Value: 0.5}, {Name: "lineNumbers", Value: false}, {Name: "id", Value: "moin"}}}, false},
		{"Many spaces", args{"width=0.5     lineNumbers=false   id='moin' "}, &Options{options: []Option{{Name: "width", Value: 0.5}, {Name: "lineNumbers", Value: false}, {Name: "id", Value: "moin"}}}, false},
		{"Empty", args{""}, &Options{}, false},
		{"Invalid", args{"width=0.5 lineNumbers=false id=hello world"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseOption(t *testing.T) {
	type args struct {
		option string
	}
	tests := []struct {
		name    string
		args    args
		want    Option
		wantErr bool
	}{
		{"String", args{"id='moin'"}, Option{"id", "moin"}, false},
		{"String with spaces", args{"id='moin moin'"}, Option{"id", "moin moin"}, false},
		{"Float", args{"width=0.5"}, Option{"width", 0.5}, false},
		{"Int", args{"width=5"}, Option{"width", 5.0}, false},
		{"Bool", args{"lineNumbers=false"}, Option{"lineNumbers", false}, false},
		{"Invalid", args{"hello"}, Option{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseOption(tt.args.option)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseOption() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseOption() got = %v, want %v", got, tt.want)
			}
		})
	}
}
