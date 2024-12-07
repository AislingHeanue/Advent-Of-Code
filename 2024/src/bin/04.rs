advent_of_code::solution!(4);

pub fn part_one(input: &str) -> Option<u32> {
    let input: Vec<Vec<char>> = input.lines().map(|line| line.chars().collect()).collect();
    let mut total = 0;
    for i in 0..input.len() {
        for j in 0..input[0].len() {
            if j + 3 < input[0].len()
                && (
                    input[i][j],
                    input[i][j + 1],
                    input[i][j + 2],
                    input[i][j + 3],
                ) == ('X', 'M', 'A', 'S')
            {
                total += 1
            }
            if j >= 3
                && (
                    input[i][j],
                    input[i][j - 1],
                    input[i][j - 2],
                    input[i][j - 3],
                ) == ('X', 'M', 'A', 'S')
            {
                total += 1
            }
            if i + 3 < input.len()
                && (
                    input[i][j],
                    input[i + 1][j],
                    input[i + 2][j],
                    input[i + 3][j],
                ) == ('X', 'M', 'A', 'S')
            {
                total += 1
            }
            if i >= 3
                && (
                    input[i][j],
                    input[i - 1][j],
                    input[i - 2][j],
                    input[i - 3][j],
                ) == ('X', 'M', 'A', 'S')
            {
                total += 1
            }
            if j + 3 < input[0].len()
                && i + 3 < input.len()
                && (
                    input[i][j],
                    input[i + 1][j + 1],
                    input[i + 2][j + 2],
                    input[i + 3][j + 3],
                ) == ('X', 'M', 'A', 'S')
            {
                total += 1
            }
            if j + 3 < input[0].len()
                && i >= 3
                && (
                    input[i][j],
                    input[i - 1][j + 1],
                    input[i - 2][j + 2],
                    input[i - 3][j + 3],
                ) == ('X', 'M', 'A', 'S')
            {
                total += 1
            }
            if j >= 3
                && i + 3 < input.len()
                && (
                    input[i][j],
                    input[i + 1][j - 1],
                    input[i + 2][j - 2],
                    input[i + 3][j - 3],
                ) == ('X', 'M', 'A', 'S')
            {
                total += 1
            }
            if j >= 3
                && i >= 3
                && (
                    input[i][j],
                    input[i - 1][j - 1],
                    input[i - 2][j - 2],
                    input[i - 3][j - 3],
                ) == ('X', 'M', 'A', 'S')
            {
                total += 1
            }
        }
    }

    Some(total)
}

pub fn part_two(input: &str) -> Option<u32> {
    let input: Vec<Vec<char>> = input.lines().map(|line| line.chars().collect()).collect();
    let mut total = 0;
    for i in 0..input.len() - 2 {
        for j in 0..input[0].len() - 2 {
            if (
                input[i][j],
                input[i][j + 2],
                input[i + 1][j + 1],
                input[i + 2][j],
                input[i + 2][j + 2],
            ) == ('M', 'M', 'A', 'S', 'S')
            {
                total += 1;
            }
            if (
                input[i][j],
                input[i][j + 2],
                input[i + 1][j + 1],
                input[i + 2][j],
                input[i + 2][j + 2],
            ) == ('M', 'S', 'A', 'M', 'S')
            {
                total += 1;
            }
            if (
                input[i][j],
                input[i][j + 2],
                input[i + 1][j + 1],
                input[i + 2][j],
                input[i + 2][j + 2],
            ) == ('S', 'S', 'A', 'M', 'M')
            {
                total += 1;
            }
            if (
                input[i][j],
                input[i][j + 2],
                input[i + 1][j + 1],
                input[i + 2][j],
                input[i + 2][j + 2],
            ) == ('S', 'M', 'A', 'S', 'M')
            {
                total += 1;
            }
        }
    }

    Some(total)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(18));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(9));
    }
}
