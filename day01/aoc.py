#!/usr/bin/env python
import argparse
import os
import re


def part1(filename: str) -> None:
    sum = 0
    left = []
    right = []
    with open(filename, "r") as file:
        for line in file:
            l, r = map(int, line.split())
            left.append(l)
            right.append(r)
        left.sort()
        right.sort()
        i = 0
        while i < len(left):
            if left[i] > right[i]:
                sum += left[i] - right[i]
            else:
                sum += right[i] - left[i]
            i += 1
    print(sum)


def part2(filename: str) -> None:
    sum = 0
    left = {}
    right = {}
    with open(filename, "r") as file:
        for line in file:
            l, r = map(int, line.split())
            if l in left:
                left[l] += 1
            else:
                left[l] = 1
            if r in right:
                right[r] += 1
            else:
                right[r] = 1
        for l, count in left.items():
            if l in right:
                sum += l * count * right[l]
    print(sum)


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
