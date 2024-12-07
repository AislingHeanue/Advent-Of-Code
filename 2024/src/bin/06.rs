use std::collections::HashSet;

advent_of_code::solution!(6);

pub fn part_one(input: &str) -> Option<u32> {
    let input: Vec<Vec<char>> = input.lines().map(|line| line.chars().collect()).collect();

    let mut start_position: (i32, i32) = (0, 0);
    for i in 0..input.len() {
        for j in 0..input[0].len() {
            if input[i][j] == '^' {
                start_position = (i.try_into().unwrap(), j.try_into().unwrap());
            }
        }
    }

    if start_position == (0, 0) {
        panic!("couldn't find ^ in the input")
    }
    let mut position = start_position;
    let mut visited: HashSet<(i32, i32)> = HashSet::new();
    visited.insert(position);
    let mut direction = Direction::Up;
    loop {
        let next_position = (
            position.0 + direction.get_offset().0,
            position.1 + direction.get_offset().1,
        );
        if next_position.0 < 0
            || next_position.1 < 0
            || next_position.0 >= input.len().try_into().unwrap()
            || next_position.1 >= input[0].len().try_into().unwrap()
        {
            break;
        }
        let next_char = input[next_position.0 as usize][next_position.1 as usize];
        if next_char == '#' {
            direction = direction.next_direction();
        } else {
            position = next_position;
            visited.insert(position);
        }
    }

    Some(visited.len().try_into().unwrap())
}

#[derive(Hash, PartialEq, Eq, Clone)]
enum Direction {
    Up,
    Down,
    Left,
    Right,
}

impl Direction {
    fn next_direction(self) -> Direction {
        match self {
            Self::Up => Self::Right,
            Self::Right => Self::Down,
            Self::Down => Self::Left,
            Self::Left => Self::Up,
        }
    }

    fn get_offset(&self) -> (i32, i32) {
        match self {
            Self::Up => (-1, 0),
            Self::Right => (0, 1),
            Self::Down => (1, 0),
            Self::Left => (0, -1),
        }
    }
}

pub fn part_two(input: &str) -> Option<u32> {
    let input: Vec<Vec<char>> = input.lines().map(|line| line.chars().collect()).collect();

    let mut start_position: (i32, i32) = (0, 0);
    for i in 0..input.len() {
        for j in 0..input[0].len() {
            if input[i][j] == '^' {
                start_position = (i.try_into().unwrap(), j.try_into().unwrap());
            }
        }
    }

    if start_position == (0, 0) {
        panic!("couldn't find ^ in the input")
    }
    let mut position = start_position;
    let mut visited: HashSet<(i32, i32)> = HashSet::new();
    let mut visited_with_direction: HashSet<((i32, i32), Direction)> = HashSet::new();
    visited.insert(position);
    let mut direction = Direction::Up;
    let mut total = 0;
    loop {
        let next_position = (
            position.0 + direction.get_offset().0,
            position.1 + direction.get_offset().1,
        );
        if next_position.0 < 0
            || next_position.1 < 0
            || next_position.0 >= input.len().try_into().unwrap()
            || next_position.1 >= input[0].len().try_into().unwrap()
        {
            break;
        }
        let next_char = input[next_position.0 as usize][next_position.1 as usize];
        if next_char == '#' {
            direction = direction.next_direction();
        } else {
            if visited.insert(next_position)
                && check_if_placing_a_wall_here_would_cause_it_to_loop(
                    &input,
                    &next_position,
                    position,
                    direction.clone().next_direction(),
                    visited_with_direction.clone(),
                )
            {
                total += 1;
            }
            position = next_position;
            visited_with_direction.insert((position, direction.clone()));
        }
    }
    Some(total)
}

fn check_if_placing_a_wall_here_would_cause_it_to_loop(
    input: &[Vec<char>],
    position_of_wall: &(i32, i32),
    mut position: (i32, i32),
    mut direction: Direction,
    mut visited_with_direction: HashSet<((i32, i32), Direction)>,
) -> bool {
    loop {
        let next_position = (
            position.0 + direction.get_offset().0,
            position.1 + direction.get_offset().1,
        );
        if next_position.0 < 0
            || next_position.1 < 0
            || next_position.0 >= input.len().try_into().unwrap()
            || next_position.1 >= input[0].len().try_into().unwrap()
        {
            return false;
        }
        let next_char = input[next_position.0 as usize][next_position.1 as usize];
        if next_char == '#' || next_position == *position_of_wall {
            direction = direction.next_direction();
        } else {
            position = next_position;
            if !visited_with_direction.insert((position, direction.clone())) {
                return true;
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(41));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(6));
    }
}
