#!/usr/bin/env python
import argparse
import os

class Node:
    def __init__(self, name):
        self.name = name
        self.connections = []

    def add_connection(self, node):
        if node.name != self.name and node not in self.connections:
            self.connections.append(node)

    def has_connection(self, node):
        return any(conn.name == node.name for conn in self.connections)

    def __str__(self):
        connections_names = [conn.name for conn in self.connections]
        return f"{self.name} -> {', '.join(connections_names)}"


def part1(filename: str) -> None:
  nodes = {}
  with open(filename, "r") as file:
    for line in file:
      names = line.strip().split("-")
      node0 = nodes.get(names[0], None)
      node1 = nodes.get(names[1], None)
      if not node0:
         node0 = Node(names[0])
         nodes[names[0]] = node0
      if not node1:
         node1 = Node(names[1])
         nodes[names[1]] = node1

      node1.add_connection(node0)
      node0.add_connection(node1)

  count = {}
  for node in nodes.values():
    for connection in node.connections:
      for connection2 in node.connections:
        if connection.name == connection2.name:
           continue
        if connection2.has_connection(connection):
          if node.name[0] == "t" or connection.name[0] == "t" or connection2.name[0] == "t":
            count[frozenset([node.name, connection.name, connection2.name])] = 1
  print(len(count))

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
