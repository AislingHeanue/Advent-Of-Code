use std::collections::{HashMap, HashSet};

use itertools::Itertools;

advent_of_code::solution!(23);

pub fn part_one(input: &str) -> Option<u32> {
    let line_re = regex::Regex::new("([a-z]{2})-([a-z]{2})").unwrap();
    let lines = input.lines();
    let mut connections = HashMap::new();
    let mut names = HashSet::new();
    for line in lines {
        let matches = line_re.captures(line).unwrap().extract::<2>().1;
        connections
            .entry(matches[0])
            .or_insert(Vec::new())
            .push(matches[1].to_string());
        connections
            .entry(matches[1])
            .or_insert(Vec::new())
            .push(matches[0].to_string());

        names.insert(matches[0]);
        names.insert(matches[1]);
    }
    let mut list = Vec::new();
    for k in connections.keys() {
        let sub_list = vec![k.to_string()];
        list.push(sub_list);
    }
    let list = cliques(list, &connections);
    let list = cliques(list, &connections);
    Some(
        list.iter()
            .filter(|v| v.iter().any(|s| s.starts_with('t')))
            .count() as u32,
    )
}

pub fn part_two(input: &str) -> Option<String> {
    let line_re = regex::Regex::new("([a-z]{2})-([a-z]{2})").unwrap();
    let lines = input.lines();
    let mut connections = HashMap::new();
    for line in lines {
        let matches = line_re.captures(line).unwrap().extract::<2>().1;
        connections
            .entry(matches[0])
            .or_insert(Vec::new())
            .push(matches[1].to_string());
        connections
            .entry(matches[1])
            .or_insert(Vec::new())
            .push(matches[0].to_string());
    }

    let mut list = Vec::new();
    for k in connections.keys() {
        let sub_list = vec![k.to_string()];
        list.push(sub_list);
    }
    let mut last_list = list.clone();
    while !list.is_empty() {
        last_list = list.clone();
        list = cliques(list, &connections);
    }
    //println!("{:?}", last_list);
    let last_set = last_list[0].clone();
    let mut last_vec: Vec<String> = last_set.into_iter().collect_vec();
    last_vec.sort();

    Some(last_vec.join(","))
}

fn cliques(input: Vec<Vec<String>>, connections: &HashMap<&str, Vec<String>>) -> Vec<Vec<String>> {
    let mut out = Vec::new();
    for v in input {
        for (key, list) in connections {
            if !v.contains(&key.to_string())
                && v.iter()
                    .all(|entry| list.contains(entry) && after_in_alphabet(entry, key))
            {
                let mut new_list = v.clone();
                new_list.push(key.to_string());
                out.push(new_list);
            }
        }
    }
    out
}

fn after_in_alphabet(first: &str, second: &str) -> bool {
    let mut v = [first, second];
    v.sort();
    first == v[0]
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(7));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some("co,de,ka,ta".to_string()));
    }
}
