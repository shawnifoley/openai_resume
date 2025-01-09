#!/usr/bin/env python3

import os
import sys

from markdown2 import markdown
from weasyprint import HTML, CSS


def read_file(filename):
    """Read the content of a file."""
    with open(filename, "r", encoding="utf-8") as file:
        return file.read()


def md_to_html_to_pdf(md_file):
    """Write HTML/PDF content to a file."""

    basename = md_file.replace(".md", "")
    resume_html = basename + ".html"
    resume_pdf = basename + ".pdf"

    extras = ["cuddled-lists", "tables", "footnotes"]
    md_content = read_file(md_file)

    html_content = markdown(md_content, extras=extras)
    with open(resume_html, "w", encoding="utf-8") as file:
        file.write(html_content)
    print("\nResume saved to " + resume_html)
    html = HTML(string=html_content)

    css = [CSS(filename="pdf.css")]
    html.write_pdf(resume_pdf, stylesheets=css)
    print("\nResume saved to " + resume_pdf)

    return


def main():
    if len(sys.argv) != 2:
        print("Usage: python md_html_pdf.py <resume in markdown>")
        sys.exit(1)

    if not os.path.exists(sys.argv[1]):
        raise FileNotFoundError("The file " + sys.argv[1] + " does not exist")

    md_file = sys.argv[1]
    md_to_html_to_pdf(md_file)


if __name__ == "__main__":
    main()
