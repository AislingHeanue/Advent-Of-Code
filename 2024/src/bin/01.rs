advent_of_code::solution!(1);

use std::collections::HashMap;

use regex::Regex;

pub fn part_one(input: &str) -> Option<u32> {
    let re = Regex::new("([0-9]+)\\s+([0-9]+)").unwrap();
    let mut left = Vec::new();
    let mut right = Vec::new();
    input.lines().for_each(|line| {
        let captures = re.captures(line).unwrap();
        left.push(captures[1].parse::<i32>().unwrap());
        right.push(captures[2].parse::<i32>().unwrap());
    });
    let mut total = 0;
    left.sort();
    right.sort();
    for i in 0..left.len() {
        total += (left[i] - right[i]).abs();
    }

    Some(total.try_into().unwrap())
}

pub fn part_two(input: &str) -> Option<u32> {
    let re = Regex::new("([0-9]+)\\s+([0-9]+)").unwrap();
    let mut left = Vec::new();
    let mut right: HashMap<u32, u32> = HashMap::new();
    input.lines().for_each(|line| {
        let captures = re.captures(line).unwrap();
        left.push(captures[1].parse::<u32>().unwrap());
        let right_key = &captures[2].parse::<u32>().unwrap();
        let right_val = right.get(right_key).or(Some(&0)).unwrap();
        right.insert(*right_key, right_val + 1);
    });
    let mut total = 0;
    for i in 0..left.len() {
        total += left[i] * right.get(&left[i]).or(Some(&0)).unwrap();
    }

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
