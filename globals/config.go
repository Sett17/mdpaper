package globals

type Config struct {
	Text struct {
		FontSize       int     `yaml:"fontSize"`
		LineHeight     float64 `yaml:"lineHeight"`
		ListLineHeight float64 `yaml:"listLineHeight"`
	} `yaml:"text"`
	Page struct {
		MarginTop       float64 `yaml:"marginTop"`
		MarginBottom    float64 `yaml:"marginBottom"`
		MarginHori      float64 `yaml:"marginHori"`
		ColumnGap       float64 `yaml:"columnGap"`
		Columns         int     `yaml:"columns"`
		PageNumbers     bool    `yaml:"pageNumbers"`
		StartPageNumber int     `yaml:"startPageNumber"`
	} `yaml:"page"`
	Toc struct {
		Enabled    bool    `yaml:"enabled"`
		LineHeight float64 `yaml:"lineHeight"`
		FontSize   int     `yaml:"fontSize"`
	} `yaml:"toc"`
	Spaces struct {
		Paragraph float64 `yaml:"paragraph"`
		Heading   float64 `yaml:"heading"`
	}
	Paper struct {
		Title  string `yaml:"title"`
		Author string `yaml:"authors"`
		Debug  bool   `yaml:"debug"`
	}
	Citation struct {
		Enabled bool   `yaml:"enabled"`
		File    string `yaml:"file"`
	}
}

var Default = Config{
	Text: struct {
		FontSize       int     `yaml:"fontSize"`
		LineHeight     float64 `yaml:"lineHeight"`
		ListLineHeight float64 `yaml:"listLineHeight"`
	}{
		FontSize:       11,
		LineHeight:     1.2,
		ListLineHeight: 1.0,
	},
	Page: struct {
		MarginTop       float64 `yaml:"marginTop"`
		MarginBottom    float64 `yaml:"marginBottom"`
		MarginHori      float64 `yaml:"marginHori"`
		ColumnGap       float64 `yaml:"columnGap"`
		Columns         int     `yaml:"columns"`
		PageNumbers     bool    `yaml:"pageNumbers"`
		StartPageNumber int     `yaml:"startPageNumber"`
	}{
		MarginTop:       20,
		MarginBottom:    20,
		MarginHori:      15,
		ColumnGap:       7,
		Columns:         2,
		PageNumbers:     true,
		StartPageNumber: 1,
	},
	Toc: struct {
		Enabled    bool    `yaml:"enabled"`
		LineHeight float64 `yaml:"lineHeight"`
		FontSize   int     `yaml:"fontSize"`
	}{
		Enabled:    true,
		LineHeight: 1.3,
		FontSize:   11,
	},
	Spaces: struct {
		Paragraph float64 `yaml:"paragraph"`
		Heading   float64 `yaml:"heading"`
	}{
		Paragraph: 2.0,
		Heading:   2.0,
	},
	Paper: struct {
		Title  string `yaml:"title"`
		Author string `yaml:"authors"`
		Debug  bool   `yaml:"debug"`
	}{
		Title:  "Paper",
		Author: "Anonymous",
		Debug:  false,
	},
	Citation: struct {
		Enabled bool   `yaml:"enabled"`
		File    string `yaml:"file"`
	}{
		Enabled: true,
		File:    "citations.bib",
	},
}

var Cfg Config = Default
