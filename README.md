# mdpaper

*Blazingly* fast highly opinionated markdown to pdf 1.5 converter aimed at writing scientific papers, e.g. in University.

## Getting started

### Installation

Either download the latest release from the releases page or install with go:

```bash
go install github.com/sett17/mdpaper@latest
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
  fontSize: 11                   # fontsize in pt for regular text
  lineHeight: 1.2                # line height for regular text
  listLineHeight: 1.0            # line height for list entries
page:
  marginTop: 20                  # whitespace in mm at the top of the page
  marginBottom: 20               # whitespace in mm at the bottom of the page
  marginHori: 15                 # whitespace in mm at the left and right of the page
  columnGap: 7                   # whitespace in mm between the columns
  columns: 2                     # number of columns (1 or 2)
  pageNumbers: true              # whether to show page numbers
  startPageNumber: 1             # page number to start with
toc:
  enabled: true                  # whether to generate a table of contents
  lineHeight: 1.3                # line height for toc entries
  fontSize: 11                   # fontsize in pt for toc entries
  heading: Table of Contents     # string to use as heading for the toc
spaces:
  paragraph: 2                   # whitespace in mm after a paragraph
  heading: 2                     # whitespace in mm after a heading
  code: 2                        # whitespace in mm after a code block
paper:
  title: Paper                   # title of the paper (will also be the file name)
  authors: Anonymous             # author of the paper
  debug: false                   # whether enable debug mode (uncompressed pdf)
citation:
  enabled: true                  # whether to enable citation support
  file: citations.bib            # path to the bib file
  heading: References            # string to use as heading for the references
code:
  style: dracula                 # style to use for code blocks
  fontSize: 10                   # fontsize in pt for code blocks
  characterSpacing: -0.75        # character spacing for code blocks
  lineNumbers: true              # whether to show line numbers
  mermaid: false                 # whether to enable mermaid diagram support
cover:
  enabled: true                  # whether to generate a cover page
  subtitle: generated by mdpaper # subtitle to use on the cover page
```

## _Beta_ Features

The features below, still have some kinks to work out, or configurations to add.

### Code

Highlighting support for code blocks is provided by [chroma](https://github.com/alecthomas/chroma). So all the styles available there can be used here.

* lines are only broken at newlines in the code (your job to not make them oveflow)
* currently no support for splitting over multiple columns
  * also means that code blocks longer than a column won't work

### Mermaid

> Mermaid lets you create diagrams and visualizations using text and code.
>
> It is a JavaScript based diagramming and charting tool that renders Markdown-inspired text definitions to create and modify diagrams dynamically.

You need to install the [mermaid-cli](https://github.com/mermaid-js/mermaid-cli) so mdpaper can use the `mmdc` command.

*For latest instructions see the [mermaid-cli](https://github.com/mermaid-js/mermaid-cli) repository.*

```bash
npm install -g @mermaid-js/mermaid-cli
```

* needs extra installation
* slow (because external JS...)
  * workaround: disabling when not working on parts including diagrams
* currently the size can't be changed

### Math

_This could currently even be considered alpha, only the regular use case works fine_

Currently only math-blocks are supported. E.g.

```latex
$$
f(n) \in O(g(n)) \Leftrightarrow \lim_{n\rightarrow\infty} sup \frac{|f(n)|}{|g(n)|}<\infty
$$
```

and not inline math (e.g. `$f(n) \in O(g(n))$`). Look in the demo for correct usage.

This is implemented by using an existing latex installation on your machine that create a PNG from the given equation. The commands needed here are `latex` and `dvipng`. These should already be available if you have any latex installation. If latex is new to you, you can use [TexLive](https://www.tug.org/texlive/) to install it.

* needs extra installation
* kinda of slow (better than mermaid)
* Use images (so no copy-pasting from the pdf)
* size can be wonky
* no support for inline math
* no support for multiple equations in one block
* latex...

## Currently unsupported commonly used markdown features

If used, they will be ignored when producing the pdf.

- Tables
- Nested lists
- Strike-through

# Acknowledgements & Known Issues

Thanks to the people behind:

- [goldmark](https://github.com/yuin/goldmark) for the markdown parser
- [bibtex](https://github.com/nickng/bibtex) for the bibtex parser
- [mermaid](https://github.com/mermaid-js/mermaid) for great looking diagrams
- [chroma](https://github.com/alecthomas/chroma) for the code highlighting
- [Furqan Software](https://github.com/FurqanSoftware/goldmark-katex) for goldmark math extender code


- Paragraphs that are split in the beginning may be out of order
  - use '\fill' in the meantime to force a column break and avoid the splitting
