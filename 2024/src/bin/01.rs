advent_of_code::solution!(1);

use std::collections::HashMap;

use regex::Regex;

pub fn part_one(input: &str) -> Option<u32> {
    let re = Regex::new("([0-9]+)\\s+([0-9]+)").unwrap();
    let (mut left, mut right): (Vec<i32>, Vec<i32>) = input
        .lines()
        .map(|line| re.captures(line).unwrap())
        .map(|captures| {
            (
                captures[1].parse::<i32>().unwrap(),
                captures[2].parse::<i32>().unwrap(),
            )
        })
        .unzip();

    left.sort();
    right.sort();
    let total = left
        .into_iter()
        .zip(right)
        .fold(0, |acc, entry| acc + (entry.0 - entry.1).abs());

    Some(total.try_into().unwrap())
}

pub fn part_two(input: &str) -> Option<u32> {
    let re = Regex::new("([0-9]+)\\s+([0-9]+)").unwrap();
    let mut left = Vec::new();
    let mut right: HashMap<u32, u32> = HashMap::new();
    for line in input.lines() {
        let captures = re.captures(line).unwrap();
        left.push(captures[1].parse::<u32>().unwrap());
        *right
            .entry(captures[2].parse::<u32>().unwrap())
            .or_default() += 1;
    }
    let total = left.into_iter().fold(0, |acc, entry| {
        acc + entry * right.get(&entry).or(Some(&0)).unwrap()
    });

    Some(total)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(11));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(31));
    }
}
