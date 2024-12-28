#!/usr/bin/env python
import argparse
import os

ruleDict = {}

def part1(filename: str) -> None:
  print(solve(filename, True))


def solve(filename: str, part1: bool) -> None:
  global ruleDict, afterDict
  sum = 0
  with open(filename, "r") as file:
    for line in file:
      if line.find("|") != -1:
        before, after = list(map(str, line.strip().split("|")))
        ruleDict[f"{before}|{after}"] = True
      elif not line.strip():
        continue
      else:
        to_sort = [x for x in line.strip().split(",")]
        sorted = quick_sort(to_sort)
        if part1 and to_sort == sorted or not part1 and to_sort != sorted:
          sum += int(sorted[len(sorted)//2])
    return sum


def quick_sort(lst: list) -> None:
  if len(lst) <= 1:
    return lst

  pivot = lst[len(lst)//2]
  sml = []
  mid = []
  big = []
  for x in lst:
    if f"{x}|{pivot}" in ruleDict:
      sml.append(x)
    elif f"{pivot}|{x}" in ruleDict:
      big.append(x)
    else:
      mid.append(x)
  return quick_sort(sml) + mid + quick_sort(big)


def part2(filename: str) -> None:
  print(solve(filename, False))

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
