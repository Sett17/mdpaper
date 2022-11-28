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
		listItem  float64 `yaml:"listItem"`
		Image     float64 `yaml:"image"`
	}
	Paper struct {
		Title  string `yaml:"title"`
		Author string `yaml:"authors"`
		Debug  bool   `yaml:"debug"`
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
		listItem  float64 `yaml:"listItem"`
		Image     float64 `yaml:"image"`
	}{
		Paragraph: 1.0,
		Heading:   1.0,
		listItem:  0,
		Image:     3.0,
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
}

var Cfg Config = Default
