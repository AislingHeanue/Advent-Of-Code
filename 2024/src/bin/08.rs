advent_of_code::solution!(8);
use std::collections::{HashMap, HashSet};

type NodeMap = HashMap<char, Vec<(i32, i32)>>;

pub fn part_one(input: &str) -> Option<u32> {
    let (node_map, bound_x, bound_y) = make_node_map(input);
    let mut anti_nodes: HashSet<(i32, i32)> = HashSet::new();

    for (_k, v) in node_map {
        for i in 0..v.len() {
            for j in 0..v.len() {
                if i != j {
                    let new_anti_node = (
                        v[i].0 + 2 * (v[j].0 - v[i].0),
                        v[i].1 + 2 * (v[j].1 - v[i].1),
                    );
                    if new_anti_node.0 >= 0
                        && new_anti_node.0 < bound_y
                        && new_anti_node.1 >= 0
                        && new_anti_node.1 < bound_x
                    {
                        anti_nodes.insert(new_anti_node);
                    }
                }
            }
        }
    }

    Some(anti_nodes.len().try_into().unwrap())
}

pub fn part_two(input: &str) -> Option<u32> {
    let (node_map, bound_x, bound_y) = make_node_map(input);
    let mut anti_nodes: HashSet<(i32, i32)> = HashSet::new();
    for (_k, v) in node_map {
        for i in 0..v.len() {
            for j in 0..v.len() {
                if i != j {
                    let offset = (v[i].0 - v[j].0, v[i].1 - v[j].1);
                    let divisor: i32 = gcd::euclid_u32(
                        offset.0.abs().try_into().unwrap(),
                        offset.1.abs().try_into().unwrap(),
                    )
                    .try_into()
                    .unwrap();
                    let offset = (offset.0 / divisor, offset.1 / divisor);
                    let mut current_position_positive = v[i];
                    while current_position_positive.0 >= 0
                        && current_position_positive.0 < bound_y
                        && current_position_positive.1 >= 0
                        && current_position_positive.1 < bound_x
                    {
                        anti_nodes.insert(current_position_positive);
                        current_position_positive = (
                            current_position_positive.0 + offset.0,
                            current_position_positive.1 + offset.1,
                        )
                    }
                    let mut current_position_negative = (v[i].0 - offset.0, v[i].1 - offset.1);
                    while current_position_negative.0 >= 0
                        && current_position_negative.0 < bound_y
                        && current_position_negative.1 >= 0
                        && current_position_negative.1 < bound_x
                    {
                        anti_nodes.insert(current_position_negative);
                        current_position_negative = (
                            current_position_negative.0 + offset.0,
                            current_position_negative.1 + offset.1,
                        )
                    }
                }
            }
        }
    }

    Some(anti_nodes.len().try_into().unwrap())
}

fn make_node_map(input: &str) -> (NodeMap, i32, i32) {
    let mut node_map: NodeMap = HashMap::new();
    let (bound_y, bound_x): (i32, i32) = (
        input
            .lines()
            .collect::<Vec<&str>>()
            .len()
            .try_into()
            .unwrap(),
        input.lines().collect::<Vec<&str>>()[0]
            .len()
            .try_into()
            .unwrap(),
    );
    input
        .lines()
        .enumerate()
        .flat_map(|(index, line)| {
            line.chars()
                .enumerate()
                .map(|(sub_index, character)| ((index, sub_index), character))
                .collect::<Vec<((usize, usize), char)>>()
        })
        .for_each(|((y, x), character)| {
            if character != '.' {
                node_map
                    .entry(character)
                    .or_default()
                    .push((y.try_into().unwrap_or_default(), x.try_into().unwrap()));
            }
        });
    (node_map, bound_y, bound_x)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(14));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(34));
    }
}
