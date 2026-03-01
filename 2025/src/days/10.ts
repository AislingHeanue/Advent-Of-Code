import { Chunk, Console, Effect, Order, Stream } from "effect";
import type { Input } from "../input";
import solver, {
  type ConstraintBound,
  type SolveResult,
  type VariableCoefficients
} from "javascript-lp-solver";

type Parsed = {
  lights: number[];
  buttons: ButtonSelection;
  maxJoltage: number;
};
type ButtonSelection = number[][];

type Output = {
  lights: boolean[];
};

export function* solve(input: Input, part: Part) {
  const stream = input.stream.pipe(Stream.map(a => parse(a, part)));
  if (part === 1) {
    return yield* stream.pipe(
      // naive approach: enumerate every possible button combination and check if they work in order (ye olde BFS)
      Stream.mapEffect(parsed =>
        Effect.gen(function* () {
          return [
            parsed.lights,
            yield* Stream.range(0, Math.pow(2, parsed.buttons.length) - 1).pipe(
              Stream.map(mask =>
                parsed.buttons.filter((_, i) => (mask & (1 << i)) !== 0)
              ),
              Stream.runCollect,
              Effect.map(Chunk.toArray)
            )
          ] as [number[], ButtonSelection[]];
        })
      ),
      Stream.map(([lights, buttonSelections]) => {
        buttonSelections.sort((a, b) => a.length - b.length);
        return [lights, buttonSelections] as [number[], ButtonSelection[]];
      }),
      Stream.map(([lights, buttonSelections]) => {
        return (
          buttonSelections.find(buttons => {
            let newLights = [...lights];
            for (const button of buttons) {
              for (const light of button) {
                newLights[light]!--;
              }
            }

            return newLights.every(l => l % 2 === 0);
          })?.length || 1000
        );
      }),
      Stream.runSum
    );
  } else {
    return yield* stream.pipe(
      // objectively better, very boring solution, just plug it into a LinAlg solver
      Stream.map(
        a =>
          solver.Solve({
            optimize: "cost",
            opType: "min",
            constraints: Array(a.lights.length)
              .keys()
              .reduce(
                (acc, i) => {
                  acc[`light-${i}`] = { equal: a.lights[i]! };
                  return acc;
                },
                {} as Record<string, ConstraintBound>
              ),
            variables: Array(a.buttons.length)
              .keys()
              .reduce(
                (acc, i) => {
                  acc[`button-${i}`] = a.buttons[i]!.reduce(
                    (acc2, light) => {
                      acc2[`light-${light}`] = 1;
                      return acc2;
                    },
                    { cost: 1 } as VariableCoefficients
                  );
                  return acc;
                },
                {} as Record<string, VariableCoefficients>
              ),
            ints: Array(a.buttons.length)
              .keys()
              .reduce(
                (acc, i) => {
                  acc[`button-${i}`] = true;
                  return acc;
                },
                {} as Record<string, boolean>
              )
          }) as SolveResult
      ),
      Stream.map(a => a.result),
      Stream.runSum
    );
  }
}

const firstPatterns = /\[(.*)\] (.*) \{(.*)\}/;
const parse = (line: string, part: Part): Parsed => {
  const [_, lights, buttons, joltages] = line.match(firstPatterns)!;
  return {
    lights:
      part === 1
        ? lights!.split("").map(c => +(c === "#"))
        : joltages!.split(",").map(joltage => +joltage),
    buttons: buttons!.split(" ").map(button =>
      button
        .slice(1, -1)
        .split(",")
        .map(num => +num)
    ),
    maxJoltage:
      part === 1
        ? 1
        : Math.max(...joltages!.split(",").map(joltage => +joltage))
  };
};

export const solution: DaySolution = {
  solve,
  examples: {
    part1: {
      input: `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`,
      expected: 7
    },
    part2: {
      input: `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`,
      expected: 33
    }
  }
};
