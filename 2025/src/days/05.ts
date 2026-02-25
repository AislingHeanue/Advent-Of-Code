import { Chunk, Console, Effect, Option, Stream } from "effect";
import type { Input } from "../input";

type IndexedRanges = {
  input: [start: number, end: number][];
  startMap: Map<number, number[]>;
  startsEnds: number[];
  startIndex: [number, number[]][];
};

export function* solve(input: Input, part: Part) {
  const [ranges, ids] = input.streamGroups;
  let indexedRanges = {
    input: new Array(yield* Stream.runCount(ranges!)),
    startMap: new Map<number, number[]>(),
    startsEnds: new Array((yield* Stream.runCount(ranges!)) * 6),
    startIndex: []
  } as IndexedRanges;

  yield* ranges!.pipe(
    Stream.map(getEnds),
    Stream.zipWithIndex,
    Stream.map(a => getIndexes(a, indexedRanges, part)),
    Stream.runDrain
  );

  indexedRanges.startIndex = indexedRanges.startMap
    .entries()
    .toArray()
    .sort((a, b) => a[0] - b[0]);

  if (part === 1) {
    return yield* ids!.pipe(
      Stream.map(getNumber),
      Stream.map(a => +(findInRange(a, indexedRanges) > 0)),
      Stream.runSum
    );
  } else {
    indexedRanges.startsEnds = new Set(indexedRanges.startsEnds)
      .values()
      .toArray()
      .sort((a, b) => a - b);
    const res = yield* Stream.fromIterable(indexedRanges.startsEnds).pipe(
      Stream.map(
        a => [a, findInRange(a, indexedRanges) > 0] as [number, boolean]
      ),
      Stream.runFold(
        { inRange: false, total: 0, previous: 0 },
        ({ inRange, total, previous }, value) => {
          // console.log(inRange, previous, total, value);
          if ((value[1] && inRange) || (!value[1] && !inRange)) {
            return { inRange, total, previous };
          }
          if (value[1]) {
            // entering a new range, modify previous
            return {
              inRange: value[1],
              total: total,
              previous: value[0]
            };
          } else {
            // in a range and coming to the end of it, modify total.
            return {
              inRange: value[1],
              total: total + value[0] - previous,
              previous
            };
          }
        }
      )
    );
    return res.total;
  }
}

const getNumber = (line: string) => {
  return +line;
};

const pattern = /([0-9]+)\-([0-9]+)/;
const getEnds = (line: string): [start: number, end: number] => {
  const matches = line.match(pattern);
  return [+matches![1]!, +matches![2]!];
};

const getIndexes = (
  [[start, end], i]: [[start: number, end: number], i: number],
  { input, startMap, startsEnds }: IndexedRanges,
  part: Part
) => {
  input[i] = [start, end];
  startMap.set(start, [...(startMap.get(start) ?? []), i]);
  if (part === 2) {
    startsEnds[6 * i] = start - 1;
    startsEnds[6 * i + 1] = start;
    startsEnds[6 * i + 2] = start + 1;
    startsEnds[6 * i + 3] = end - 1;
    startsEnds[6 * i + 4] = end;
    startsEnds[6 * i + 5] = end + 1;
  }
};

const findInRange = (id: number, indexedRanges: IndexedRanges): number => {
  let count = 0;
  let i = 0;
  while (
    i < indexedRanges.startIndex.length &&
    indexedRanges.startIndex[i]![0] <= id
  ) {
    for (const j of indexedRanges.startIndex[i]![1]) {
      if (indexedRanges.input[j]![1] >= id) {
        count++;
      }
    }
    i++;
  }
  return count;
};

export const solution: DaySolution = {
  solve,
  examples: {
    part1: {
      input: `3-5
10-14
16-20
12-18

1
5
8
11
17
32`,
      expected: 3
    },
    part2: {
      input: `3-5
10-14
16-20
12-18

1
5
8
11
17
32`,
      expected: 14
    }
  }
};
