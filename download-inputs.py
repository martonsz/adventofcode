#!/usr/bin/env python

import argparse
import os
import re
import aocd
from pathlib import Path
from datetime import datetime

def clear_inputs():
  # Find file with regex
  for root, dirs, files in os.walk('.'):
    for file in files:
      if re.search(r'input\d*.txt', file):
        print(f'Removing {root}/{file}')
        os.remove(os.path.join(root, file))
      if re.search(r'example\d*.txt', file):
        print(f'Removing {root}/{file}')
        os.remove(os.path.join(root, file))
  print("All input files removed.")

def setup_token(clear_token=False):
  aocd_dir = os.getenv('AOCD_DIR', os.path.join(Path.home(), '.config', 'aocd'))
  if not os.path.exists(aocd_dir):
    os.makedirs(aocd_dir)

  if clear_token:
    os.remove(f'{aocd_dir}/token')

  if not os.path.exists(f'{aocd_dir}/token'):
    print('Please enter your Advent of Code session token.')
    token = input('Token: ')
    with open(f'{aocd_dir}/token', 'w') as f:
        f.write(token.strip())

  try:
    aocd.get_data()
  except Exception as e :
    print(f'That token does not work. Please try again. Error msg: "{e}"')
    setup_token(True)
    return

def download_inputs():
  year = datetime.year

  for day in range(1, 26):
    if day < 10:
      day = f'0{day}'
    os.makedirs(f'day{day}', exist_ok=True)
    if not os.path.exists(f'day{day}/input.txt'):
      print(f'Downloading input for day {day}')
      data = aocd.get_data(day=int(day))
      with open(f'day{day}/input.txt', 'w') as f:
        f.write(data)
    if not os.path.exists(f'day{day}/example.txt'):
      print(f'Downloading example for day {day}')
      data = aocd.get_puzzle(day=int(day))
      if len(data.examples) == 1:
        with open(f'day{day}/example.txt', 'w') as f:
          f.write(data.examples[0].input_data)
      else:
        for i, example in enumerate(data.examples):
          with open(f'day{day}/example{i+1}.txt', 'w') as f:
            f.write(example.input_data)


if __name__ == "__main__":
  parser = argparse.ArgumentParser(description="Download input files")
  parser.add_argument(
    "-c",
    "--clear",
    action="store_true",
    default=False,
    help="Clear all input files",
  )
  args = parser.parse_args()

  setup_token()
  if args.clear:
    clear_inputs()
    download_inputs()
