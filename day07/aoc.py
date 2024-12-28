#!/usr/bin/env python
import argparse
import os
import itertools

def part1(filename: str) -> None:
  answer, _ = part1Solve(filename)
  print(answer)

def part1Solve(filename: str) -> tuple[str, dict]:
  with open(filename, "r") as file:
    answerSum = 0
    rows = {}
    rowIndex = 0
    for line in file:
      answer, numbers = line.split (":")
      numbers = list(map(int, numbers.strip().split()))
      combinations = itertools.product("+*", repeat=len(numbers)-1)
      for combination in combinations:
        i = 0
        sum = 0
        for group in itertools.pairwise(numbers):
          if i == 0:
            sum = calculate(group[0], combination[i], group[1])
          else:
            sum = calculate(sum, combination[i], group[1])
          i += 1
        if sum == int(answer):
          rows[rowIndex] = True
          answerSum += int(sum)
          break
      rowIndex += 1
    return answerSum, rows


def calculate(left: int, symbol: str, right: int) -> int:
  if symbol == "+":
    return left + right
  elif symbol == "*":
    return left * right
  elif symbol == "|":
    return int(f"{left}{right}")


def part2(filename: str) -> None:
  part1Answer, skipRows = part1Solve(filename)
  with open(filename, "r") as file:
    answerSum = part1Answer
    rowIndex = 0
    for line in file:
      if rowIndex in skipRows:
        rowIndex += 1
        continue
      rowIndex += 1
      answer, numbers = line.split (":")
      numbers = list(map(int, numbers.strip().split()))
      combinations = itertools.product("+*|", repeat=len(numbers)-1)
      for combination in combinations:
        i = 0
        sum = 0
        for group in itertools.pairwise(numbers):
          if i == 0:
            sum = calculate(group[0], combination[i], group[1])
          else:
            sum = calculate(sum, combination[i], group[1])
          i += 1
        if sum == int(answer):
          answerSum += int(sum)
          break
    print(answerSum)


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
