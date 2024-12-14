advent_of_code::solution!(13);
use itertools::Itertools;
use ndarray::array;
use ndarray_linalg::Solve;
use regex::Regex;

pub fn part_one(input: &str) -> Option<u64> {
    solve(input, false)
}

pub fn part_two(input: &str) -> Option<u64> {
    solve(input, true)
}

pub fn solve(input: &str, part_b: bool) -> Option<u64> {
    let button_re = Regex::new("Button [AB]: X\\+([0-9]+), Y\\+([0-9]+)").unwrap();
    let prize_re = Regex::new("Prize: X=([0-9]+), Y=([0-9]+)").unwrap();
    Some(
        input
            .lines()
            //.filter(|line| !line.is_empty())
            .chunks(4)
            .into_iter()
            .map(|i| {
                let vals: Vec<&str> = i.collect();
                //println!("{:?}", vals);
                let (x1, y1) = button_re
                    .captures(vals[0])
                    .unwrap()
                    .extract::<2>()
                    .1
                    .map(|res| res.parse::<f64>().unwrap())
                    .into();

                let (x2, y2) = button_re
                    .captures(vals[1])
                    .unwrap()
                    .extract::<2>()
                    .1
                    .map(|res| res.parse::<f64>().unwrap())
                    .into();

                let (xb, yb) = prize_re
                    .captures(vals[2])
                    .unwrap()
                    .extract::<2>()
                    .1
                    .map(|res| {
                        let num = res.parse::<f64>().unwrap();
                        if part_b {
                            num + 10000000000000.0
                        } else {
                            num
                        }
                    })
                    .into();
                //println!("{} {}, {} {} -> {} {}", x1, y1, x2, y2, xb, yb);
                (array![[x1, x2], [y1, y2]], array![xb, yb])
            })
            .map(|(a, b)| {
                let res = a.solve(&b).unwrap();
                let (xa, ya) = (res[0].round(), res[1].round());
                (a, b, xa, ya)
            })
            .filter(|(a, b, xa, ya)| {
                let x = array![xa.round(), ya.round()];
                let new_b = a.dot(&x);
                b[0] == new_b[0] && b[1] == new_b[1]
            })
            .fold(0, |acc, (_a, _b, xa, ya)| acc + 3 * xa as u64 + ya as u64),
    )
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(480));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(875318608908));
    }
}
