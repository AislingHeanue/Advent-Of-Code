use std::collections::HashMap;

use itertools::Itertools;

advent_of_code::solution!(11);

pub fn part_one(input: &str) -> Option<u64> {
    solve(input, 25)
}

pub fn part_two(input: &str) -> Option<u64> {
    solve(input, 75)
}

pub fn solve(input: &str, steps: u32) -> Option<u64> {
    let mut counts: HashMap<u64, usize> = input
        .lines()
        .next()?
        .split(" ")
        .map(|num| num.parse::<u64>().unwrap())
        .counts();

    for _i in 0..steps {
        counts = map_blink(counts /*HashMap::new()*/);
    }

    Some(counts.values().sum::<usize>().try_into().unwrap())
}

// m is used here for per-value memoization, since this is an idea I had earlier in the question
// for optimisation, but it didn't pan out. I still want to keep it here so that I remember how to
// use it.
fn map_blink(
    in_map: HashMap<u64, usize>, /* mut m: HashMap<u64, Vec<u64>> */
) -> HashMap<u64, usize> {
    let mut out_map = HashMap::new();

    for (k, v) in in_map.iter() {
        //let new_vals = if let Some(vals) = m.get(k) {
        //    vals
        //} else {
        let new_vals = if *k == 0 {
            &vec![1]
        } else {
            let digits = k.ilog10() + 1;
            if digits % 2 == 0 {
                let power_of_10 = 10_u64.pow(digits / 2);
                let first_half = k / power_of_10;
                let second_half = k % power_of_10;
                &vec![first_half, second_half]
            } else {
                &vec![k * 2024]
            }
        };
        //m.insert(*k, new_vals.clone());
        //new_vals
        //};
        for val in new_vals {
            *out_map.entry(*val).or_insert(0) += v;
        }
    }
    out_map
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(55312));
    }
}
