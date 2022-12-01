# mdpaper

*Blazingly* fast highly opinionated markdown to pdf 1.5 converter aimed at writing scientific papers, e.g. in University.

## Getting started

### Installation

Either download the latest release from the releases page or install with go:

```bash
go install github.com/sett17/mdpaper
```

## Usage

Every available option can be set in the markdown file with a YAML frontmatter.

```bash
mdpaper my_paper.md
```

## Options

Options regarding the generation of the PDF can be set in a `config.yaml` file. If this file does not exist, it will be created with default values.

```yaml
text:
  fontSize: 11                # fontsize in pt for regular text
  lineHeight: 1.2             # line height for regular text
  listLineHeight: 1.0         # line height for list entries
page:
  marginTop: 20               # whitespace in mm at the top of the page
  marginBottom: 20            # whitespace in mm at the bottom of the page
  marginHori: 15              # whitespace in mm at the left and right of the page
  columnGap: 7                # whitespace in mm between the columns
  columns: 2                  # number of columns (1 or 2)
  pageNumbers: true           # whether to show page numbers
  startPageNumber: 1          # page number to start with
toc:
  enabled: true               # whether to generate a table of contents
  lineHeight: 1.3             # line height for toc entries
  fontSize: 11                # fontsize in pt for toc entries
  heading: Table of Contents  # string to use as heading for the toc
spaces:
  paragraph: 1                # whitespace in mm after a paragraph
  heading: 1                  # whitespace in mm after a heading
paper:
  title: Paper                # title of the paper (also name of pdf file)
  authors: Anonymous          # author of the paper
  debug: false                # whether to enable debug mode (no compression)
  mermaid: false              # whether to enable mermaid support (requires mermaid-cli on your machine)
citation:
  enabled: true               # whether to enable citation support
  file: citations.bib         # path to the bibtex file
  heading: References         # string to use as heading for the references
```

### Mermaid

> Mermaid lets you create diagrams and visualizations using text and code.
> 
> It is a JavaScript based diagramming and charting tool that renders Markdown-inspired text definitions to create and modify diagrams dynamically.

You need to install the [mermaid-cli](https://github.com/mermaid-js/mermaid-cli) so mdpaper can use the `mmdc` command.

*For latest instructions see the [mermaid-cli](https://github.com/mermaid-js/mermaid-cli) repository.*
```bash
npm install -g @mermaid-js/mermaid-cli
```

## Currently unsupported commonly used markdown features

If used, they will be ignored when producing the pdf.

- Tables
- Nested lists
- Code blocks
- Strike-through

# Acknowledgements & Known Issues

Thanks to the people behind:
- [goldmark](https://github.com/yuin/goldmark) for the markdown parser
- [bibtex](https://github.com/nickng/bibtex) for the bibtex parser
- [mermaid](https://github.com/mermaid-js/mermaid) for great looking diagrams


- Paragraphs that are split in the beginning may be out of order
  - use '\fill' in the meantime to force a column break and avoid the splitting
