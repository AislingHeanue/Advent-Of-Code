use itertools::Itertools;

advent_of_code::solution!(15);

pub fn part_one(input: &str) -> Option<u32> {
    let grid: Vec<Vec<char>> = input
        .lines()
        .filter(|line| !line.is_empty() && line.starts_with("#"))
        .map(|line| line.chars().collect())
        .collect();
    let arrows: Vec<char> = input
        .lines()
        .filter(|line| !line.is_empty() && !line.starts_with("#"))
        .flat_map(|line| line.chars())
        .collect();

    Some(solve(grid, arrows))
}

pub fn part_two(input: &str) -> Option<u32> {
    let grid: Vec<Vec<char>> = input
        .lines()
        .filter(|line| !line.is_empty() && line.starts_with("#"))
        .map(|line| {
            line.chars()
                .flat_map(|c| match c {
                    '#' => vec!['#', '#'],
                    '.' => vec!['.', '.'],
                    '@' => vec!['@', '.'],
                    'O' => vec!['[', ']'],
                    _ => panic!("unknown character"),
                })
                .collect()
        })
        .collect();
    let arrows: Vec<char> = input
        .lines()
        .filter(|line| !line.is_empty() && !line.starts_with("#"))
        .flat_map(|line| line.chars())
        .collect();

    Some(solve(grid, arrows))
}

fn solve(mut grid: Vec<Vec<char>>, arrows: Vec<char>) -> u32 {
    let mut current_position = (0, 0);
    for (y, v) in grid.iter().enumerate() {
        if let Some(x) = v.iter().find_position(|c| **c == '@') {
            current_position = (x.0, y);
        }
    }
    for arrow in arrows {
        let mut spaces_to_check = vec![next_position(current_position, arrow)];
        let mut spaces_that_need_to_move = Vec::new();
        let mut index = 0;
        let mut is_valid = true;
        while index < spaces_to_check.len() {
            let (new_x, new_y) = spaces_to_check[index];
            match grid[new_y][new_x] {
                '.' => {
                    index += 1;
                    continue;
                }
                'O' => {
                    spaces_that_need_to_move.push((new_x, new_y));
                    spaces_to_check.push(next_position((new_x, new_y), arrow));
                }
                '[' => {
                    spaces_that_need_to_move.push((new_x, new_y));
                    spaces_to_check.push(next_position((new_x, new_y), arrow));
                    // only apply the fancy boxes rule for vertical pushes
                    if next_position((new_x, new_y), arrow).1 != new_y {
                        spaces_to_check.push(next_position((new_x + 1, new_y), arrow));
                        spaces_that_need_to_move.push((new_x + 1, new_y));
                    }
                }
                ']' => {
                    spaces_that_need_to_move.push((new_x, new_y));
                    spaces_to_check.push(next_position((new_x, new_y), arrow));
                    // only apply the fancy boxes rule for vertical pushes
                    if next_position((new_x, new_y), arrow).1 != new_y {
                        spaces_to_check.push(next_position((new_x - 1, new_y), arrow));
                        spaces_that_need_to_move.push((new_x - 1, new_y));
                    }
                }
                '#' => {
                    is_valid = false;
                    break;
                }
                _ => panic!("unknown character"),
            }
            index += 1;
            //println!("{} {:?} {}", index, spaces_to_check[index], arrow);
        }
        if is_valid {
            spaces_that_need_to_move
                .into_iter()
                .unique()
                .rev()
                .for_each(|space| {
                    let next = next_position(space, arrow);
                    let c = grid[space.1][space.0];
                    if c != '.' {
                        grid[space.1][space.0] = '.';
                        grid[next.1][next.0] = c;
                    }
                });
            let next_robot_position = next_position(current_position, arrow);
            grid[next_robot_position.1][next_robot_position.0] = '@';
            grid[current_position.1][current_position.0] = '.';
            current_position = next_robot_position;
        }
    }

    gps_score(grid)
}

fn next_position((x, y): (usize, usize), c: char) -> (usize, usize) {
    match c {
        'v' => (x, y + 1),
        '>' => (x + 1, y),
        '<' => (x - 1, y),
        '^' => (x, y - 1),
        _ => panic!("unknown character"),
    }
}

fn gps_score(grid: Vec<Vec<char>>) -> u32 {
    let mut total = 0;
    for (y, v) in grid.into_iter().enumerate() {
        for (x, c) in v.into_iter().enumerate() {
            if c == 'O' || c == '[' {
                total += x as u32 + 100 * y as u32
            }
        }
    }
    total
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(10092));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(9021));
    }
}
