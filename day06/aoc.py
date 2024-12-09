#!/usr/bin/env python
import argparse
import os
import re


def part1(filename: str) -> None:
  with open(filename, "r") as file:
    one_line = file.read()
    line_length = one_line.find("\n") + 1
    start = one_line.find("^")
    row = int(start / line_length)
    col = start  % line_length
    grid = [line.strip() for line in one_line.split("\n") if line.strip()]
    row_size = len(grid)
    col_size = len(grid[0])
    step_row = -1
    step_col = 0
    visited = {}
    while row >= 0 and row < row_size and col >= 0 and col < col_size:
      print(row, col, grid[row][col])
      visited[str(row) + str(col)] = True
      try:
        if step_row != 0:
          # Up or down
          if grid[row + step_row][col] == '#':
            step_col = 1 if step_row == -1 else -1
            step_row = 0
        else:
          # Left or right
          if grid[row][col + step_col] == '#':
            step_row = 1 if step_col == 1 else -1
            step_col = 0
        row += step_row
        col += step_col
      except IndexError:
        break

    print("")
    print(len(visited))


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
  args = parser.parse_args()

  part = args.part if args.part else env_part if env_part else "part1"

  if part == "part1":
    part1(args.input)
  else:
    part2(args.input)


if __name__ == "__main__":
  main()
