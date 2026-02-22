import type { Effect } from "effect"
import type { Input } from "./input"
import type { Day } from "./day"

declare global {
  // Part type
  type Part = 1 | 2

  // Solve function - a generator function that can yield* Effects
  // Written as: function*(input, part) { ... yield* someEffect; return value }
  type SolveGen<T> = (input: Input, part: Part) => Generator<any, T | undefined, any>

  // Example test case
  interface Example {
    input: string
    expected: unknown
  }

  // Solution module structure
  interface DaySolution {
    solve: SolveGen<unknown>
    examples?: {
      part1?: Example
      part2?: Example
    }
  }

  // Solution registry type
  type SolutionRegistry = Map<Day, DaySolution>

  // Test result type
  interface TestResult {
    day: Day
    part: Part
    passed: boolean
    expected: unknown
    actual: unknown
  }
}

export { }
