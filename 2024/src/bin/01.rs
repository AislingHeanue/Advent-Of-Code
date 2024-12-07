advent_of_code::solution!(1);
use regex::Regex;

pub fn part_one(input: &str) -> Option<u32> {
    let (mut left, mut right) = get_vectors(input);
    left.sort();
    right.sort();
    Some(
        left.into_iter()
            .zip(right)
            .fold(0, |acc, entry| acc + entry.0.abs_diff(entry.1)),
    )
}

pub fn part_two(input: &str) -> Option<u32> {
    let (left, right) = get_vectors(input);
    Some(left.into_iter().fold(0, |acc, entry| {
        acc + entry * right.iter().filter(|num| **num == entry).count() as u32
    }))
}

pub fn get_vectors(input: &str) -> (Vec<u32>, Vec<u32>) {
    let re = Regex::new("([0-9]+)\\s+([0-9]+)").unwrap();
    input
        .lines()
        .map(|line| re.captures(line).unwrap().extract::<2>().1)
        .map(|captures| {
            (
                captures[0].parse::<u32>().unwrap(),
                captures[1].parse::<u32>().unwrap(),
            )
        })
        .collect()
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
