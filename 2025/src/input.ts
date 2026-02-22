import { Stream } from "effect";

/**
 * Input wrapper class that provides utility methods for parsing puzzle input.
 * Extend this class to add your own helper methods.
 */
export class Input {
  constructor(public readonly raw: string) {}

  /** Get the raw string */
  toString(): string {
    return this.raw;
  }

  /** Get trimmed input */
  get trimmed(): string {
    return this.raw.trim();
  }

  /** Split into lines (trimmed, empty lines removed) */
  get lines(): string[] {
    return this.raw.trim().split("\n");
  }

  get stream(): Stream.Stream<string> {
    return Stream.fromIterable(this.lines);
  }

  /** Split into lines including empty ones */
  get linesRaw(): string[] {
    return this.raw.split("\n");
  }

  /** Parse each line as a number */
  get numbers(): number[] {
    return this.lines.map(Number);
  }

  /** Parse each line as an integer */
  get integers(): number[] {
    return this.lines.map(line => parseInt(line, 10));
  }

  /** Get input as a 2D grid of characters */
  get grid(): string[][] {
    return this.lines.map(line => line.split(""));
  }

  /** Get input as a 2D grid of numbers */
  get numGrid(): number[][] {
    return this.lines.map(line => line.split("").map(Number));
  }

  /** Split by blank lines into groups */
  get groups(): string[][] {
    return this.raw
      .trim()
      .split("\n\n")
      .map(group => group.split("\n"));
  }

  /** Get all numbers in the input (including negatives) */
  get allNumbers(): number[] {
    const matches = this.raw.match(/-?\d+/g);
    return matches ? matches.map(Number) : [];
  }

  /** Split first line by delimiter */
  firstLine(delimiter: string | RegExp = ","): string[] {
    const line = this.lines[0];
    return line ? line.split(delimiter) : [];
  }

  streamSplit(delimiter: string | RegExp = ","): Stream.Stream<string> {
    return Stream.fromIterable(this.firstLine(delimiter));
  }

  /** Split first line by delimiter and parse as numbers */
  firstLineNumbers(delimiter: string | RegExp = ","): number[] {
    return this.firstLine(delimiter).map(Number);
  }
}
