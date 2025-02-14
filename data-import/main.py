#!/usr/bin/env python3

import argparse
import sys
from saxonche import PySaxonProcessor


def transform_xml_to_json(xml_file, xslt_file):
    with PySaxonProcessor(license=False) as proc:
        xslt_proc = proc.new_xslt30_processor()
        result = xslt_proc.transform_to_string(source_file=xml_file, stylesheet_file=xslt_file)
        return result


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("xml")
    parser.add_argument("xslt")
    args = parser.parse_args()

    try:
        json_output = transform_xml_to_json(args.xml, args.xslt)
        print(json_output)
    except Exception as e:
        print("An error occurred during transformation:", str(e), file=sys.stderr)
        sys.exit(1)


if __name__ == '__main__':
    main()
