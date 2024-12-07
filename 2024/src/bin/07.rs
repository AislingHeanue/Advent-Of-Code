advent_of_code::solution!(7);

pub fn part_one(input: &str) -> Option<u64> {
    Some(
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
                does_tree_contain(*first, second.clone(), false)
            })
            .map(|(first, _second)| first)
            .sum::<u64>(),
    )
}

fn does_tree_contain(target: u64, mut nums: Vec<u64>, is_b: bool) -> bool {
    // numbers are evaluated left to right so
    // if 18 = a+b*c + 8, then 10 = a+b*c
    // if 18 = a+b*c * 0, then 18 = a+b*c
    // if 18 = a+b*c * 3 and 3|18, then 6 = a+b*c
    if nums.len() == 1 {
        return target == nums[0];
    }
    // assume nums.len() is never 0
    let last_num = nums.pop().unwrap();

    if last_num > target {
        return false;
    }
    let mut power_of_10 = 1;
    let mut reduced_last_num = last_num;
    while reduced_last_num > 0 {
        reduced_last_num /= 10;
        power_of_10 *= 10;
    }

    does_tree_contain(target - last_num, nums.clone(), is_b)
        || (target % last_num == 0
            && last_num != 0
            && does_tree_contain(target / last_num, nums.clone(), is_b))
        || (is_b
            && target % power_of_10 == last_num
            && does_tree_contain((target - last_num) / power_of_10, nums, is_b))
        || (last_num == 0 && target == last_num)
}

pub fn part_two(input: &str) -> Option<u64> {
    Some(
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
                does_tree_contain(*first, second.clone(), true)
            })
            .map(|(first, _second)| first)
            .sum::<u64>(),
    )
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
