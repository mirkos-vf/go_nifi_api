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


_v_format = {
    "name": '',
    "path": '',
    "type": '',
    "param": ''
}


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
    parser.epilog = __doc__

    args = parser.parse_args(sysargs[1:])
    _generate_code(args.in_json)
    return 0


def _generate_code(input_json):
    global output_go
    data = json.load(input_json)

    for key, values in data.items():
        if key != "AccessApi":
            output_go = open(key + '.go', mode='w', buffering=-1)

            tmpfile = tempfile.NamedTemporaryFile(suffix='.go', delete=False)
            _write_files(tmpfile, values)
            tmpfile.close()

            new_content = subprocess.check_output(
                ['goimports', tmpfile.name]
            ).decode('utf-8')

            output_go.write(new_content)
            output_go.flush()
            os.remove(tmpfile.name)


def _write_files(outfile, data):
    _fwrite(outfile, """\
     // this documentation
     package go_nifi_api

     import "./types"
     
     // WARNING: This file is generated!

     """)

    for i in data['methods']:
        _format().setdefault("name", '{}{}'.format(_get_method(i["method"]), i["path"].replace('/', '')))
        _format().setdefault("path", i["path"].lower())

        if "body" in i["response"][0]:
            _format().setdefault("type", i["response"][0]["body"])

            if "path" in i["request"][0]:
                param = "{}, ".format(i["request"][0]["param"])

            _fwrite(outfile, """\
            // {name} this godoc
            func (a *app) {name}({param}token, method string) *types.{type} {{
                variables := types.{type}{{}}
                
                url := fmt.Sprintf("%s/%s/%s", a.host, "{path}", "")
                bytes, _ := a.Do(url, token, method, nil)
                _ = json.Unmarshal(bytes, &variables)
    
                return &variables
            }}\n\n
            """.format(**_format()))


def _fwrite(outfile, content):
    print(textwrap.dedent(content), end='', file=outfile)


def _format():
    return _v_format

def _get_method(type):
    return {
        'POST': 'create',
        'GET ': 'get',
        'DELETE': 'delete',
        'PUT': 'put',
    }[type]


if __name__ == '__main__':
    sys.exit(main())
