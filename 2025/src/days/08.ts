import { Chunk, Console, Effect, Option, Order, Stream } from "effect";
import type { Input } from "../input";

export function* solve(input: Input, part: Part) {
  // test case vs real input parameters
  const connectionsToMake =
    part === 1
      ? input.lines.length > 20
        ? 1000
        : 10
      : Number.MAX_SAFE_INTEGER;
  return yield* Stream.cross(
    input.stream.pipe(Stream.map(getCoords), Stream.zipWithIndex),
    input.stream.pipe(Stream.map(getCoords), Stream.zipWithIndex)
  ).pipe(
    // get pairwise distances
    Stream.filter(a => a[0][1] < a[1][1]),
    Stream.map(
      ([[pointA, leftNum], [pointB, rightNum]]) =>
        [
          [leftNum, rightNum],
          Math.pow(pointA[0] - pointB[0], 2) +
            Math.pow(pointA[1] - pointB[1], 2) +
            Math.pow(pointA[2] - pointB[2], 2)
        ] as [[number, number], number]
    ),
    // sort distances
    Stream.runCollect,
    Stream.map(a => a.pipe(Chunk.sortWith(a => a[1], Order.number))),
    Stream.flattenChunks,
    Stream.map(a => a[0]),
    Stream.take(connectionsToMake),
    // convert the list of connections to a set of connected graphs
    Stream.runFoldWhile(
      {
        blobs: input.lines.map((_, i) => i),
        sizes: input.lines.map(() => 1),
        previous: Option.none() as Option.Option<[number, number]>
      },
      acc => acc.blobs.length > 1,
      (acc, [left, right]) => {
        // "union-find" algorithm
        const find = (x: number): number =>
          acc.blobs[x] === x ? x : (acc.blobs[x] = find(acc.blobs[x]!));

        const aSet = find(left);
        const bSet = find(right);
        if (aSet != bSet) {
          acc.blobs[aSet] = bSet;
          acc.sizes[bSet]! += acc.sizes[aSet]!;
          acc.sizes[aSet]! = 0;
          acc.previous = Option.some([left, right]);
        }
        return acc;
      }
    ),
    // get the three biggest node sizes and multiply them
    // or find the last processed pair and multiply its x coords
    Effect.map(a =>
      part === 1
        ? a.sizes
            // .map(b => b)
            .sort((a, b) => b - a)
            // .reverse()
            .slice(0, 3)
            .reduce((acc, item) => acc * item, 1)
        : getCoords(input.lines[Option.getOrThrow(a.previous)[0]]!)[0]! *
          getCoords(input.lines[Option.getOrThrow(a.previous)[1]]!)[0]!
    )
  );
}

const pattern = /([0-9]+),([0-9]+),([0-9]+)/;
const getCoords = (input: string) => {
  const matches = input.match(pattern)!;
  return [+matches[1]!, +matches[2]!, +matches[3]!] as [number, number, number];
};

export const solution: DaySolution = {
  solve,
  examples: {
    part1: {
      input: `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`,
      expected: 40
    },
    part2: {
      input: `162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689`,
      expected: 25272
    }
  }
};
