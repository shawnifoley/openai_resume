#!/usr/bin/env python3

import os
import sys

import mistune
import pdfkit


def read_file(filename):
    """Read the content of a file."""
    with open(filename, "r", encoding="utf-8") as file:
        return file.read()


def convert_markdown_to_html(markdown_content):
    """Convert Markdown content to HTML."""
    return mistune.markdown(markdown_content)


def write_html_file(html_content, output_file):
    """Write HTML content to a file."""
    with open(output_file, "w", encoding="utf-8") as file:
        file.write(html_content)
    print("\nResume saved to " + output_file)


def convert_html_to_pdf(input_html, output_pdf):
    """Write PDF content to a file."""
    pdfkit.from_string(input_html, output_pdf)
    print("\nResume saved to " + output_pdf)


def main():
    if len(sys.argv) != 2:
        print("Usage: python md_html_pdf.py <resume in markdown>")
        sys.exit(1)

    if not os.path.exists(sys.argv[1]):
        raise FileNotFoundError("The file " + sys.argv[1] + " does not exist")

    resume_content = read_file(sys.argv[1])
    resume_html = sys.argv[1] + ".html"
    resume_pdf = sys.argv[1] + ".pdf"
    html_content = convert_markdown_to_html(resume_content)
    write_html_file(html_content, resume_html)
    convert_html_to_pdf(html_content, resume_pdf)


if __name__ == "__main__":
    main()
