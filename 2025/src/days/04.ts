import { Console, Effect, Option, Stream } from "effect";
import type { Input } from "../input";

const printing = false;

export function* solve(input: Input, part: Part) {
  const grid = input.grid.map(line => line.map(c => c === "@"));

  let pipeline = Stream.range(0, grid[0]!.length - 1).pipe(
    Stream.zipWithPreviousAndNext,
    Stream.cross(
      Stream.range(0, grid.length - 1).pipe(Stream.zipWithPreviousAndNext)
    ),
    // Stream.tap(a => (a[1][1] === 0 ? print("\n") : Effect.void)),
    // Stream.tap(a =>
    //   printing && Option.isSome(canAccess(a, grid))
    //     ? print("x")
    //     : print(grid[a[0][1]]![a[1][1]]! ? "@" : ".")
    // ),
    Stream.filterMap(a => canAccess(a, grid)),
    Stream.runCollect,
    Effect.map(chunk => {
      for (const a of chunk) {
        modifyGrid(a, grid);
      }

      return chunk.length;
    })
  );

  const res = Effect.iterate(
    { cont: true, returned: 0 },
    {
      while: ({ cont }) => cont,
      body: ({ returned }) =>
        pipeline.pipe(
          Effect.map(a => ({
            cont: a > 0 && part !== 1,
            returned: a + returned
          }))
        )
    }
  );

  const returned = (yield* res).returned;
  yield* print("\n");
  return returned;
}

const print = (message: string) =>
  Effect.sync(() => {
    if (printing) {
      process.stdout.write(message);
    }
  });

const canAccess = (
  input: [
    [Option.Option<number>, number, Option.Option<number>],
    [Option.Option<number>, number, Option.Option<number>]
  ],
  grid: boolean[][]
): Option.Option<[number, number]> => {
  if (!grid[input[0][1]]![input[1][1]]!) {
    return Option.none();
  }
  let count = 0;
  for (let xi = 0; xi <= 2; xi++) {
    for (let yi = 0; yi <= 2; yi++) {
      let y = input[0][yi];
      let x = input[1][xi];
      if (Option.isOption(y)) {
        y = Option.getOrUndefined(y);
      }
      if (Option.isOption(x)) {
        x = Option.getOrUndefined(x);
      }
      if (x !== undefined && y !== undefined) {
        count += +grid[y]![x]!;
      }
    }
  }
  return count <= 4 ? Option.some([input[0][1], input[1][1]]) : Option.none();
};

const modifyGrid = (input: [number, number], grid: boolean[][]) => {
  grid[input[0]]![input[1]] = false;
};

export const solution: DaySolution = {
  solve,
  examples: {
    part1: {
      input: `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`,
      expected: 13
    },
    part2: {
      input: `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`,
      expected: 43
    }
  }
};
