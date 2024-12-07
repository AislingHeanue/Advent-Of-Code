use std::collections::HashMap;

use itertools::Itertools;

advent_of_code::solution!(5);

pub fn part_one(input: &str) -> Option<u32> {
    let rule_re = regex::Regex::new("([0-9]+)\\|([0-9]+)").unwrap();
    let rule_map: HashMap<(u32, u32), bool> = rule_re
        .captures_iter(input)
        .map(|c| c.extract::<2>().1.map(|val| val.parse::<u32>().unwrap()))
        .map(|[first, second]| ((first, second), true))
        .collect();

    let page_re = regex::Regex::new("[0-9]{2}(,[0-9]{2})+").unwrap();
    Some(
        page_re
            .find_iter(input)
            .map(|m| m.as_str().split(",").map(|val| val.parse::<u32>().unwrap()))
            .filter(|m| {
                m.clone()
                    .tuple_windows()
                    .all(|(first, second)| !rule_map.contains_key(&(second, first)))
            })
            .map(|list| list.collect())
            .map(|list: Vec<u32>| list[(list.len() - 1) / 2])
            .sum(),
    )
}

pub fn part_two(input: &str) -> Option<u32> {
    let rule_re = regex::Regex::new("([0-9]+)\\|([0-9]+)").unwrap();
    let rule_map: HashMap<(u32, u32), bool> = rule_re
        .captures_iter(input)
        .map(|c| c.extract::<2>().1.map(|val| val.parse::<u32>().unwrap()))
        .map(|[first, second]| ((first, second), true))
        .collect();

    let page_re = regex::Regex::new("[0-9]{2}(,[0-9]{2})+").unwrap();
    Some(
        page_re
            .find_iter(input)
            .map(|m| m.as_str().split(",").map(|val| val.parse::<u32>().unwrap()))
            .filter(|m| {
                m.clone()
                    .tuple_windows()
                    .any(|(first, second)| rule_map.contains_key(&(second, first)))
            })
            .map(|list| {
                let mut new_list: Vec<u32> = list.collect();
                for _ in 0..new_list.len() {
                    for j in 0..new_list.len() - 1 {
                        if rule_map.contains_key(&(new_list[j + 1], new_list[j])) {
                            (new_list[j + 1], new_list[j]) = (new_list[j], new_list[j + 1]);
                        }
                    }
                }
                new_list
            })
            .map(|list: Vec<u32>| list[(list.len() - 1) / 2])
            .sum(),
    )
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(143));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(123));
    }
}
