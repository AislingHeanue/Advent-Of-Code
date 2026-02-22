import { Effect, Console, Match } from "effect"
import { parseArgs } from "./src/cli"
import { download, read, submit } from "./src/aocjs"
import { runDay, testDay } from "./src/runner"
import { getSolution, getAllSolutions } from "./src/days"
import { allDays } from "./src/day"

const program = Effect.gen(function* () {
  const command = yield* parseArgs(process.argv)

  yield* Match.value(command).pipe(
    Match.tag("DownloadCommand", (cmd) =>
      download(cmd.day)
    ),
    Match.tag("ReadCommand", (cmd) =>
      read(cmd.day)
    ),
    Match.tag("SolveCommand", (cmd) =>
      Effect.gen(function* () {
        const solution = getSolution(cmd.day)
        if (!solution) {
          yield* Console.error(`No solution found for day ${cmd.day}`)
          return
        }
        const result = yield* runDay(cmd.day, solution)
        
        // Submit if requested
        if (cmd.submit) {
          const answer = cmd.submit === 1 ? result.part1 : result.part2
          if (answer === undefined) {
            yield* Console.error(`Part ${cmd.submit} not implemented`)
            return
          }
          yield* submit(cmd.day, cmd.submit, String(answer))
        }
      })
    ),
    Match.tag("AllCommand", () =>
      Effect.gen(function* () {
        const solutions = getAllSolutions()
        for (const day of allDays()) {
          const solution = solutions.get(day)
          if (solution) {
            yield* runDay(day, solution)
          }
        }
      })
    ),
    Match.tag("TestCommand", (cmd) =>
      Effect.gen(function* () {
        const solutions = getAllSolutions()
        let totalPassed = 0
        let totalFailed = 0

        if (cmd.day !== undefined) {
          // Test single day
          const solution = getSolution(cmd.day)
          if (!solution) {
            yield* Console.error(`No solution found for day ${cmd.day}`)
            return
          }
          const results = yield* testDay(cmd.day, solution)
          totalPassed += results.filter(r => r.passed).length
          totalFailed += results.filter(r => !r.passed).length
        } else {
          // Test all days
          for (const day of allDays()) {
            const solution = solutions.get(day)
            if (solution) {
              const results = yield* testDay(day, solution)
              totalPassed += results.filter(r => r.passed).length
              totalFailed += results.filter(r => !r.passed).length
            }
          }
        }

        yield* Console.log(`\n${totalPassed} passed, ${totalFailed} failed`)
      })
    ),
    Match.exhaustive
  )
})

Effect.runPromise(program)
  .then(() => {
    process.exit(0)
  })
  .catch((error) => {
    console.error("Error:", error.message || error)
    process.exit(1)
  })
