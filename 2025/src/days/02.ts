import { Console, Effect, Option, Stream } from "effect";
import type { Input } from "../input";

export function* solve(input: Input, part: Part) {
  return yield* input.streamSplit(/,/).pipe(
    Stream.filterMap(extractStartEnd),
    Stream.mapEffect(input => sumInvalidIds(input, part), {
      concurrency: "unbounded"
    }),
    Stream.runSum
  );
}

const pattern = /([0-9]+)-([0-9]+)/;
const extractStartEnd = (line: string) => {
  const match = line.match(pattern);
  return match
    ? Option.some({ start: +match[1]!, end: +match[2]! })
    : Option.none();
};

const getNextInvalidNumber = (n: number, part: Part) =>
  Effect.gen(function* () {
    const digits = Math.floor(Math.log10(n)) + 1;
    let nextN;
    if (part === 1) {
      if (digits % 2 !== 0) {
        const divisor = Math.pow(10, (digits - 1) / 2);
        nextN = 10 * divisor * divisor + divisor; // next is 11, 1010 or 100100
      } else {
        nextN = yield* getNextInvalidNumberWithCopies(n, 2);
      }
    } else {
      return yield* getFactors(digits).pipe(
        Stream.mapEffect(copies => getNextInvalidNumberWithCopies(n, copies), {
          // concurrency: "unbounded"
        }),
        Stream.orElseIfEmpty(() => 11), // single digit numbers exist and always map to 11 as the next value
        Stream.runFold(Infinity, (acc, it) => (acc < it ? acc : it))
      );
    }

    return nextN;
  });

const getNextInvalidNumberWithCopies = (n: number, copies: number) =>
  Effect.gen(function* () {
    const digits = Math.floor(Math.log10(n)) + 1;
    let nextN;
    const divisor = Math.pow(10, digits / copies);
    let firstChunk = Math.floor(n / Math.pow(divisor, copies - 1));
    nextN = yield* addCopies(firstChunk, copies);
    while (n > nextN) {
      nextN = yield* addCopies(firstChunk++, copies);
    }
    return nextN;
  });

const addCopies = (chunk: number, copies: number) => {
  const digits = Math.floor(Math.log10(chunk)) + 1;
  return Stream.range(1, copies).pipe(
    Stream.map(n => chunk * Math.pow(10, digits * (n - 1))),
    Stream.runSum
  );
};

const getFactors = (n: number) =>
  Stream.range(2, n).pipe(Stream.filter(a => n % a === 0));

const sumInvalidIds = (input: { start: number; end: number }, part: Part) =>
  Effect.gen(function* () {
    return yield* Effect.iterate(
      { n: yield* getNextInvalidNumber(input.start, part), sum: 0 },
      {
        while: ({ n }) => n <= input.end,
        body: ({ n, sum }: { n: number; sum: number }) =>
          Effect.gen(function* () {
            // console.log("starting", input.start, "ending", input.end, "n", n);

            return {
              n: yield* getNextInvalidNumber(n + 1, part),
              sum: sum + n
            };
          })
      }
    ).pipe(Effect.map(it => it.sum));
  });

export const solution: DaySolution = {
  solve,
  examples: {
    part1: {
      input: `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`,
      expected: 1227775554
    },
    part2: {
      input: `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`,
      expected: 4174379265
    }
  }
};
