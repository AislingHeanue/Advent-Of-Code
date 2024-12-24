use std::collections::{HashMap, HashSet};

use regex::Regex;

advent_of_code::solution!(24);

pub fn part_one(input: &str) -> Option<u64> {
    let first_re = Regex::new("([xy][0-9][0-9]): ([01])").unwrap();
    let second_re =
        Regex::new("([a-z0-9]{3}) ([A-Z]{2,3}) ([a-z0-9]{3}) -> ([a-z0-9]{3})").unwrap();

    let mut states: HashMap<String, bool> = input
        .lines()
        .filter(|line| line.len() < 7 && !line.is_empty())
        .map(|line| first_re.captures(line).unwrap().extract::<2>().1)
        .map(|matches| {
            (
                matches[0].to_string(),
                matches[1].parse::<u32>().unwrap() == 1,
            )
        })
        .collect();

    let mut instructions: HashSet<Instruction> = input
        .lines()
        .filter(|line| line.len() > 7)
        .map(|line| second_re.captures(line).unwrap().extract::<4>().1)
        .map(|matches| Instruction {
            left: matches[0].to_string(),
            operator: Operator::from_string(matches[1]),
            right: matches[2].to_string(),
            output: matches[3].to_string(),
        })
        .collect();

    while !instructions.is_empty() {
        let copy_instructions = instructions.clone();
        for instruction in copy_instructions {
            let left_state = states.get(&instruction.left);
            let right_state = states.get(&instruction.right);
            if let (Some(l), Some(r)) = (left_state, right_state) {
                let out = instruction.operator.evaluate(*l, *r);
                instructions.remove(&instruction);
                states.insert(instruction.output, out);
            }
        }
    }

    Some(
        states
            .into_iter()
            .filter(|(k, _v)| k.starts_with('z'))
            .map(|(k, v)| {
                let num = &k[1..].parse::<u64>().unwrap();
                (v as u64) << num
            })
            .sum(),
    )
}

pub fn part_two(_input: &str) -> Option<u32> {
    None
}

#[derive(Clone, PartialEq, Eq, Hash)]
struct Instruction {
    left: String,
    operator: Operator,
    right: String,
    output: String,
}

#[derive(Clone, PartialEq, Eq, Hash)]
enum Operator {
    Xor,
    Or,
    And,
}

impl Operator {
    fn from_string(s: &str) -> Operator {
        match s {
            "AND" => Self::And,
            "OR" => Self::Or,
            "XOR" => Self::Xor,
            _ => panic!("invalid operator"),
        }
    }
    fn evaluate(&self, left: bool, right: bool) -> bool {
        match self {
            Self::Or => left | right,
            Self::And => left & right,
            Self::Xor => left ^ right,
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(2024));
    }

    //#[test]
    //fn test_part_two() {
    //    let result = part_two(&advent_of_code::template::read_file("examples", DAY));
    //    assert_eq!(result, None);
    //}
}

// x00 AND y00 -> bdk 0th carry
// x00 XOR y00 -> z00 0th parity (z)
// x01 XOR y01 -> rsq 1st parity no carry
// bdk XOR rsq -> z01 1st parity with 0th carry (z)
// x01 AND y01 -> qkj 1st carry ignoring z0
// rsq AND bdk -> jnh 1st carry where z0 is 0
// qkj OR jnh -> spt 1st carry
//
