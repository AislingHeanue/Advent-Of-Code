import { Effect, Console, Duration } from "effect";
import * as fs from "fs";
import * as path from "path";
import { Day, formatDay } from "./day";
import { Input } from "./input";

// ANSI codes for formatting
const ANSI_BOLD = "\x1b[1m";
const ANSI_ITALIC = "\x1b[3m";
const ANSI_RESET = "\x1b[0m";
const ANSI_GREEN = "\x1b[32m";
const ANSI_RED = "\x1b[31m";

// Read input file
export const readInput = (day: Day): Effect.Effect<Input, Error> =>
  Effect.try({
    try: () => {
      const filePath = path.join(
        process.cwd(),
        "data",
        "inputs",
        `${formatDay(day)}.txt`
      );
      const raw = fs.readFileSync(filePath, "utf-8");
      return new Input(raw);
    },
    catch: error =>
      new Error(`Could not read input file for day ${day}: ${error}`)
  });

// Format duration
const formatDuration = (duration: Duration.Duration): string => {
  const millis = Duration.toMillis(duration);
  if (millis < 1) {
    return `${(millis * 1000).toFixed(1)}us`;
  } else if (millis < 1000) {
    return `${millis.toFixed(1)}ms`;
  } else {
    return `${(millis / 1000).toFixed(2)}s`;
  }
};

// Run a single part and time it
const runPart = <T>(
  partName: string,
  solve: SolveGen<T>,
  input: Input,
  part: Part
): Effect.Effect<T | undefined, unknown, unknown> =>
  Effect.gen(function* () {
    const start = performance.now();
    // Wrap the user's generator in Effect.gen
    const result = yield* Effect.gen(() => solve(input, part));
    const elapsed = performance.now() - start;
    const duration = Duration.millis(elapsed);

    if (result === undefined) {
      yield* Console.log(
        `${partName}: ${ANSI_ITALIC}not implemented${ANSI_RESET}`
      );
    } else if (result !== "skip") {
      const resultStr = String(result);
      if (resultStr.includes("\n")) {
        yield* Console.log(`${partName}: (${formatDuration(duration)})`);
        yield* Console.log(resultStr);
      } else {
        yield* Console.log(
          `${partName}: ${ANSI_BOLD}${resultStr}${ANSI_RESET} (${formatDuration(duration)})`
        );
      }
    }

    return result;
  });

// Result from running a day (for submit)
export interface DayResult {
  part1: unknown;
  part2: unknown;
}

// Run both parts of a day's solution
export const runDay = (
  day: Day,
  solution: DaySolution
): Effect.Effect<DayResult, unknown, unknown> =>
  Effect.gen(function* () {
    yield* Console.log(`\n--- Day ${formatDay(day)} ---`);

    const input = yield* readInput(day);

    // run sequentially to test execution times accurately
    const part1 = yield* runPart("Part 1", solution.solve, input, 1);
    const part2 = yield* runPart("Part 2", solution.solve, input, 2);

    return { part1, part2 };
  });

// Create the solution registry
export const createRegistry = (): SolutionRegistry => new Map();

// Register a solution
export const registerSolution = (
  registry: SolutionRegistry,
  day: Day,
  solution: DaySolution
): void => {
  registry.set(day, solution);
};

// Run tests for a single day
export const testDay = (
  day: Day,
  solution: DaySolution
): Effect.Effect<TestResult[], unknown, unknown> =>
  Effect.gen(function* () {
    const results: TestResult[] = [];

    if (!solution.examples) {
      yield* Console.log(
        `Day ${formatDay(day)}: ${ANSI_ITALIC}no examples defined${ANSI_RESET}`
      );
      return results;
    }

    // Test Part 1
    if (solution.examples.part1) {
      const input = new Input(solution.examples.part1.input);
      const actual = yield* Effect.gen(() => solution.solve(input, 1));
      const passed = actual === solution.examples.part1.expected;
      results.push({
        day,
        part: 1,
        passed,
        expected: solution.examples.part1.expected,
        actual
      });

      if (passed) {
        yield* Console.log(
          `Day ${formatDay(day)} Part 1: ${ANSI_GREEN}PASS${ANSI_RESET}`
        );
      } else {
        yield* Console.log(
          `Day ${formatDay(day)} Part 1: ${ANSI_RED}FAIL${ANSI_RESET} (expected ${solution.examples.part1.expected}, got ${actual})`
        );
      }
    }

    // Test Part 2
    if (solution.examples.part2) {
      const input = new Input(solution.examples.part2.input);
      const actual = yield* Effect.gen(() => solution.solve(input, 2));
      const passed = actual === solution.examples.part2.expected;
      results.push({
        day,
        part: 2,
        passed,
        expected: solution.examples.part2.expected,
        actual
      });

      if (passed) {
        yield* Console.log(
          `Day ${formatDay(day)} Part 2: ${ANSI_GREEN}PASS${ANSI_RESET}`
        );
      } else {
        yield* Console.log(
          `Day ${formatDay(day)} Part 2: ${ANSI_RED}FAIL${ANSI_RESET} (expected ${solution.examples.part2.expected}, got ${actual})`
        );
      }
    }

    return results;
  });

// Re-export Input for convenience
export { Input } from "./input";
