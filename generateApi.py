import sys
import argparse


def main(sysargs=sys.argv[:]):
    parser = argparse.ArgumentParser(description='Generate api!')

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

    args = parser.parse_args(sysargs[1:])

    return 0


if __name__ == '__main__':
    sys.exit(main())
