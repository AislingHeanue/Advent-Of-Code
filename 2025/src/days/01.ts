import { Stream } from "effect";
import type { Input } from "../input";

const pattern = /\b([LR])([1-9][0-9]*)/;

const R = (n: number, amount: number) => (n + amount) % 100;
const L = (n: number, amount: number) => (n - amount) % 100;

export function* solve(input: Input, part: Part) {
  // left = down
  // right = up
  // 0 - 1 = 99 (use mod 100)
  let n = 50;
  const result = yield* input.stream.pipe(
    Stream.runFold({ n, count: 0 }, (input, line) => {
      const match = line.match(pattern)!;
      let amount = +match[2]!;
      const n = input.n;
      if (amount === 0) {
        return input;
      }
      let zeros = Math.floor(amount / 100);
      amount -= zeros * 100;
      switch (match[1] as "L" | "R") {
        case "L":
          input.n = n - amount;
          if (input.n < 0) {
            input.n += 100;
            if (n !== 0) {
              zeros += 1;
            }
          }
          break;
        case "R":
          input.n = n + amount;
          if (input.n >= 100) {
            input.n -= 100;
            if (input.n !== 0) {
              zeros += 1;
            }
          }
          break;
      }
      if (input.n === 0) {
        input.count += 1;
      }
      if (part === 2) {
        input.count += zeros;
      }

      return input;
    })
  );

  return result.count;
}

export const solution: DaySolution = {
  solve,
  examples: {
    part1: {
      input: `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`,
      expected: 3
    },
    part2: {
      input: `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`,
      expected: 6
    }
    // part2: {
    //   input: `example input here`,
    //   expected: 456
    // }
  }
};
