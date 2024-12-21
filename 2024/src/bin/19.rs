use std::collections::HashMap;

advent_of_code::solution!(19);

pub fn part_one(input: &str) -> Option<u32> {
    let lines: Vec<&str> = input.lines().collect();
    let towels: Vec<&str> = lines[0].split(", ").collect();
    let designs = &lines[2..];

    let mut m = HashMap::new();
    Some(
        designs
            .iter()
            .filter(|design| {
                get_num_possible(
                    design,
                    &towels
                        .clone()
                        .into_iter()
                        .filter(|towel| design.contains(*towel))
                        .collect(),
                    &mut m,
                ) > 0
            })
            .count() as u32,
    )
}

pub fn part_two(input: &str) -> Option<u64> {
    let lines: Vec<&str> = input.lines().collect();
    let towels: Vec<&str> = lines[0].split(", ").collect();
    let designs = &lines[2..];

    let mut m = HashMap::new();
    Some(
        designs
            .iter()
            .map(|design| {
                get_num_possible(
                    design,
                    &towels
                        .clone()
                        .into_iter()
                        .filter(|towel| design.contains(*towel))
                        .collect(),
                    &mut m,
                )
            })
            .sum(),
    )
}

fn get_num_possible<'a>(design: &'a str, towels: &Vec<&str>, m: &mut HashMap<&'a str, u64>) -> u64 {
    if let Some(res) = m.get(design) {
        return *res;
    }

    let res = towels
        .iter()
        .filter(|towel| design.starts_with(*towel))
        .map(|towel| {
            if design == *towel {
                1
            } else {
                get_num_possible(design.strip_prefix(*towel).unwrap(), towels, m)
            }
        })
        .sum();
    m.insert(design, res);

    res
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(6));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(16));
    }
}
