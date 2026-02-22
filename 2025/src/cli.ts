import { Effect, Data } from "effect"
import { Day, parseDay } from "./day"

// Command types
export class DownloadCommand extends Data.TaggedClass("DownloadCommand")<{
  day: Day
}> { }

export class ReadCommand extends Data.TaggedClass("ReadCommand")<{
  day: Day
}> { }

export class SolveCommand extends Data.TaggedClass("SolveCommand")<{
  day: Day
  submit?: 1 | 2
}> { }

export class AllCommand extends Data.TaggedClass("AllCommand")<Record<string, never>> { }

export class TestCommand extends Data.TaggedClass("TestCommand")<{
  day?: Day
}> { }

export type Command = DownloadCommand | ReadCommand | SolveCommand | AllCommand | TestCommand

// Parse errors
export class ParseError extends Data.TaggedError("ParseError")<{
  message: string
}> { }

// Parse command line arguments
export const parseArgs = (args: string[]): Effect.Effect<Command, ParseError> => {
  // Skip node and script path
  const cliArgs = args.slice(2)

  if (cliArgs.length === 0) {
    return Effect.fail(new ParseError({
      message: "No command specified. Available: download, read, solve, all"
    }))
  }

  const command = cliArgs[0]
  const rest = cliArgs.slice(1)

  return Effect.gen(function* () {
    switch (command) {
      case "download": {
        const dayStr = rest[0]
        if (dayStr === undefined) {
          return yield* Effect.fail(new ParseError({ message: "download requires a day number" }))
        }
        const day = yield* Effect.mapError(
          parseDay(parseInt(dayStr, 10)),
          () => new ParseError({ message: `Invalid day: ${dayStr}. Must be 1-12.` })
        )
        return new DownloadCommand({ day })
      }

      case "read": {
        const dayStr = rest[0]
        if (dayStr === undefined) {
          return yield* Effect.fail(new ParseError({ message: "read requires a day number" }))
        }
        const day = yield* Effect.mapError(
          parseDay(parseInt(dayStr, 10)),
          () => new ParseError({ message: `Invalid day: ${dayStr}. Must be 1-12.` })
        )
        return new ReadCommand({ day })
      }

      case "solve": {
        const dayStr = rest[0]
        if (dayStr === undefined) {
          return yield* Effect.fail(new ParseError({ message: "solve requires a day number" }))
        }
        const day = yield* Effect.mapError(
          parseDay(parseInt(dayStr, 10)),
          () => new ParseError({ message: `Invalid day: ${dayStr}. Must be 1-12.` })
        )

        // Check for --submit flag
        const submitIdx = rest.indexOf("--submit")
        let submit: 1 | 2 | undefined
        if (submitIdx !== -1) {
          const partStr = rest[submitIdx + 1]
          if (partStr !== undefined) {
            const part = parseInt(partStr, 10)
            if (part === 1 || part === 2) {
              submit = part
            }
          }
        }

        return new SolveCommand({ day, submit })
      }

      case "all": {
        return new AllCommand({})
      }

      case "test": {
        const dayStr = rest[0]
        if (dayStr === undefined) {
          // Test all days
          return new TestCommand({})
        }
        const day = yield* Effect.mapError(
          parseDay(parseInt(dayStr, 10)),
          () => new ParseError({ message: `Invalid day: ${dayStr}. Must be 1-12.` })
        )
        return new TestCommand({ day })
      }

      default:
        return yield* Effect.fail(new ParseError({
          message: `Unknown command: ${command}. Available: download, read, solve, test, all`
        }))
    }
  })
}
