import { Console, Effect, Option, Stream } from "effect";
import type { Input } from "../input";

const printing = false;

export function* solve(input: Input, part: Part) {
  return yield* input.stream.pipe(
    Stream.runFold(
      {
        total: 0,
        previousLine: Option.none() as Option.Option<number[]>,
        timelines: 1
      },
      (acc, item) => {
        const propagation = propagate([acc.previousLine, item], part, printing);
        return {
          total: acc.total + propagation[1],
          previousLine: Option.some(propagation[0]),
          timelines: propagation[2]
        };
      }
    ),
    Effect.map(a => (part === 1 ? a.total : a.timelines))
  );
}

const propagate = (
  [previousLineOption, line]: [Option.Option<number[]>, string],
  part: number,
  printing: boolean
): [number[], number, number] => {
  if (Option.isNone(previousLineOption)) {
    if (printing) {
      console.log(line);
    }
    return [line.split("").map(input => (input === "S" ? 1 : 0)), 0, 1];
  }
  const previousLine = Option.getOrThrow(previousLineOption);
  const lineArr = line.split("").map(input => (input === "^" ? -1 : 0));
  let splits = 0;
  let timelines = 0;
  for (let i = 0; i < line.length; i++) {
    if (previousLine[i]! > 0) {
      timelines += previousLine[i]!;
      if (line[i] === "^") {
        splits++;
        timelines += previousLine[i]!;
        lineArr[i - 1]! += previousLine[i]!;
        lineArr[i + 1]! += previousLine[i]!;
      } else {
        lineArr[i]! += previousLine[i]!;
      }
    }
  }
  if (printing) {
    console.log(
      lineArr
        .map(num => {
          if (part === 1) {
            return num === -1 ? "^" : num ? "|" : ".";
          } else {
            return num === -1 ? "^" : num ? num % 10 : ".";
          }
        })
        .join(""),
      splits,
      timelines
    );
  }
  return [lineArr, splits, timelines];
};

export const solution: DaySolution = {
  solve,
  examples: {
    part1: {
      input: `.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`,
      expected: 21
    },
    part2: {
      input: `.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`,
      expected: 40
    }
  }
};
