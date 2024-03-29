package globals

type Config struct {
	Text struct {
		FontSize       int     `yaml:"fontSize"`
		LineHeight     float64 `yaml:"lineHeight"`
		ListLineHeight float64 `yaml:"listLineHeight"`
		ListMarker     string  `yaml:"listMarker"`
		FigureText     string  `yaml:"figureText"`
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
		Heading    string  `yaml:"heading"`
	} `yaml:"toc"`
	Tof struct {
		Enabled    bool    `yaml:"enabled"`
		LineHeight float64 `yaml:"lineHeight"`
		FontSize   int     `yaml:"fontSize"`
		Heading    string  `yaml:"heading"`
	} `yaml:"tof"`
	Margins struct {
		Paragraph     float64 `yaml:"paragraph"`
		HeadingTop    float64 `yaml:"heading_top"`
		HeadingBottom float64 `yaml:"heading_bottom"`
		List          float64 `yaml:"list"`
		Image         float64 `yaml:"image"`
		Code          float64 `yaml:"code"`
	} `yaml:"margins"`
	Paper struct {
		Title  string `yaml:"title"`
		Author string `yaml:"authors"`
		Debug  bool   `yaml:"debug"`
	}
	Citation struct {
		Enabled       bool    `yaml:"enabled"`
		File          string  `yaml:"file"`
		Heading       string  `yaml:"heading"`
		BibFontSize   int     `yaml:"bibFontSize"`
		BibLineHeight float64 `yaml:"bibLineHeight"`
		CSLFile       string  `yaml:"csl"`
		LocaleFile    string  `yaml:"locale"`
	} `yaml:"citation"`
	Code struct {
		Style            string  `yaml:"style"`
		FontSize         int     `yaml:"fontSize"`
		CharacterSpacing float64 `yaml:"characterSpacing"`
		LineNumbers      bool    `yaml:"lineNumbers"`
		Dot              bool    `yaml:"dot"`
	} `yaml:"code"`
	Table struct {
		Padding float64 `yaml:"padding"`
	} `yaml:"table"`
	Cover struct {
		Enabled  bool   `yaml:"enabled"`
		Subtitle string `yaml:"subtitle"`
		Abstract string `yaml:"abstract"`
	} `yaml:"cover"`
}

var Default = Config{
	Text: struct {
		FontSize       int     `yaml:"fontSize"`
		LineHeight     float64 `yaml:"lineHeight"`
		ListLineHeight float64 `yaml:"listLineHeight"`
		ListMarker     string  `yaml:"listMarker"`
		FigureText     string  `yaml:"figureText"`
	}{
		FontSize:       11,
		LineHeight:     1.2,
		ListLineHeight: 1.0,
		ListMarker:     "-",
		FigureText:     "Figure",
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
		Heading    string  `yaml:"heading"`
	}{
		Enabled:    true,
		LineHeight: 1.3,
		FontSize:   12,
		Heading:    "Table of Contents",
	},
	Tof: struct {
		Enabled    bool    `yaml:"enabled"`
		LineHeight float64 `yaml:"lineHeight"`
		FontSize   int     `yaml:"fontSize"`
		Heading    string  `yaml:"heading"`
	}{
		Enabled:    true,
		LineHeight: 1.3,
		FontSize:   12,
		Heading:    "Table of Figures",
	},
	Margins: struct {
		Paragraph     float64 `yaml:"paragraph"`
		HeadingTop    float64 `yaml:"heading_top"`
		HeadingBottom float64 `yaml:"heading_bottom"`
		List          float64 `yaml:"list"`
		Image         float64 `yaml:"image"`
		Code          float64 `yaml:"code"`
	}{
		Paragraph:     2.0,
		HeadingTop:    3.0,
		HeadingBottom: 1.0,
		List:          2.0,
		Image:         4.0,
		Code:          2.0,
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
		Enabled       bool    `yaml:"enabled"`
		File          string  `yaml:"file"`
		Heading       string  `yaml:"heading"`
		BibFontSize   int     `yaml:"bibFontSize"`
		BibLineHeight float64 `yaml:"bibLineHeight"`
		CSLFile       string  `yaml:"csl"`
		LocaleFile    string  `yaml:"locale"`
	}{
		Enabled:       true,
		File:          "citations.json",
		Heading:       "References",
		BibFontSize:   12,
		BibLineHeight: 1.4,
		CSLFile:       "",
		LocaleFile:    "",
	},
	Table: struct {
		Padding float64 `yaml:"padding"`
	}{
		Padding: 1.5,
	},
	Code: struct {
		Style            string  `yaml:"style"`
		FontSize         int     `yaml:"fontSize"`
		CharacterSpacing float64 `yaml:"characterSpacing"`
		LineNumbers      bool    `yaml:"lineNumbers"`
		Dot              bool    `yaml:"dot"`
	}{
		Style:            "dracula",
		FontSize:         10,
		CharacterSpacing: -.75,
		LineNumbers:      true,
		Dot:              true,
	},
	Cover: struct {
		Enabled  bool   `yaml:"enabled"`
		Subtitle string `yaml:"subtitle"`
		Abstract string `yaml:"abstract"`
	}{
		Enabled:  true,
		Subtitle: "generated by mdpaper",
		Abstract: "",
	},
}

var Cfg = Default

var DidConfig = false
