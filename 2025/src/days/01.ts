import type { Input } from "../input";

export const solve = (input: Input, part: Part): number | undefined => {
  return undefined;
};

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
L821`,
      expected: 123
    }
    // part2: {
    //   input: `example input here`,
    //   expected: 456
    // }
  }
};
