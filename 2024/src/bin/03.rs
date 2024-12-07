advent_of_code::solution!(3);

pub fn part_one(input: &str) -> Option<u32> {
    regex::Regex::new("mul\\(([0-9]{1,3}),([0-9]{1,3})\\)")
        .unwrap()
        .captures_iter(input)
        .map(|c| c.extract::<2>().1.map(|num| num.parse::<u32>().unwrap()))
        .map(|[num1, num2]| num1 * num2)
        .reduce(|acc, e| acc + e)
}

pub fn part_two(input: &str) -> Option<u32> {
    let mul_re = regex::Regex::new("mul\\(([0-9]{1,3}),([0-9]{1,3})\\)").unwrap();
    let do_re = regex::Regex::new("do\\(\\)").unwrap();
    let dont_re = regex::Regex::new("don't\\(\\)").unwrap();

    let mut locations: Vec<Location> = mul_re
        .captures_iter(input)
        .map(|c| (c.get(0).unwrap().start(), c.extract::<2>().1))
        .map(|(a, [b, c])| Location::Mul(a, b.parse::<u32>().unwrap() * c.parse::<u32>().unwrap()))
        .collect();

    locations.append(
        &mut do_re
            .find_iter(input)
            .map(|c| Location::Do(c.start()))
            .collect(),
    );

    locations.append(
        &mut dont_re
            .find_iter(input)
            .map(|c| Location::Dont(c.start()))
            .collect(),
    );

    locations.sort_by(|first, second| {
        let first = match first {
            Location::Do(s) => s,
            Location::Dont(s) => s,
            Location::Mul(s, _) => s,
        };
        let second = match second {
            Location::Do(s) => s,
            Location::Dont(s) => s,
            Location::Mul(s, _) => s,
        };
        first.cmp(second)
    });

    Some(
        locations
            .iter()
            .fold((0, true), |(total, doing), e| match e {
                Location::Do(_) => (total, true),
                Location::Dont(_) => (total, false),
                Location::Mul(_, num) => {
                    if doing {
                        (total + num, doing)
                    } else {
                        (total, doing)
                    }
                }
            })
            .0,
    )
}

enum Location {
    Do(usize),
    Dont(usize),
    Mul(usize, u32),
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file_part(
            "examples", DAY, 1,
        ));
        assert_eq!(result, Some(161));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file_part(
            "examples", DAY, 2,
        ));
        assert_eq!(result, Some(48));
    }
}
