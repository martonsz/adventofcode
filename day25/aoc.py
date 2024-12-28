#!/usr/bin/env python
import argparse
import os
import itertools

def part1(filename: str) -> None:
  locks, keys = get_locks_and_keys(filename)
  count = 0
  for lock in locks:
    for key in keys:
      fit = True
      for x in zip(lock, key):
        if x[0] + x[1] > 7:
          fit = False
          break
      if fit:
        count += 1
  print(count)

def get_locks_and_keys(filename: str) -> tuple[list, list]:
  locks = []
  keys = []
  with open(filename, "r") as file:
    count = list(itertools.repeat(0, 5))
    is_lock = False
    is_key = False
    for line in file:
      line = line.strip()
      if not line:
        if is_lock:
          locks.append(count)
        else:
          keys.append(count)
        #print(f"is_lock: {is_lock} {count}")
        count = list(itertools.repeat(0, 5))
        is_lock, is_key = False, False
        continue
      for i, col in enumerate(line):
        if col == "#":
          count[i] += 1
      if not is_lock and not is_key:
        if 1 in count:
          is_lock = True
        else:
          is_key = True
    if is_lock:
      locks.append(count)
    else:
      keys.append(count)
    return locks, keys


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
