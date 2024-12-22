use std::{cmp::Ordering, collections::HashMap, hash::Hash};

use itertools::Itertools;

advent_of_code::solution!(21);

pub fn part_one(input: &str) -> Option<u64> {
    solve(input, 2)
}

pub fn part_two(input: &str) -> Option<u64> {
    solve(input, 25)
}

fn solve(input: &str, remotes: u64) -> Option<u64> {
    // small gird order of operations: right, up, down, left so we can't hit down-left ever
    // big grid order of operations: right, up, down, left so we can't hit up-left ever
    let mut arrows_pad = HashMap::new();
    arrows_pad.insert('^', (1, 0));
    arrows_pad.insert('A', (2, 0));
    arrows_pad.insert('<', (0, 1));
    arrows_pad.insert('v', (1, 1));
    arrows_pad.insert('>', (2, 1));

    let mut number_pad = HashMap::new();
    number_pad.insert('7', (0, 0));
    number_pad.insert('8', (1, 0));
    number_pad.insert('9', (2, 0));
    number_pad.insert('4', (0, 1));
    number_pad.insert('5', (1, 1));
    number_pad.insert('6', (2, 1));
    number_pad.insert('1', (0, 2));
    number_pad.insert('2', (1, 2));
    number_pad.insert('3', (2, 2));
    number_pad.insert('0', (1, 3));
    number_pad.insert('A', (2, 3));

    let mut current_map = None;
    for _ in 0..remotes {
        current_map = Some(distance_pairs(&arrows_pad, current_map, &(0, 0)));
    }
    let m3 = distance_pairs(&number_pad, current_map, &(0, 3));

    Some(
        input
            .lines()
            .map(|line| {
                let num = &line[..3].parse::<u64>().unwrap();
                (*num, line.chars().collect::<Vec<char>>())
            })
            .map(|(num, chars)| {
                let mut path_size = 0;
                let mut current_position = 'A';
                for next_position in chars {
                    path_size += m3.get(&(current_position, next_position)).unwrap();
                    current_position = next_position;
                }

                num * path_size
            })
            .sum(),
    )
}

fn distance_pairs(
    pad: &HashMap<char, (usize, usize)>,
    existing_map: Option<HashMap<(char, char), u64>>,
    dont_go_here: &(usize, usize),
) -> HashMap<(char, char), u64> {
    let mut m = HashMap::new();
    for (k1, v1) in pad {
        for (k2, v2) in pad {
            let dirs = get_directions(v1, v2, dont_go_here);

            m.insert(
                (*k1, *k2),
                dirs.into_iter()
                    .map(|d| get_steps(&d, existing_map.clone()))
                    .min()
                    .unwrap(),
            );
        }
    }
    m
}

fn get_steps(d: &Vec<Direction>, existing_map: Option<HashMap<(char, char), u64>>) -> u64 {
    // input: ><^
    // then we would need A to >, > to <, < to ^, ^ to A, click A
    match existing_map {
        None => {
            let mut new_d = d.clone(); // 1 for each button press, plus 'A'
            new_d.push(Direction::A);
            new_d.len() as u64
        }
        Some(m) => {
            let mut total = 0;
            let mut current_position = 'A';
            for next_position in d {
                total += m.get(&(current_position, next_position.char())).unwrap();
                current_position = next_position.char();
            }
            total += m.get(&(current_position, 'A')).unwrap();

            total
        }
    }
}

// napkin debugging section:

// got:
//                  3                                7            9                 A
//              ^   A         ^^          <<         A       >>   A        vvv      A
//          <   A > A     <   AA    v <   AA   >>  ^ A    v  AA ^ A  v <   AAA >  ^ A
// 379A: v<<A>>^AvA^A  v<<A>>^AA  v<A<A>>^AA  vAA^<A>A  v<A>^AA<A>Av<A<A>>^AAAvA^<A>A
// want:
//                  3                            7                9                 A
//              ^   A           <<        ^^     A           >>   A        vvv      A
//          <   A > A    v <<   AA   >  ^ AA   > A        v  AA ^ A  v <   AAA >  ^ A
// 379A: <v<A>>^AvA^A  <vA<AA>>^AA  vA<^A>AA  vA^A      <vA>^AA<A>A<v<A>A>^AAAvA<^A>A

#[derive(Clone, Hash, PartialEq, Eq)]
enum Direction {
    Up,
    Left,
    Down,
    Right,
    A,
}

impl Direction {
    fn next(&self, (x, y): (usize, usize)) -> (usize, usize) {
        match self {
            Self::Right => (x + 1, y),
            Self::Up => (x, y - 1),
            Self::Down => (x, y + 1),
            Self::Left => (x - 1, y),
            Self::A => panic!("next A"),
        }
    }
    fn char(&self) -> char {
        match self {
            Self::Right => '>',
            Self::Up => '^',
            Self::Down => 'v',
            Self::Left => '<',
            Self::A => 'A',
        }
    }
}

fn get_directions(
    (from_x, from_y): &(usize, usize),
    (to_x, to_y): &(usize, usize),
    dont_go_here: &(usize, usize),
) -> Vec<Vec<Direction>> {
    let mut directions = Vec::new();
    match from_x.cmp(to_x) {
        Ordering::Greater => {
            for _ in 0..from_x - to_x {
                directions.push(Direction::Left);
            }
        }
        Ordering::Less => {
            for _ in 0..to_x - from_x {
                directions.push(Direction::Right);
            }
        }
        Ordering::Equal => {}
    }
    match from_y.cmp(to_y) {
        Ordering::Greater => {
            for _ in 0..from_y - to_y {
                directions.push(Direction::Up);
            }
        }
        Ordering::Less => {
            for _ in 0..to_y - from_y {
                directions.push(Direction::Down);
            }
        }
        Ordering::Equal => {}
    }

    directions
        .clone()
        .into_iter()
        .permutations(directions.len())
        .filter(|directions| {
            let mut current_position = (*from_x, *from_y);
            for direction in directions {
                current_position = direction.next(current_position);
                // discard any paths through the illegal points
                if current_position == *dont_go_here {
                    return false;
                }
            }

            true
        })
        .unique()
        .collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(126384));
    }
}
