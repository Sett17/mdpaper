# mdpaper

Highly opinionated markdown to pdf 1.5 converter aimed at writing *scientific* papers, e.g. in University.

## Getting started

### Installation

Either download the latest release from the releases page or install with go:

```bash
go install github.com/sett17/mdpaper
```

## Supported elements

- [x] Headings
  - 1-6 `#`
- [x] Paragraph
  - Just text, fallback for everything else
- [ ] Generate TOC from headings
- [ ] Generate title page
- [x] bold, italic, inline code inside of paragraph
  - `**bold**`, `*italic*`, \`inline code\`
- [x] `\fill` sequence
  - fills the remaining space in the column with whitespace
- [ ] Image
  - `![Text subtitle](path/to/img.png){options}`
- [ ] Unordered list
  - `- item`
- [ ] Ordered list
  - `1. item`
- [ ] Checkboxes
  - `- [x] item`
- [ ] Tables
  - `| col1 | col2 |` etc.

## Usage

Every available option can be set in the markdown file with a YAML frontmatter.

```bash
mdpaper my_paper.md
```

### Options

| Option     |                            Description                            | Default     |
|------------|:-----------------------------------------------------------------:|:------------|
| title      |                        Title of the paper                         | `Paper`     |
| fontSize   |                   The font size to any integer                    | `11`        |
| lineHeight |            The relative space between line beginnings             | `1.2`       |
| margin     | The left/right margin in mm (top and bottom are 1.5x this number) | `15.0`      |
| columns    |             The number of columns in the page (max 2)             | `2`         |
| authors    |             List of strings that are the Author names             | `Anonymous` |

# Acknowledgements

- [goldmark](https://github.com/yuin/goldmark) for the *blazingly fast* markdown parser