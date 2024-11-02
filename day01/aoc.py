#!/usr/bin/env python
import argparse
import os
import re

mapping = { "0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, 
    "zero": 0, "one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9 }
mappingReverse = {
    "0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
    "orez": 0, "eno": 1, "owt": 2, "eerht": 3, "ruof": 4, "evif": 5, "xis": 6, "neves": 7, "thgie": 8, "enin": 9 }

def part1(filename: str) -> None:
    sum = 0
    with open(filename, 'r') as file:
        for line in file:
            digits = ''.join(re.findall(r'\d', line.strip()))
            if not digits:
                continue
            sum += int(f"{digits[0]}{digits[-1]}")
    print(sum)

def part2(filename: str) -> None:
    sum = 0
    with open(filename, 'r') as file:
        for line in file:
            sum += findWords(line.strip())
    print(sum)

def findWords(line: str) -> int:
    leftIndex = len(line)
    leftDigit = 0
    rightIndex = leftIndex
    rightDigit = 0
    for numberName, value in mapping.items():
        i = line.find(numberName)
        if i > -1 and i < leftIndex:
            leftDigit = value
            leftIndex = i
    for numberName, value in mappingReverse.items():
        i = line[::-1].find(numberName)
        if i > -1 and i < rightIndex:
            rightDigit = value
            rightIndex = i
    if leftIndex == len(line) - rightIndex:
        return leftDigit
    else:
        return int(f"{leftDigit}{rightDigit}")

def main():
    env_part = os.getenv('part')
    parser = argparse.ArgumentParser(description="Process input file and part flag.")
    parser.add_argument('-i', '--input', type=str, default='input.txt', help="Input file (default: input.txt)")
    group = parser.add_mutually_exclusive_group()
    group.add_argument('-1', action='store_const', dest='part', const='part1', help="Set part to 'part1'")
    group.add_argument('-2', action='store_const', dest='part', const='part2', help="Set part to 'part2'")
    args = parser.parse_args()
    
    part = args.part if args.part else env_part if env_part else "part1"
    
    if part == "part1":
        part1(args.input)
    else:
        part2(args.input)

if __name__ == "__main__":
    main()