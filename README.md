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

Options regarding the generation of the PDF can be set in the YAML frontmatter of the markdown file.
For example
```yaml
---
title: "My paper"
author:
  - "John Doe"
margin: 10.0
dbg: true
---

# My paper

the rest of the paper...
```

| Option        |                            Description                            | Default     |
|---------------|:-----------------------------------------------------------------:|:------------|
| title         |                        Title of the paper                         | `Paper`     |
| fontSize      |                   The font size to any integer                    | `11`        |
| lineHeight    |            The relative space between line beginnings             | `1.2`       |
| margin        | The left/right margin in mm (top and bottom are 1.5x this number) | `15.0`      |
| columns       |             The number of columns in the page (max 2)             | `2`         |
| authors       |             List of strings that are the Author names             | `Anonymous` |
| toc           |           Weather to generate a Table of Contents page            | 'true'      |
| tocLineHeight |             The line height used for the ToC entries              | '1.3'       |
| dbg           |                 Weather to draw debug rectangles                  | 'false'     |

# Acknowledgements & Known Issues

- [goldmark](https://github.com/yuin/goldmark) for the *blazingly fast* markdown parser

- Paragraphs that are split in the beginning may be out of order
  - use '\fill' in the meantime to force a column break and avoid the splitting
