import { Effect, Data, Console } from "effect"
import { Client } from "aocjs"
import TurndownService from "turndown"
import * as fs from "fs"
import * as path from "path"
import { Day, formatDay } from "./day"

// HTML to Markdown converter
const turndown = new TurndownService({
  headingStyle: "atx",
  codeBlockStyle: "fenced"
})

// Clean HTML before conversion - remove scripts, styles, and non-puzzle content
const cleanPuzzleHtml = (html: string): string => {
  let cleaned = html
    // Remove script tags and their content
    .replace(/<script\b[^>]*>[\s\S]*?<\/script>/gi, "")
    // Remove style tags and their content
    .replace(/<style\b[^>]*>[\s\S]*?<\/style>/gi, "")
    // Remove inline event handlers
    .replace(/\s+on\w+="[^"]*"/gi, "")
    .replace(/\s+on\w+='[^']*'/gi, "")
  
  // Extract just the article content (puzzle descriptions)
  const articleMatch = cleaned.match(/<article[\s\S]*<\/article>/gi)
  if (articleMatch) {
    cleaned = articleMatch.join("\n\n")
  }
  
  return cleaned
}

// Error types
export class AocTokenMissingError extends Data.TaggedError("AocTokenMissingError")<{
  message: string
}> {}

export class AocClientError extends Data.TaggedError("AocClientError")<{
  message: string
}> {}

// Get year from environment or default to 2025
const getYear = (): number => {
  const year = process.env.AOC_YEAR
  if (year) {
    const parsed = parseInt(year, 10)
    if (!isNaN(parsed)) return parsed
  }
  return 2025
}

// Get session token from environment
const getToken = (): Effect.Effect<string, AocTokenMissingError> =>
  Effect.gen(function* () {
    const token = process.env.AOC_SESSION
    if (!token) {
      return yield* Effect.fail(new AocTokenMissingError({
        message: "AOC_SESSION environment variable is not set. Set it to your adventofcode.com session cookie value."
      }))
    }
    return token
  })

// Create AOC client
const createClient = () =>
  Effect.gen(function* () {
    const token = yield* getToken()
    return new Client({ session: token })
  })

// Get file paths
const getInputPath = (day: Day): string =>
  path.join(process.cwd(), "data", "inputs", `${formatDay(day)}.txt`)

const getPuzzlePath = (day: Day): string =>
  path.join(process.cwd(), "data", "puzzles", `${formatDay(day)}.md`)

// Ensure directory exists
const ensureDir = (filePath: string): void => {
  const dir = path.dirname(filePath)
  if (!fs.existsSync(dir)) {
    fs.mkdirSync(dir, { recursive: true })
  }
}

// Download puzzle description only
export const downloadPuzzle = (day: Day) =>
  Effect.gen(function* () {
    const client = yield* createClient()
    const year = getYear()
    
    // Fetch puzzle description (returns main element HTML)
    const puzzleHtml = yield* Effect.tryPromise({
      try: () => client.getProblem(year, day as number),
      catch: (error) => new AocClientError({
        message: `Failed to fetch puzzle: ${error}`
      })
    })
    
    // Clean HTML and convert to Markdown
    const cleanedHtml = cleanPuzzleHtml(puzzleHtml)
    const puzzleMarkdown = turndown.turndown(cleanedHtml)
    
    const puzzlePath = getPuzzlePath(day)
    ensureDir(puzzlePath)
    fs.writeFileSync(puzzlePath, puzzleMarkdown)
    yield* Console.log(`Successfully wrote puzzle to "${puzzlePath}".`)
  })

// Download puzzle input and description
export const download = (day: Day) =>
  Effect.gen(function* () {
    const client = yield* createClient()
    const year = getYear()
    
    // Fetch input
    const input = yield* Effect.tryPromise({
      try: () => client.getInput(year, day as number),
      catch: (error) => new AocClientError({
        message: `Failed to fetch input: ${error}`
      })
    })
    
    const inputPath = getInputPath(day)
    ensureDir(inputPath)
    fs.writeFileSync(inputPath, input)
    yield* Console.log(`Successfully wrote input to "${inputPath}".`)
    
    // Also download puzzle description
    yield* downloadPuzzle(day)
  })

// Read puzzle description (displays local file)
export const read = (day: Day) =>
  Effect.gen(function* () {
    const puzzlePath = getPuzzlePath(day)
    
    if (fs.existsSync(puzzlePath)) {
      const content = fs.readFileSync(puzzlePath, "utf-8")
      yield* Console.log(content)
    } else {
      yield* Console.log(`Puzzle file not found at "${puzzlePath}". Run download first.`)
    }
  })

// Submit an answer
export const submit = (day: Day, part: 1 | 2, answer: string) =>
  Effect.gen(function* () {
    const client = yield* createClient()
    const year = getYear()
    
    const success = yield* Effect.tryPromise({
      try: () => client.submit(year, day as number, part, answer),
      catch: (error) => new AocClientError({
        message: `Failed to submit answer: ${error}`
      })
    })
    
    if (success) {
      yield* Console.log(`Part ${part} answer "${answer}" was correct!`)
      
      // Re-download puzzle to get part 2 description after part 1 is correct
      if (part === 1) {
        yield* Console.log(`Downloading part 2 description...`)
        yield* downloadPuzzle(day)
      }
    } else {
      yield* Console.log(`Part ${part} answer "${answer}" was incorrect.`)
    }
    
    return success
  })
