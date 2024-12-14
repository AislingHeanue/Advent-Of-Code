use core::f32;
use std::collections::HashMap;

advent_of_code::solution!(14);

pub fn part_one(input: &str) -> Option<u32> {
    let (max_x, max_y) = match &input[0..5] {
        "p=0,4" => (11, 7),
        _ => (101, 103),
    };

    let (boundary_x, boundary_y) = ((max_x - 1) / 2, (max_y - 1) / 2);

    let steps = 100;
    let bot_re = regex::Regex::new("p=(-?[0-9]+),(-?[0-9]+) v=(-?[0-9]+),(-?[0-9]+)").unwrap();

    Some(
        input
            .lines()
            .map(|line| {
                bot_re
                    .captures(line)
                    .unwrap()
                    .extract::<4>()
                    .1
                    .map(|num| num.parse::<i32>().unwrap())
            })
            .map(|res| {
                (
                    (res[0] + res[2] * steps) % max_x,
                    (res[1] + res[3] * steps) % max_y,
                )
            })
            .map(|(mut x, mut y)| {
                if x < 0 {
                    x += max_x;
                }
                if y < 0 {
                    y += max_y;
                }
                (x, y)
            })
            .fold(HashMap::new(), |mut acc, (x, y)| {
                if (..boundary_x).contains(&x) && (..boundary_y).contains(&y) {
                    *acc.entry(Corner::TopLeft).or_insert(0) += 1;
                } else if ((boundary_x + 1)..).contains(&x) && (..boundary_y).contains(&y) {
                    *acc.entry(Corner::TopRight).or_insert(0) += 1;
                } else if (..boundary_x).contains(&x) && ((boundary_y + 1)..).contains(&y) {
                    *acc.entry(Corner::BottomLeft).or_insert(0) += 1;
                } else if ((boundary_x + 1)..).contains(&x) && ((boundary_y + 1)..).contains(&y) {
                    *acc.entry(Corner::BottomRight).or_insert(0) += 1;
                }

                acc
            })
            .into_iter()
            .fold(1, |acc, (_, val)| acc * val),
    )
}

pub fn part_two(input: &str) -> Option<u32> {
    let (max_x, max_y) = match &input[0..5] {
        "p=0,4" => (11, 7),
        _ => (101, 103),
    };
    let bot_re = regex::Regex::new("p=(-?[0-9]+),(-?[0-9]+) v=(-?[0-9]+),(-?[0-9]+)").unwrap();

    // current_state hold the current position and velocities of all particles, and the product of
    // the variance in x and y values, and the number of steps taken.
    type State = (Vec<(i32, i32, i32, i32)>, f32, u32);
    let mut current_state: State = (
        input
            .lines()
            .map(|line| {
                bot_re
                    .captures(line)
                    .unwrap()
                    .extract::<4>()
                    .1
                    .map(|num| num.parse::<i32>().unwrap())
            })
            .map(|res| (res[0], res[1], res[2], res[3]))
            .collect(),
        f32::INFINITY,
        0,
    );

    while current_state.1 > 250000.0 {
        current_state.2 += 1;
        current_state.0 = current_state
            .0
            .into_iter()
            .map(|res| {
                (
                    (res.0 + res.2) % max_x,
                    (res.1 + res.3) % max_y,
                    res.2,
                    res.3,
                )
            })
            .map(|(mut x, mut y, vx, vy)| {
                if x < 0 {
                    x += max_x;
                }
                if y < 0 {
                    y += max_y;
                }
                (x, y, vx, vy)
            })
            .collect();

        current_state.1 = current_state
            .0
            .iter()
            .fold((Vec::new(), Vec::new()), |(mut xs, mut ys), e| {
                xs.push(e.0);
                ys.push(e.1);
                (xs, ys)
            })
            .product_variance();
    }

    Some(current_state.2)
}

trait ProductVariance {
    fn product_variance(&self) -> f32;
}

impl ProductVariance for (Vec<i32>, Vec<i32>) {
    fn product_variance(&self) -> f32 {
        variance(&self.0) * variance(&self.1)
    }
}

fn mean(v: &[i32]) -> f32 {
    v.iter().map(|val| *val as f32).sum::<f32>() / v.len() as f32
}
fn variance(v: &[i32]) -> f32 {
    let m = mean(v);

    v.iter()
        .fold(0.0, |acc, e| acc + (*e as f32 - m) * (*e as f32 - m))
        / (v.len() as f32 - 1.0)
}

#[derive(PartialEq, Eq, Hash)]
enum Corner {
    TopLeft,
    TopRight,
    BottomLeft,
    BottomRight,
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(12));
    }
}
