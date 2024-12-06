advent_of_code::solution!(2);
use itertools::Itertools;
pub fn part_one(input: &str) -> Option<u32> {
    Some(
        input
            .lines()
            .into_iter()
            .map(|line| {
                line.split(" ")
                    .into_iter()
                    .map(|num| num.parse::<i32>().unwrap())
                    .tuple_windows()
                    .fold(
                        /*safe and not increasing or decreasing yet*/ (true, 0),
                        |acc, (first, second)| {
                            if !acc.0 {
                                return (false, 0);
                            }
                            if acc.1 * (first - second) < 0 {
                                // last first difference was a different sign to this one, not safe
                                //println!("changed sign, {}, {}", first, second);
                                return (false, 0);
                            } else {
                                let difference = first - second;
                                if (first - second).abs() < 1 || (first - second).abs() > 3 {
                                    //println!("bad difference, {}, {}", first, second);
                                    return (false, 0);
                                }
                                return (true, difference);
                            }
                        },
                    )
                    .0
            })
            .filter(|safe| *safe)
            .count()
            .try_into()
            .unwrap(),
    )
}

pub fn part_two(input: &str) -> Option<u32> {
    Some(
        input
            .lines()
            .into_iter()
            .map(|line| {
                line.split(" ")
                    .into_iter()
                    .map(|num| num.parse::<i32>().unwrap())
                    .collect()
            })
            .map(|num_list: Vec<i32>| {
                (0..num_list.len())
                    .map(|i| {
                        let mut smaller_num_list = num_list.clone();
                        smaller_num_list.remove(i);
                        smaller_num_list
                            .into_iter()
                            .tuple_windows()
                            .fold(
                                /*safe and not increasing or decreasing yet*/ (true, 0),
                                |acc, (first, second)| {
                                    if !acc.0 {
                                        return (false, 0);
                                    }
                                    if acc.1 * (first - second) < 0 {
                                        // last first difference was a different sign to this one, not safe
                                        //println!("changed sign, {}, {}", first, second);
                                        return (false, 0);
                                    } else {
                                        let difference = first - second;
                                        if (first - second).abs() < 1 || (first - second).abs() > 3
                                        {
                                            //println!("bad difference, {}, {}", first, second);
                                            return (false, 0);
                                        }
                                        return (true, difference);
                                    }
                                },
                            )
                            .0
                    })
                    .any(|this| this)
            })
            .filter(|safe| {
                //println!("{}", *safe);
                *safe
            })
            .count()
            .try_into()
            .unwrap(),
    )
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(2));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(5));
    }
}
