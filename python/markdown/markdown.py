import re


def process_paragraph(line: str) -> str:
    if not re.match("<h|<ul|<p|<li", line):
        line = f"<p>{line}</p>"
    return line


def process_italics(line: str) -> str:
    return re.sub(r"_(.*?)_", r"<em>\1</em>", line)


def process_bold(line: str) -> str:
    return re.sub(r"__(.*?)__", r"<strong>\1</strong>", line)


def process_header(line: str) -> str:
    match = re.match("^(#{1,6}) (.*)", line)
    if match is not None:
        headers, text = match.groups()
        size = len(headers)
        line = f"<h{size}>{line[size+1:]}</h{size}>"
    return line


def process_unordered_list(line: str, in_list: bool, in_list_append: bool) -> tuple:
    m = re.match(r"\* (.*)", line)
    if m:
        if not in_list:
            in_list = True
            curr = m.group(1)
            m1 = re.match("(.*)__(.*)__(.*)", curr)
            if m1:
                curr = (
                    m1.group(1) + "<strong>" + m1.group(2) + "</strong>" + m1.group(3)
                )
            m1 = re.match("(.*)_(.*)_(.*)", curr)
            if m1:
                curr = m1.group(1) + "<em>" + m1.group(2) + "</em>" + m1.group(3)
            line = "<ul><li>" + curr + "</li>"
        else:
            is_bold = False
            is_italic = False
            curr = m.group(1)
            m1 = re.match("(.*)__(.*)__(.*)", curr)
            if m1:
                is_bold = True
            m1 = re.match("(.*)_(.*)_(.*)", curr)
            if m1:
                is_italic = True
            if is_bold:
                curr = (
                    m1.group(1) + "<strong>" + m1.group(2) + "</strong>" + m1.group(3)
                )
            if is_italic:
                curr = m1.group(1) + "<em>" + m1.group(2) + "</em>" + m1.group(3)
            line = "<li>" + curr + "</li>"
    else:
        if in_list:
            in_list_append = True
            in_list = False

    return line, in_list, in_list_append


def parse(markdown: str) -> str:
    lines = markdown.split("\n")
    res = ""
    in_list = False
    in_list_append = False
    for line in lines:
        line = process_header(line)
        line, in_list, in_list_append = process_unordered_list(
            line, in_list, in_list_append
        )
        line = process_paragraph(line)
        line = process_bold(line)
        line = process_italics(line)

        if in_list_append:
            line = "</ul>" + line
            in_list_append = False
        res += line
    if in_list:
        res += "</ul>"
    return res
