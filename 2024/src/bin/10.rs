use itertools::Itertools;

advent_of_code::solution!(10);

pub fn part_one(input: &str) -> Option<u32> {
    get_count(input, true)
}

pub fn part_two(input: &str) -> Option<u32> {
    get_count(input, false)
}

pub fn get_count(input: &str, unique: bool) -> Option<u32> {
    let grid: Vec<Vec<u32>> = input
        .lines()
        .map(|line| line.chars().map(|c| c.to_digit(10).unwrap()).collect())
        .collect();

    let mut heads = Vec::new();
    for i in 0..grid.len() {
        for j in 0..grid[0].len() {
            if grid[i][j] == 0 {
                heads.push((i, j));
            }
        }
    }
    Some(
        heads
            .iter()
            .map(|start| {
                let s = get_reachable(0, start, &grid);
                if unique {
                    s.iter().unique().count()
                } else {
                    s.len()
                }
            })
            .sum::<usize>()
            .try_into()
            .unwrap(),
    )
}

fn get_reachable(
    current_num: u32,
    (x, y): &(usize, usize),
    grid: &Vec<Vec<u32>>,
) -> Vec<(usize, usize)> {
    let mut s = Vec::new();
    if current_num == 9 {
        s = vec![(*x, *y)];
        return s;
    }
    if *x != 0 && grid[x - 1][*y] == current_num + 1 {
        s.append(&mut get_reachable(current_num + 1, &(x - 1, *y), grid))
    }
    if *y != 0 && grid[*x][y - 1] == current_num + 1 {
        s.append(&mut get_reachable(current_num + 1, &(*x, y - 1), grid))
    }
    if x + 1 != grid.len() && grid[x + 1][*y] == current_num + 1 {
        s.append(&mut get_reachable(current_num + 1, &(x + 1, *y), grid))
    }
    if y + 1 != grid[0].len() && grid[*x][y + 1] == current_num + 1 {
        s.append(&mut get_reachable(current_num + 1, &(*x, y + 1), grid))
    }

    s
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(36));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(81));
    }
}
