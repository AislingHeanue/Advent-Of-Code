import { Chunk, Effect, identity, Option, Order, Stream } from "effect";
import type { Input } from "../input";

type Point2D = { x: number; y: number };

export function* solve(input: Input, part: Part) {
  const firstPosition = getCoords(input.lines[0]!);
  const coords = input.stream.pipe(Stream.map(getCoords));
  const safeRanges =
    part === 1
      ? undefined
      : yield* setupQuestion2Garbage(coords, firstPosition);

  const biggest = yield* Stream.cross(
    Stream.fromIterable(input.lines.map(getCoords)),
    Stream.fromIterable(input.lines.map(getCoords))
  ).pipe(
    Stream.filter(
      a => a[0].x < a[1].x || (a[0].x === a[1].x && a[0].y < a[1].y)
    ),
    Stream.map(a => [a, getArea(a)] as [[Point2D, Point2D], number]),
    Stream.runCollect,
    Stream.map(Chunk.sortWith(a => a[1], Order.reverse(Order.number))),
    Stream.flattenChunks,
    part === 1
      ? Stream.take(1)
      : Stream.find(a => checkAllShaded(a[0], safeRanges!)),
    Stream.runHead,
    Effect.map(Option.getOrThrow)
  );
  return biggest[1]!;
}

const pattern = /([0-9]+),([0-9]+)/;
const getCoords = (line: string): Point2D => {
  const matches = line.match(pattern)!;
  return { x: +matches[1]!, y: +matches[2]! };
};

const getArea = ([left, right]: [Point2D, Point2D]) => {
  return (Math.abs(left.x - right.x) + 1) * (Math.abs(left.y - right.y) + 1);
};

const getHorizontalPoints = ([point, next]: [Point2D, Point2D]): Point2D[] => {
  if (point.x !== next.x) {
    // skip the leftmost point
    return Array(Math.abs(next.x - point.x))
      .keys()
      .map(a => ({
        x: a + Math.min(point.x, next.x),
        y: point.y
      }))
      .toArray();
  }
  return [];
};

const getContourPoints = ([point, next]: [Point2D, Point2D]): Point2D[] => {
  if (point.x === next.x) {
    return Array(Math.abs(next.y - point.y) + 1)
      .keys()
      .map(a => ({ x: point.x, y: a + Math.min(point.y, next.y) }))
      .toArray();
  } else {
    return Array(Math.abs(next.x - point.x) + 1)
      .keys()
      .map(a => ({ x: a + Math.min(point.x, next.x), y: point.y }))
      .toArray();
  }
};

const pointToString = (p: Point2D) => {
  return `${p.x}%${p.y}`;
};

const setupQuestion2Garbage = (
  coords: Stream.Stream<Point2D>,
  firstPosition: Point2D
) =>
  Effect.gen(function* () {
    const coordsWithNext = coords.pipe(
      Stream.zipWithNext,
      Stream.map(
        a =>
          [a[0], Option.getOrElse(a[1], () => firstPosition)] as [
            Point2D,
            Point2D
          ]
      )
    );
    const leftFacingContourTiles = yield* coordsWithNext.pipe(
      Stream.map(getHorizontalPoints),
      Stream.flattenIterables,
      Stream.runCollect,
      Effect.map(Chunk.toArray),
      Effect.map(a => new Set(a.map(pointToString)))
    );
    const contourTiles = yield* coordsWithNext.pipe(
      Stream.map(getContourPoints),
      Stream.flattenIterables,
      Stream.runCollect,
      Effect.map(Chunk.toArray),
      Effect.map(a => new Set(a.map(pointToString)))
    );
    return yield* coords.pipe(
      Stream.runFold(
        { xs: new Set<number>(), ys: new Set<number>() },
        ({ xs, ys }, { x, y }) => {
          xs.add(x);
          ys.add(y);
          return { xs, ys };
        }
      ),
      Stream.map(({ xs, ys }) => {
        const xArray = xs
          .keys()
          .toArray()
          .toSorted((a, b) => a - b);
        const yArray = ys
          .keys()
          .toArray()
          .toSorted((a, b) => a - b);
        return { xs: xArray, ys: yArray };
      }),
      Stream.runHead,
      Effect.map(Option.getOrThrow),
      Effect.map(
        ranges =>
          new Map(
            ranges!.xs.map(x => {
              let inLoop = false;
              let currentRangeStart = undefined as number | undefined;
              let safeForThisX = [] as { min: number; max: number }[];
              ranges!.ys.forEach((y, j) => {
                // for each x march through the list of ys.
                // If on the contour, always safe.
                // If not on the contour, safe if inLoop
                // inLoop is flipped whenever leftFacingContourTiles is encountered
                if (leftFacingContourTiles!.has(pointToString({ x, y }))) {
                  inLoop = !inLoop;
                }
                if (
                  currentRangeStart === undefined &&
                  contourTiles!.has(pointToString({ x, y }))
                ) {
                  currentRangeStart = y;
                } else if (
                  currentRangeStart !== undefined &&
                  !contourTiles!.has(pointToString({ x, y })) &&
                  !inLoop
                ) {
                  safeForThisX.push({
                    min: currentRangeStart,
                    max: ranges!.ys[j - 1]!
                  });
                  currentRangeStart = undefined;
                } else if (
                  currentRangeStart !== undefined &&
                  y === ranges!.ys[ranges!.ys.length - 1]
                ) {
                  safeForThisX.push({ min: currentRangeStart, max: y });
                  currentRangeStart = undefined;
                }
              });
              return [x, safeForThisX];
            })
          )
      )
    );
  });

const checkAllShaded = (
  corners: [Point2D, Point2D],
  safeRanges: Map<number, { min: number; max: number }[]>
) => {
  let left = Math.min(corners[0].x, corners[1].x);
  let right = Math.max(corners[0].x, corners[1].x);
  let up = Math.min(corners[0].y, corners[1].y);
  let down = Math.max(corners[0].y, corners[1].y);
  for (const [x, xRanges] of safeRanges) {
    if (x < left || x > right) {
      continue;
    }
    if (!xRanges.some(a => a.min <= up && a.max >= down)) {
      return false;
    }
  }
  return true;
};

export const solution: DaySolution = {
  solve,
  examples: {
    part1: {
      input: `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`,
      expected: 50
    },
    part2: {
      input: `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`,
      expected: 24
    }
  }
};
