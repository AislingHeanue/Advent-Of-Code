import { Day } from "../day"
import "../types" // Load global types

import { solution as day01 } from "./01"
import { solution as day02 } from "./02"
import { solution as day03 } from "./03"
import { solution as day04 } from "./04"
import { solution as day05 } from "./05"
import { solution as day06 } from "./06"
import { solution as day07 } from "./07"
import { solution as day08 } from "./08"
import { solution as day09 } from "./09"
import { solution as day10 } from "./10"
import { solution as day11 } from "./11"
import { solution as day12 } from "./12"

// All solutions indexed by day
const solutions: Record<number, DaySolution> = {
  1: day01,
  2: day02,
  3: day03,
  4: day04,
  5: day05,
  6: day06,
  7: day07,
  8: day08,
  9: day09,
  10: day10,
  11: day11,
  12: day12,
}

export const getSolution = (day: Day): DaySolution | undefined =>
  solutions[day]

export const getAllSolutions = (): Map<Day, DaySolution> => {
  const map = new Map<Day, DaySolution>()
  for (const [dayNum, solution] of Object.entries(solutions)) {
    map.set(parseInt(dayNum, 10) as Day, solution)
  }
  return map
}
