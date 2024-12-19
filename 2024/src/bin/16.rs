advent_of_code::solution!(16);

use std::collections::{HashMap, HashSet};

use priority_queue::PriorityQueue;

pub fn part_one(input: &str) -> Option<u32> {
    let (grid, mut super_cool_map, start_pos, end_pos) = parse_input(input);

    find_lowest_score(&grid, &mut super_cool_map, start_pos);

    let end_val = Direction::all()
        .into_iter()
        .map(|d| (end_pos.0, end_pos.1, d))
        .map(|(x, y, d)| (d, *super_cool_map.get(&(x, y, d)).unwrap_or(&u32::MAX)))
        .min_by(|(_, v1), (_, v2)| v1.cmp(v2))
        .unwrap();

    Some(end_val.1)
}

pub fn part_two(input: &str) -> Option<u32> {
    let (grid, mut super_cool_map, start_pos, end_pos) = parse_input(input);

    find_lowest_score(&grid, &mut super_cool_map, start_pos);

    // assumption: only one of the 'directions' of the end tile will have the lowest score
    let end_val = Direction::all()
        .into_iter()
        .map(|d| (end_pos.0, end_pos.1, d))
        .map(|(x, y, d)| (d, *super_cool_map.get(&(x, y, d)).unwrap_or(&u32::MAX)))
        .min_by(|(_, v1), (_, v2)| v1.cmp(v2))
        .unwrap();

    let mut marked_positions = HashSet::new();
    marked_positions.insert(end_pos);
    find_valid_backwards_paths(
        (end_pos.0, end_pos.1, end_val.0),
        &mut marked_positions,
        &super_cool_map,
    );

    Some(marked_positions.len() as u32)
}

type InputResult = (
    Vec<Vec<char>>,
    HashMap<(usize, usize, Direction), u32>,
    (usize, usize, Direction),
    (usize, usize),
);

fn parse_input(input: &str) -> InputResult {
    let grid: Vec<Vec<char>> = input.lines().map(|line| line.chars().collect()).collect();

    let mut super_cool_map = HashMap::new();

    let mut start_pos = (0, 0, Direction::Right);
    let mut end_pos = (0, 0);
    for y in 0..grid.len() {
        for x in 0..grid[0].len() {
            if grid[y][x] == 'S' {
                super_cool_map.insert((x, y, Direction::Right), 0);
                start_pos = (x, y, Direction::Right);
            }
            if grid[y][x] == 'E' {
                end_pos = (x, y);
            }
        }
    }

    (grid, super_cool_map, start_pos, end_pos)
}

fn find_lowest_score(
    grid: &[Vec<char>],
    m: &mut HashMap<(usize, usize, Direction), u32>,
    start_pos: (usize, usize, Direction),
) {
    let mut positions = PriorityQueue::new();
    positions.push(start_pos, -0_isize);
    while !positions.is_empty() {
        let current_position = positions.pop().unwrap().0;
        for direction in Direction::all() {
            let next_x_y = direction.next_position((current_position.0, current_position.1));
            let next = (next_x_y.0, next_x_y.1, direction);
            let (turn_cost, possible) = current_position.2.get_scores_for_turns(&direction);
            if !possible {
                // trying to turn 180 degrees
                continue;
            }
            let cost_of_next_coming_from_here = turn_cost + 1 + m.get(&current_position).unwrap();
            let next_grid_point = grid[next.1][next.0];
            if next_grid_point == '#' {
                // skip walls
                continue;
            }

            let current_score_for_this_tile_and_direction = m.get(&next);
            match current_score_for_this_tile_and_direction {
                Some(n) if *n < cost_of_next_coming_from_here => {}
                _ => {
                    m.insert(next, cost_of_next_coming_from_here);
                    positions.push(next, -(cost_of_next_coming_from_here as isize));
                }
            }
        }
    }
}

fn find_valid_backwards_paths(
    pos: (usize, usize, Direction),
    m: &mut HashSet<(usize, usize)>,
    scores: &HashMap<(usize, usize, Direction), u32>,
) {
    let this_score = scores.get(&pos).unwrap();
    let next = pos.2.opposite().next_position((pos.0, pos.1));

    // possible score differences:
    //   same direction: 1
    //   different direction: 1001
    //   opposite: not possible
    for direction in Direction::all() {
        if let Some(next_score) = scores.get(&(next.0, next.1, direction)) {
            match direction {
                a if a == pos.2.opposite() => continue,
                b if b == pos.2 => {
                    if *this_score as i32 - *next_score as i32 == 1 {
                        m.insert(next);
                        find_valid_backwards_paths((next.0, next.1, direction), m, scores)
                    }
                }
                _ => {
                    if *this_score as i32 - *next_score as i32 == 1001 {
                        m.insert(next);
                        find_valid_backwards_paths((next.0, next.1, direction), m, scores)
                    }
                }
            }
        }
    }
}

#[derive(PartialEq, Clone, Debug, Copy, Hash, Eq)]
enum Direction {
    Up,
    Right,
    Down,
    Left,
}

impl Direction {
    fn all() -> [Direction; 4] {
        [
            Direction::Up,
            Direction::Right,
            Direction::Down,
            Direction::Left,
        ]
    }

    fn get_scores_for_turns(&self, d: &Direction) -> (u32, bool) {
        match (self, d) {
            // same direction = no turn = free
            (x, y) if *x == *y => (0, true),
            // 180 turn not allowed
            (x, y) if *x == y.opposite() => (0, false),
            // remaining matches must be a 90 degree turn
            _ => (1000, true),
        }
    }

    fn opposite(&self) -> Direction {
        match self {
            Self::Up => Self::Down,
            Self::Right => Self::Left,
            Self::Down => Self::Up,
            Self::Left => Self::Right,
        }
    }

    fn next_position(&self, (x, y): (usize, usize)) -> (usize, usize) {
        match self {
            Self::Up => (x, y - 1),
            Self::Right => (x + 1, y),
            Self::Down => (x, y + 1),
            Self::Left => (x - 1, y),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(7036));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(45));
    }
}
