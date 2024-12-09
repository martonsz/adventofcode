#!/usr/bin/env python
import argparse
import os
import regex

pattern = regex.compile(r"XMAS|SAMX")
pattern2 = regex.compile(r"MS|SM")

def part1(filename: str) -> None:
  with open(filename, "r") as file:
    text = [line.strip() for line in file.readlines()]
    max_col = len(text[0])
    max_row = len(text)
    cols = [[] for _ in range(max_col)]
    rows = [[] for _ in range(max_row)]
    fdiag = [[] for _ in range(max_row + max_col - 1)]
    bdiag = [[] for _ in range(len(fdiag))]
    min_bdiag = -max_row + 1

    for x in range(max_col):
      for y in range(max_row):
        cols[x].append(text[y][x])
        rows[y].append(text[y][x])
        fdiag[x+y].append(text[y][x])
        bdiag[x-y-min_bdiag].append(text[y][x])

    count = 0
    for line in rows:
      count += count_in_line("".join(line))
    for line in cols:
      count += count_in_line("".join(line))
    for line in fdiag:
      count += count_in_line("".join(line))
    for line in bdiag:
      count += count_in_line("".join(line))
    print(count)


def count_in_line(line: str) -> int:
  match = regex.findall(pattern, line, overlapped=True)
  return len(match)


def part2(filename: str) -> None:
  count = 0
  with open(filename, "r") as file:
    text = [line.strip() for line in file.readlines()]
    max_col = len(text[0])
    max_row = len(text)
    for r in range(max_row):
      for c in range(max_col):
        if r == 0 or r == max_row - 1 or c == 0 or c == max_col -1:
          continue
        if text[r][c] == 'A':
          diag = text[r-1][c-1] + text[r+1][c+1]
          if pattern2.search(diag):
            diag = text[r-1][c+1] + text[r+1][c-1]
            if pattern2.search(diag):
              count += 1
    print(count)


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
  args = parser.parse_args()

  part = args.part if args.part else env_part if env_part else "part1"

  if part == "part1":
    part1(args.input)
  else:
    part2(args.input)


if __name__ == "__main__":
  main()
