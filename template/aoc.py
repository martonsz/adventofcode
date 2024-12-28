#!/usr/bin/env python
import argparse
import os


def part1(filename: str) -> None:
  pass


def part2(filename: str) -> None:
  pass


def main():
  env_part = os.getenv("part")
  parser = argparse.ArgumentParser(description="Process input file and part flag.")
  parser.add_argument(
    "-i",
    "--input",
    type=str,
    default="input.txt",
    help="Input file (default: input.txt)",
  )
  group = parser.add_mutually_exclusive_group()
  group.add_argument(
    "-1",
    action="store_const",
    dest="part",
    const="part1",
    help="Set part to 'part1'",
  )
  group.add_argument(
    "-2",
    action="store_const",
    dest="part",
    const="part2",
    help="Set part to 'part2'",
  )
  parser.add_argument(
    "-b",
    "--begining",
    action="store_true",
    help="Pring row in the begining. On windows vscode prints metadata in the beginning without a new line",
  )
  args = parser.parse_args()

  part = args.part if args.part else env_part if env_part else "part1"

  if args.begining:
    print("")

  if part == "part1":
    part1(args.input)
  else:
    part2(args.input)


if __name__ == "__main__":
  main()
