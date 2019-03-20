from __future__ import print_function, unicode_literals

import argparse
import json
import os
import subprocess
import sys
import tempfile
import textwrap


class _FancyFormatter(argparse.ArgumentDefaultsHelpFormatter,
                      argparse.RawDescriptionHelpFormatter):
    pass


def main(sysargs=sys.argv[:]):
    parser = argparse.ArgumentParser(
        description='Generate code!',
        formatter_class=_FancyFormatter)

    parser.add_argument(
        '-i', '--in-json',
        type=argparse.FileType('r'),
        default=sys.stdin,
        help='Input JSON file which defines each type to be generated'
    )
    parser.add_argument(
        '-o', '--out-go',
        type=argparse.FileType('w'),
        default=sys.stdout,
        help='Output file/stream to which generated source will be written'
    )
    parser.epilog = __doc__

    args = parser.parse_args(sysargs[1:])
    _generate_code(args.out_go, args.in_json)
    return 0


def _generate_code(output_go, input_json):
    data = json.load(input_json)

    tmpfile = tempfile.NamedTemporaryFile(suffix='.go', delete=False)
    _write_files(tmpfile)
    tmpfile.close()

    new_content = subprocess.check_output(
        ['goimports', tmpfile.name]
    ).decode('utf-8')

    print(new_content, file=output_go, end='')
    output_go.flush()
    os.remove(tmpfile.name)


def _write_files(outfile):
    _fwrite(outfile, " 11111 ")


def _fwrite(outfile, content):
    print(textwrap.dedent(content), end='', file=outfile)


if __name__ == '__main__':
    sys.exit(main())
