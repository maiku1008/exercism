import re


def process_paragraph(line: str) -> str:
    if not re.match("<h|<ul|<p|<li", line):
        line = f"<p>{line}</p>"
    return line


def process_italics(line: str) -> str:
    return re.sub(r"_(.*?)_", r"<em>\1</em>", line)


def process_bold(line: str) -> str:
    return re.sub(r"__(.*?)__", r"<strong>\1</strong>", line)


def process_list_item(line: str) -> str:
    return re.sub(r"^\* (.*)", r"<li>\1</li>", line)


def process_header(line: str) -> str:
    match = re.match("^(#{1,6}) (.*)", line)
    if match is not None:
        headers, text = match.groups()
        size = len(headers)
        line = f"<h{size}>{line[size+1:]}</h{size}>"
    return line


def parse(markdown: str) -> str:
    lines = markdown.split("\n")
    parsed = []
    in_list = False
    for line in lines:
        line = process_header(line)
        line = process_list_item(line)
        line = process_paragraph(line)
        line = process_bold(line)
        line = process_italics(line)
        if line.startswith("<li>"):
            if not in_list:
                parsed.append("<ul>")
            in_list = True
        else:
            if in_list:
                parsed.append("</ul>")
            in_list = False
        parsed.append(line)
    if in_list:
        parsed.append("</ul>")

    return "".join(parsed)
