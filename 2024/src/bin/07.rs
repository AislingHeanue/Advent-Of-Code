advent_of_code::solution!(7);

pub fn part_one(input: &str) -> Option<u64> {
    Some(solve(input, false))
}

pub fn part_two(input: &str) -> Option<u64> {
    Some(solve(input, true))
}

pub fn solve(input: &str, is_b: bool) -> u64 {
    input
        .lines()
        .map(|line| line.split(": ").collect())
        .map(|line_split: Vec<&str>| {
            (
                line_split[0].parse::<u64>().unwrap(),
                line_split[1]
                    .split(" ")
                    .map(|val| val.parse::<u64>().unwrap())
                    .collect(),
            )
        })
        .filter(|(first, second): &(u64, Vec<u64>)| {
            does_tree_contain(*first, second.split_last().unwrap(), is_b)
        })
        .map(|(first, _second)| first)
        .sum::<u64>()
}

fn does_tree_contain(target: u64, (last_num, nums): (&u64, &[u64]), is_b: bool) -> bool {
    // numbers are evaluated left to right so
    // if 18 = a+b*c + 8, then 10 = a+b*c
    // if 18 = a+b*c * 3 and 3|18, then 6 = a+b*c
    if nums.is_empty() || *last_num > target {
        return *last_num == target;
    }
    let plus_condition = || does_tree_contain(target - last_num, nums.split_last().unwrap(), is_b);
    let product_condition = || {
        target % last_num == 0
            && *last_num != 0
            && does_tree_contain(target / last_num, nums.split_last().unwrap(), is_b)
    };
    let concat_condition = || {
        is_b && target % get_power_of_10(*last_num) == *last_num
            && does_tree_contain(
                target / get_power_of_10(*last_num),
                nums.split_last().unwrap(),
                is_b,
            )
    };
    plus_condition() || product_condition() || concat_condition()
}

fn get_power_of_10(mut num: u64) -> u64 {
    let mut power_of_10 = 1;
    while num > 0 {
        num /= 10;
        power_of_10 *= 10;
    }

    power_of_10
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(3749));
    }
    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(11387));
    }
}
