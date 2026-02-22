import { Schema } from "effect"

// Day represents a valid advent day (1-12 for 2025)
export const Day = Schema.Int.pipe(
  Schema.between(1, 12),
  Schema.brand("Day")
)

export type Day = typeof Day.Type

// Format day as two digits (e.g., 1 -> "01")
export const formatDay = (day: Day): string => 
  day.toString().padStart(2, "0")

// Parse a string to a Day
export const parseDay = Schema.decodeUnknown(Day)

// All valid days (1-12 for 2025)
export const allDays = (): readonly Day[] =>
  Array.from({ length: 12 }, (_, i) => (i + 1) as Day)
