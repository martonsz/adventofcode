#!/usr/bin/env python
import argparse
import os


def part1(filename: str) -> None:
  with open(filename, "r") as file:
    count = 0
    for line in file:
      numbers = list(map(int, line.strip().split()))
      prev = numbers[0]
      increase = True
      decrease = True
      for n in numbers[1:]:
        if increase:
          increase = (prev < n and n - prev <= 3)
        if decrease:
          decrease = (prev > n and prev - n <= 3)
        if not (increase or decrease):
          break
        prev = n
      if increase or decrease:
        count += 1
    print(count)


def part2(filename: str) -> None:
  with open(filename, "r") as file:
    count = 0
    for line in file:
      numbers = list(map(int, line.strip().split()))
      if part2IsIncrease(numbers):
        count += 1
      elif part2IsIncrease(numbers[::-1]):
        count += 1
      else:
        for skip_index in range(len(numbers)):
          newNumbers = [value for i, value in enumerate(numbers) if i != skip_index]
          if part2IsIncrease(newNumbers):
            count += 1
            break
          elif part2IsIncrease(newNumbers[::-1]):
            count += 1
            break
    print(count)

def part2IsIncrease(numbers: list[int]) -> bool:
      prev = numbers[0]
      increase = True
      decrease = True
      for n in numbers[1:]:
        if increase:
          increase = (prev < n and n - prev <= 3)
        if decrease:
          decrease = (prev > n and prev - n <= 3)
        if not (increase or decrease):
          break
        prev = n
      return increase or decrease


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
