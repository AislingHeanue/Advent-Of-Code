import { Chunk, Effect, Stream } from "effect";
import type { Input } from "../input";
import { memoize } from "micro-memoize";
import { maxSafeInteger } from "effect/FastCheck";

let map: Map<string, string[]>;

export function* solve(input: Input, part: Part) {
  map = yield* input.stream.pipe(
    Stream.map(parse),
    Stream.runCollect,
    Effect.map(a => new Map(Chunk.toArray(a)))
  );
  const start = part === 1 ? "you" : "svr";
  return traverseCached(start, { seenFft: false, seenDac: false }, part);
}

const parse = (line: string) => {
  const [name, outString] = line.split(": ");
  const outs = outString!.split(" ");
  return [name, outs] as [string, string[]];
};

const traverse = (
  name: string,
  { seenDac, seenFft }: { seenDac: boolean; seenFft: boolean },
  part: Part
): number => {
  if (name === "out") {
    return +((seenDac && seenFft) || part === 1);
  }
  if (name === "fft") {
    seenFft = true;
  }
  if (name === "dac") {
    seenDac = true;
  }
  return map
    .get(name)!
    .reduce(
      (acc, item) => acc + traverseCached(item, { seenDac, seenFft }, part),
      0
    );
};

const traverseCached = memoize(traverse, {
  maxSize: Number.MAX_SAFE_INTEGER,
  isKeyItemEqual: "shallow"
});

export const solution: DaySolution = {
  solve,
  examples: {
    part1: {
      input: `aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out`,
      expected: 5
    },
    part2: {
      input: `svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`,
      expected: 2
    }
  }
};
