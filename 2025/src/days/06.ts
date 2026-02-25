import { Chunk, Console, Effect, Stream } from "effect";
import type { Input } from "../input";

export function* solve(input: Input, part: Part) {
  if (part === 1) {
    return transpose(input.lines.map(line => line.split(/\s/).filter(Boolean)))
      .map(a =>
        doMaths(
          a.slice(undefined, -1).map(Number),
          a[a.length - 1]![a[a.length - 1]!.length - 1]!
        )
      )
      .reduce((acc, item) => acc + item);
  } else {
    return yield* Stream.fromIterable(transpose(input.grid).reverse()).pipe(
      Stream.map(a => {
        return [+a.slice(undefined, -1).join("").trim(), a[a.length - 1]] as [
          number,
          string
        ];
      }),
      Stream.split(a => !a[0]),
      Stream.map(Chunk.toArray),
      Stream.map(a =>
        doMaths(a.map(a => a[0]).map(Number), a[a.length - 1]![1]!)
      ),
      Stream.runSum
    );
  }
}

const transpose = (a: string[][]) => {
  const length = Math.max(...a.map(internal => internal.length));
  let out = new Array<string[]>(length);
  // quickest way i know how to transpose right now
  for (let i = 0; i < length; i++) {
    out[i] = a.map(internal => internal[i] || " ");
  }
  return out;
};

const doMaths = (input: number[], symbol: string) => {
  switch (symbol) {
    case "+":
      return input.reduce((acc, item) => acc + item, 0);
    case "*":
      return input.reduce((acc, item) => acc * item, 1);
    default:
      console.log("ERROR UNEXPECTED INPUT", input);
      return -1;
  }
};

export const solution: DaySolution = {
  solve,
  examples: {
    part1: {
      input: `123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +`,
      expected: 4277556
    },
    part2: {
      input: `123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +`,
      expected: 3263827
    }
  }
};
