advent_of_code::solution!(12);

pub fn part_one(input: &str) -> Option<u32> {
    let regions = get_regions(input);
    Some(
        regions
            .into_iter()
            .map(|region| region.area * region.perimeter)
            .sum(),
    )
}

pub fn part_two(input: &str) -> Option<u32> {
    let regions = get_regions(input);
    Some(
        regions
            .into_iter()
            .map(|region| region.area * region.corners)
            .sum(),
    )
}

fn get_regions(input: &str) -> Vec<RegionInfo> {
    let mut regions = Vec::new() as Vec<RegionInfo>;
    let mut characters: Vec<Vec<(char, bool)>> = input
        .lines()
        .map(|line| line.chars().map(|c| (c, false)).collect())
        .collect();

    for y_grid in 0..characters.len() {
        for x_grid in 0..characters[0].len() {
            let c = characters[y_grid][x_grid].0;
            if characters[y_grid][x_grid].1 {
                continue;
            }
            characters[y_grid][x_grid].1 = true;

            let mut area = 1;
            let mut perimeter = 0;
            let mut corners = 0;
            let mut points = vec![(y_grid, x_grid)];
            let mut i = 0;
            while i < points.len() {
                let (y, x) = points[i];
                //println!("pondering {} {}", x, y);

                let mut surrounding = Vec::new();
                if x != 0 {
                    surrounding.push(Direction::Left);
                } else {
                    perimeter += 1;
                }

                if y != 0 {
                    surrounding.push(Direction::Up);
                } else {
                    perimeter += 1;
                }

                if x != characters[0].len() - 1 {
                    surrounding.push(Direction::Right);
                } else {
                    perimeter += 1;
                }

                if y != characters.len() - 1 {
                    surrounding.push(Direction::Down);
                } else {
                    perimeter += 1;
                }

                let mut connected_to_this = Vec::new();
                for (direction, (ys, xs)) in surrounding
                    .into_iter()
                    .map(|dir| (dir, dir.next_point(y, x)))
                {
                    if characters[ys][xs].0 == c {
                        connected_to_this.push(direction);
                        if !characters[ys][xs].1 {
                            characters[ys][xs].1 = true;
                            area += 1;
                            points.push((ys, xs));
                        }
                        //println!("{} {} ({}): has {} {}", x_grid, y_grid, c, perimeter, area);
                    } else {
                        perimeter += 1;
                    }
                }
                // get convex corners
                //println!("{}", connected_to_this.len());
                corners += match connected_to_this.len() {
                    0 => 4,
                    1 => 2,
                    2 => {
                        if connected_to_this[0].perpendicular(&connected_to_this[1]) {
                            1
                        } else {
                            0
                        }
                    }
                    _ => 0,
                };
                // get concave corners
                for dir in &connected_to_this {
                    if *dir == Direction::Up || *dir == Direction::Down {
                        for dir2 in &connected_to_this {
                            if *dir2 == Direction::Left || *dir2 == Direction::Right {
                                let (y1, x1) = dir.next_point(y, x);
                                let (y2, x2) = dir2.next_point(y1, x1);

                                if characters[y2][x2].0 != c {
                                    corners += 1;
                                }
                            }
                        }
                    }
                }
                i += 1;
            }

            regions.push(RegionInfo {
                area,
                perimeter,
                corners,
            });
            //println!("");
        }
    }
    //println!("{:?}", regions);
    regions
}

#[derive(Debug)]
struct RegionInfo {
    perimeter: u32,
    area: u32,
    corners: u32,
}

#[derive(Clone, Copy, PartialEq)]
enum Direction {
    Right,
    Down,
    Left,
    Up,
}

impl Direction {
    fn next_point(&self, y: usize, x: usize) -> (usize, usize) {
        match self {
            Self::Right => (y, x + 1),
            Self::Down => (y + 1, x),
            Self::Left => (y, x - 1),
            Self::Up => (y - 1, x),
        }
    }

    fn perpendicular(&self, dir: &Direction) -> bool {
        match self {
            Self::Up | Self::Down => *dir == Self::Left || *dir == Self::Right,
            Self::Left | Self::Right => *dir == Self::Up || *dir == Self::Down,
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(1930));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(1206));
    }
}

//
// RRRR
// RRRR
//   RRR
//   R
//
// M
// M
// MMM
