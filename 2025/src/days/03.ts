import { Console, Effect, Stream } from "effect";
import type { Input } from "../input";

export function* solve(input: Input, part: Part) {
  const nums = part === 1 ? 2 : 12;
  return yield* input.stream.pipe(
    Stream.map(toNumberArray),
    Stream.mapEffect(a => getHighest(a, nums)),
    // Stream.tap(Console.log),
    Stream.runSum
  );
}

const toNumberArray = (line: string) => line.split("").map(a => +a);

const getHighest = (line: number[], nums: number) =>
  Effect.gen(function* () {
    const out = new Array(nums);
    let index = 0;
    for (let i = 1; i <= nums; i++) {
      out[i - 1] = Math.max(...line.slice(index, -nums + i || undefined));
      index += line.slice(index).indexOf(out[i - 1]) + 1;
    }
    return +out.join("");
  });

export const solution: DaySolution = {
  solve,
  examples: {
    part1: {
      input: `987654321111111
811111111111119
234234234234278
818181911112111`,
      expected: 357
    },
    part2: {
      input: `987654321111111
811111111111119
234234234234278
818181911112111`,
      expected: 3121910778619
    }
  }
};
