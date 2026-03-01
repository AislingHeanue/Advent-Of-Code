import { Chunk, Effect, Stream } from "effect";
import type { Input } from "../input";
let map: Map<string, string[]>;
let traverseCached: typeof traverse;

export function* solve(input: Input, part: Part) {
  map = yield* input.stream.pipe(
    Stream.map(parse),
    Stream.runCollect,
    Effect.map(a => new Map(Chunk.toArray(a)))
  );
  const start = part === 1 ? "you" : "svr";
  traverseCached = yield* Effect.cachedFunction(traverse);
  return yield* traverseCached(
    JSON.stringify({
      name: start,
      seenFft: false,
      seenDac: false,
      part
    })
  );
}

const parse = (line: string) => {
  const [name, outString] = line.split(": ");
  const outs = outString!.split(" ");
  return [name, outs] as [string, string[]];
};

const traverse = (data: string): Effect.Effect<number, any, any> =>
  Effect.gen(function* () {
    let { name, seenDac, seenFft, part } = JSON.parse(data);
    if (name === "out") {
      return +((seenDac && seenFft) || part === 1);
    }
    if (name === "fft") {
      seenFft = true;
    }
    if (name === "dac") {
      seenDac = true;
    }
    let total: number = 0;
    for (const item of map.get(name)!) {
      total += yield* traverseCached!(
        JSON.stringify({
          name: item,
          seenDac,
          seenFft,
          part
        })
      );
    }
    return total;
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
