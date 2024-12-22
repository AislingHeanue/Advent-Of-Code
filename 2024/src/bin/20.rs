use std::collections::HashMap;

advent_of_code::solution!(20);

pub fn part_one(input: &str) -> Option<u32> {
    solve(input, 2)
}

pub fn part_two(input: &str) -> Option<u32> {
    solve(input, 20)
}

fn solve(input: &str, limit: u32) -> Option<u32> {
    let grid: Vec<Vec<char>> = input.lines().map(|line| line.chars().collect()).collect();

    let mut start_pos = None;
    for y in 0..grid.len() {
        for x in 0..grid[0].len() {
            if grid[y][x] == 'S' {
                start_pos = Some((x, y));
            }
        }
    }
    let grid_with_numbers = get_numbers(&grid, start_pos.unwrap());

    let mut cheats = HashMap::new();
    for y in 0..grid.len() {
        for x in 0..grid[0].len() {
            if grid_with_numbers[y][x] == -1 {
                continue;
            }
            for y2 in y..grid.len() {
                for x2 in 0..grid[0].len() {
                    if y == y2 && x < x2 {
                        continue; // don't double count horizontal lines (edge case)
                    }
                    if grid_with_numbers[y2][x2] == -1 {
                        continue;
                    }
                    let d = distance((x, y), (x2, y2));
                    if d <= limit {
                        let savings =
                            (grid_with_numbers[y][x] - grid_with_numbers[y2][x2]).abs() - d as i32;
                        if savings > 0 {
                            *cheats.entry(savings).or_insert(0) += 1;
                        }
                    }
                }
            }
        }
    }

    //println!("{:?}", cheats);
    Some(
        cheats
            .into_iter()
            .filter(|(k, _v)| *k >= 100)
            .map(|(_k, v)| v)
            .sum(),
    )
}

fn distance((x1, y1): (usize, usize), (x2, y2): (usize, usize)) -> u32 {
    ((x1 as i32 - x2 as i32).abs() + (y1 as i32 - y2 as i32).abs()) as u32
}

fn get_numbers(grid: &[Vec<char>], start_pos: (usize, usize)) -> Vec<Vec<i32>> {
    let mut grid_with_numbers: Vec<Vec<i32>> = grid
        .iter()
        .map(|line| line.iter().map(|_c| -1).collect())
        .collect();
    let mut current_position = start_pos;
    let mut d = 0;
    grid_with_numbers[current_position.1][current_position.0] = d;
    loop {
        d += 1;
        for next in [
            (current_position.0 + 1, current_position.1),
            (current_position.0 - 1, current_position.1),
            (current_position.0, current_position.1 + 1),
            (current_position.0, current_position.1 - 1),
        ] {
            match grid[next.1][next.0] {
                '.' => {
                    if grid_with_numbers[next.1][next.0] == -1 {
                        current_position = next;
                        grid_with_numbers[current_position.1][current_position.0] = d;
                        break;
                    }
                }
                'E' => {
                    current_position = next;
                    grid_with_numbers[current_position.1][current_position.0] = d;
                    return grid_with_numbers;
                }
                _ => {}
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
        assert_eq!(result, Some(0));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(0));
    }
}
