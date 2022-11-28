# mdpaper

Blazingly fast highly opinionated markdown to pdf 1.5 converter aimed at writing scientific papers, e.g. in University.

## Getting started

### Installation

Either download the latest release from the releases page or install with go:

```bash
go install github.com/sett17/mdpaper
```

## Supported features

- [x] Headings
  - 1-6 `#`
- [x] Paragraph
  - Just text, fallback for everything else
- [x] Generate TOC from headings
  - & pdf outline
- [ ] Generate title page
- [x] bold, italic, inline code inside of paragraph
  - `**bold**`, `*italic*`, \`inline code\`
- [x] `\fill` sequence
  - fills the remaining space in the column with whitespace
- [x] Image
  - `![Text subtitle](path/to/img.png) sizeMultplier`
- [x] Unordered list
  - `- item`
  - nested lists are not currently supported
- [x] Ordered list
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

Options regarding the generation of the PDF can be set in a `config.yaml` file. If this file does not exist, it will be created with default values.

```yaml
text:
  fontSize: 11         # fontsize in pt for regular text
  lineHeight: 1.2      # line height for regular text
  listLineHeight: 1.0  # line height for list entries
page:
  marginTop: 20        # whitespace in mm at the top of the page
  marginBottom: 20     # whitespace in mm at the bottom of the page
  marginHori: 15       # whitespace in mm at the left and right of the page
  columnGap: 7         # whitespace in mm between the columns
  columns: 2           # number of columns (1 or 2)
  pageNumbers: true    # whether to show page numbers
  startPageNumber: 1   # page number to start with
toc:
  enabled: true        # whether to generate a table of contents
  lineHeight: 1.3      # line height for toc entries
  fontSize: 11         # fontsize in pt for toc entries
spaces:
  paragraph: 1         # mm of whitespace after a paragraph
  heading: 1           # mm of whitespace after a heading
paper:
  title: Paper         # title of the paper (also name of pdf file)
  authors: Anonymous   # author of the paper
  debug: false         # whether to enable debug mode (no compression)
```

# Acknowledgements & Known Issues

- [goldmark](https://github.com/yuin/goldmark) for the *blazingly fast* markdown parser

- Paragraphs that are split in the beginning may be out of order
  - use '\fill' in the meantime to force a column break and avoid the splitting
