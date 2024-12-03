#!/usr/bin/env python
import argparse
import os
import re


def part1(filename: str) -> None:
  with open(filename, "r") as file:
    all = file.read().replace('\n', '')
    print(getSum(all))

def part2(filename: str) -> None:
  with open(filename, "r") as file:
    all = file.read().replace('\n', '')
    sum = 0
    pattern = r"don't\(\).*?(?=$|do\(\))"
    newLine =  re.sub(pattern, "", all)
    sum += getSum(newLine)
    print(sum)


def getSum(line: str) -> int:
  sum = 0
  pattern = r"mul\((\d+),(\d+)\)"
  for match in re.finditer(pattern, line.strip()):
    num1 = int(match.group(1))
    num2 = int(match.group(2))
    sum += num1 * num2
  return sum


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
