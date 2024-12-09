#!/usr/bin/env python
import argparse
import os

grid = []
visited = {}
startRow = 0
startCol = 0
row_size = 0
col_size = 0

def part1(filename: str) -> None:
  solvePart1(filename)
  print(len(visited))


def solvePart1(filename: str) -> None:
  global grid
  global visited
  global startRow
  global startCol
  global row_size
  global col_size
  with open(filename, "r") as file:
    one_line = file.read()
    line_length = one_line.find("\n") + 1
    start = one_line.find("^")
    startRow = row = int(start / line_length)
    startCol = col = start  % line_length
    grid = [line.strip() for line in one_line.split("\n") if line.strip()]
    row_size = len(grid)
    col_size = len(grid[0])
    step_row = -1
    step_col = 0

    while row >= 0 and row < row_size and col >= 0 and col < col_size:
      #print(row, col,grid[row][col], len(visited))
      if grid[row][col] == '#':
        # Step back
        row += step_row * -1
        col += step_col * -1
        # Rotate
        if step_row != 0:
          step_col = 1 if step_row == -1 else -1
          step_row = 0
        else:
          step_row = 1 if step_col == 1 else -1
          step_col = 0
      else:
        visited[f"{str(row)},{str(col)}"] = True
        row += step_row
        col += step_col


def part2(filename: str) -> None:
  global grid
  global visited
  global startRow
  global startCol
  global row_size
  global col_size
  solvePart1(filename)
  step_row = -1
  step_col = 0
  loopCount = 0
  loopDict = {}
  row = startRow
  col = startCol

  for v in visited.keys():
    v_split = v.split(",")
    r = int(v_split[0])
    c = int(v_split[1])
  #for r in range(row_size):
  #  for c in range(col_size):
    if grid[r][c] != '.':
      continue
    grid[r] = f"{grid[r][:c]}#{grid[r][c+1:]}"
    while row >= 0 and row < row_size and col >= 0 and col < col_size:
      if grid[row][col] == '#':
        if f"r{row}c{col}sr{step_row}sc{step_col}" in loopDict:
          loopCount += 1
          break
        loopDict[f"r{row}c{col}sr{step_row}sc{step_col}"] = True
        # Step back
        row += step_row * -1
        col += step_col * -1
        # Rotate
        if step_row != 0:
          step_col = 1 if step_row == -1 else -1
          step_row = 0
        else:
          step_row = 1 if step_col == 1 else -1
          step_col = 0
      else:
        row += step_row
        col += step_col
    loopDict = {}
    grid[r] = f"{grid[r][:c]}.{grid[r][c+1:]}"
    row = startRow
    col = startCol
    step_row = -1
    step_col = 0
  print(loopCount)


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
