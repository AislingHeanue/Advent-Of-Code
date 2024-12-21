use std::collections::{HashMap, HashSet};

use priority_queue::PriorityQueue;

advent_of_code::solution!(18);

pub fn part_one(input: &str) -> Option<u32> {
    let coord_re = regex::Regex::new("([0-9]+),([0-9]+)").unwrap();
    let (mut max_x, mut max_y) = (7, 7);
    let mut num_steps = 12;
    let coords: &[(usize, usize)] = &input
        .lines()
        .map(|line| {
            coord_re
                .captures(line)
                .unwrap()
                .extract::<2>()
                .1
                .map(|num| num.parse().unwrap())
                .map(|num| {
                    if num > 6 {
                        (max_x, max_y) = (71, 71);
                        num_steps = 1024;
                    }
                    num
                })
        })
        .map(|res| (res[0], res[1]))
        .collect::<Vec<(usize, usize)>>()[..num_steps];

    let mut banned_points = HashSet::new();
    for c in coords {
        banned_points.insert(*c);
    }
    //println!("{:?}", coords);

    let mut scores = HashMap::new();
    for y in 0..max_y {
        for x in 0..max_x {
            scores.insert((x, y), 71 * 71);
        }
    }
    find_lowest_distance(banned_points, max_x, max_y, &mut scores, (0, 0));

    //println!("{:?}", scores);
    scores.get(&(max_x - 1, max_y - 1)).copied()
}

pub fn part_two(input: &str) -> Option<String> {
    let coord_re = regex::Regex::new("([0-9]+),([0-9]+)").unwrap();
    let (mut max_x, mut max_y) = (7, 7);
    let mut num_steps = 12;
    let coords: &[(usize, usize)] = &input
        .lines()
        .map(|line| {
            coord_re
                .captures(line)
                .unwrap()
                .extract::<2>()
                .1
                .map(|num| num.parse().unwrap())
                .map(|num| {
                    if num > 6 {
                        (max_x, max_y) = (71, 71);
                        num_steps = 1024;
                    }
                    num
                })
        })
        .map(|res| (res[0], res[1]))
        .collect::<Vec<(usize, usize)>>();

    //let mut reachable = HashSet::new();
    //for y in 0..max_y {
    //    for x in 0..max_x {
    //        reachable.insert((x, y));
    //    }
    //}
    for i in 0..71 * 71 {
        let sub_coords = &coords[..i];
        let mut banned_points = HashSet::new();
        for c in sub_coords {
            banned_points.insert(*c);
        }
        //println!("{:?}", coords);

        let reachable = find_can_reach(
            banned_points,
            max_x,
            max_y,
            //&mut reachable,
            //&mut reachable,
            (0, 0),
            (max_x - 1, max_y - 1),
        );

        //println!("{:?}", scores);
        if !reachable {
            return Some(format!("{},{}", coords[i - 1].0, coords[i - 1].1));
        }
    }

    None
}
fn find_lowest_distance(
    banned: HashSet<(usize, usize)>,
    max_x: usize,
    max_y: usize,
    m: &mut HashMap<(usize, usize), u32>,
    //reachable: &mut HashSet<(usize, usize)>,
    start_pos: (usize, usize),
) {
    let mut positions = PriorityQueue::new();
    positions.push(start_pos, -0_isize);
    m.insert(start_pos, 0);
    while !positions.is_empty() {
        let current_position = positions.pop().unwrap().0;
        for next_x_y in [
            (current_position.0 as i32 + 1, current_position.1 as i32),
            (current_position.0 as i32 - 1, current_position.1 as i32),
            (current_position.0 as i32, current_position.1 as i32 + 1),
            (current_position.0 as i32, current_position.1 as i32 - 1),
        ] {
            if next_x_y.0 < 0
                || next_x_y.0 >= max_x as i32
                || next_x_y.1 < 0
                || next_x_y.1 >= max_y as i32
            {
                continue;
            }
            let next = (next_x_y.0 as usize, next_x_y.1 as usize);
            if banned.contains(&next) {
                continue;
            }

            let cost_of_next_coming_from_here = 1 + m.get(&current_position).unwrap();

            let current_score_for_this_tile_and_direction = m.get(&next);
            match current_score_for_this_tile_and_direction {
                Some(n) if *n > cost_of_next_coming_from_here => {
                    m.insert(next, cost_of_next_coming_from_here);
                    positions.push(next, -(cost_of_next_coming_from_here as isize));
                }
                _ => {}
            }
        }
    }
}

fn find_can_reach(
    banned: HashSet<(usize, usize)>,
    max_x: usize,
    max_y: usize,
    start_pos: (usize, usize),
    end_pos: (usize, usize),
) -> bool {
    let mut positions = Vec::new();
    let mut m = HashSet::new();
    positions.push(start_pos);
    m.insert(start_pos);
    while let Some(current_position) = positions.pop() {
        for next_x_y in [
            (current_position.0 as i32 + 1, current_position.1 as i32),
            (current_position.0 as i32 - 1, current_position.1 as i32),
            (current_position.0 as i32, current_position.1 as i32 + 1),
            (current_position.0 as i32, current_position.1 as i32 - 1),
        ] {
            if next_x_y.0 < 0
                || next_x_y.0 >= max_x as i32
                || next_x_y.1 < 0
                || next_x_y.1 >= max_y as i32
            {
                continue;
            }
            let next = (next_x_y.0 as usize, next_x_y.1 as usize);
            if banned.contains(&next) {
                continue;
            }
            if next == end_pos {
                return true;
            }

            if m.insert(next) {
                positions.push(next);
            }
        }
    }
    false
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(22));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some("6,1".to_string()));
    }
}
