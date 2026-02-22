import type { Input } from "./input"
import type { Day } from "./day"

declare global {
  // Part type
  type Part = 1 | 2

  // Solution function type: solve(input, part) where part defaults to 1
  type SolveFn<T> = (input: Input, part: Part) => T | undefined

  // Example test case
  interface Example {
    input: string
    expected: unknown
  }

  // Solution module structure
  interface DaySolution {
    solve: SolveFn<unknown>
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
