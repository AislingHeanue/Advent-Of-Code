advent_of_code::solution!(22);

use std::collections::HashMap;

pub fn part_one(input: &str) -> Option<u64> {
    let (data, _map) = setup(input);
    Some(data.into_iter().sum())
}

pub fn part_two(input: &str) -> Option<u64> {
    let (_data, map) = setup(input);
    println!("{:?}", map.iter().max_by_key(|(_k, v)| **v).unwrap());
    Some(*map.iter().max_by_key(|(_k, v)| **v).unwrap().1)
}

fn setup(input: &str) -> (Vec<u64>, HashMap<(i64, i64, i64, i64), u64>) {
    let (vs, ms): (Vec<u64>, Vec<HashMap<(i64, i64, i64, i64), u64>>) = input
        .lines()
        .map(|line| line.parse::<u64>().unwrap())
        .map(|mut num| {
            let mut v = Vec::with_capacity(2000);
            let mut m = HashMap::new();
            for _ in 0..2000 {
                let new_num = evolve(num);
                v.push((new_num % 10, (new_num % 10) as i64 - (num % 10) as i64));
                num = new_num;
            }
            for i in 3..2000 {
                m.entry((v[i - 3].1, v[i - 2].1, v[i - 1].1, v[i].1))
                    .or_insert(v[i].0);
            }
            (num, m)
        })
        .unzip();
    let mut out_m = HashMap::new();
    for m in ms {
        for (key, val) in m {
            *out_m.entry(key).or_insert(0) += val;
        }
    }

    (vs, out_m)
}

fn mix(first: u64, second: u64) -> u64 {
    first ^ second
}

fn prune(num: u64) -> u64 {
    num % 16777216
}

fn evolve(num: u64) -> u64 {
    let first = prune(mix(num, num << 6));
    let second = prune(mix(first, first >> 5));
    prune(mix(second, second << 11))
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file_part(
            "examples", DAY, 1,
        ));
        assert_eq!(result, Some(37327623));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file_part(
            "examples", DAY, 2,
        ));
        assert_eq!(result, Some(23));
    }
}
