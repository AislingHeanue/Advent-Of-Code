use core::f32;
use std::cmp::Ordering;
advent_of_code::solution!(14);

pub fn part_one(input: &str) -> Option<u32> {
    let s = setup(input);
    let (max_x, max_y) = s.maximums;
    let (boundary_x, boundary_y) = s.boundaries;
    let mut current_points = s.current_points;
    for (x, y) in &mut current_points {
        x.next(max_x, 100);
        y.next(max_y, 100);
    }
    Some(
        current_points
            .into_iter()
            .fold(vec![0, 0, 0, 0], |mut v, (x, y)| {
                match (x.position.cmp(&boundary_x), y.position.cmp(&boundary_y)) {
                    (Ordering::Less, Ordering::Less) => v[0] += 1,
                    (Ordering::Greater, Ordering::Less) => v[1] += 1,
                    (Ordering::Less, Ordering::Greater) => v[2] += 1,
                    (Ordering::Greater, Ordering::Greater) => v[3] += 1,
                    _ => {}
                }
                v
            })
            .iter()
            .product(),
    )
}

pub fn part_two(input: &str) -> Option<u32> {
    let s = setup(input);
    let (max_x, max_y) = s.maximums;
    let (mut xs, mut ys): (Vec<Point>, Vec<Point>) = s.current_points.into_iter().unzip();
    let mut steps = 0;
    loop {
        steps += 1;
        for i in 0..xs.len() {
            xs[i].next(max_x, 1);
            ys[i].next(max_y, 1);
        }
        if xs.variance() * ys.variance() < 250000.0 {
            break;
        }
    }
    Some(steps)
}

fn setup(input: &str) -> Setup {
    let (max_x, max_y) = match &input[0..5] {
        "p=0,4" => (11, 7),
        _ => (101, 103),
    };

    let (boundary_x, boundary_y) = ((max_x - 1) / 2, (max_y - 1) / 2);

    let bot_re = regex::Regex::new("p=(-?[0-9]+),(-?[0-9]+) v=(-?[0-9]+),(-?[0-9]+)").unwrap();

    let current_points: Vec<(Point, Point)> = input
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
                Point {
                    position: res[0] as u32,
                    velocity: res[2],
                },
                Point {
                    position: res[1] as u32,
                    velocity: res[3],
                },
            )
        })
        .collect();
    Setup {
        maximums: (max_x, max_y),
        boundaries: (boundary_x, boundary_y),
        current_points,
    }
}

struct Setup {
    maximums: (u32, u32),
    boundaries: (u32, u32),
    current_points: Vec<(Point, Point)>,
}

#[derive(Clone, Copy)]
struct Point {
    position: u32,
    velocity: i32,
}

trait Variance {
    fn variance(&self) -> f32;
}

trait Mean {
    fn mean(&self) -> f32;
}

impl Point {
    fn next(&mut self, max: u32, times: u32) {
        let mut new_val = self.position as i32 + self.velocity * times as i32;
        while new_val >= max as i32 {
            new_val -= max as i32;
        }
        while new_val < 0 {
            new_val += max as i32;
        }
        self.position = new_val as u32;
    }
}

impl Variance for Vec<Point> {
    fn variance(&self) -> f32 {
        let m = self.mean();

        self.iter().fold(0.0, |acc, e| {
            acc + (e.position as f32 - m) * (e.position as f32 - m)
        }) / (self.len() as f32 - 1.0)
    }
}

impl Mean for Vec<Point> {
    fn mean(&self) -> f32 {
        self.iter().map(|val| val.position as f32).sum::<f32>() / self.len() as f32
    }
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
