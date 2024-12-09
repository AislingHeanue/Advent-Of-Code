advent_of_code::solution!(9);

pub fn part_one(input: &str) -> Option<u64> {
    let (mut drive_vec, mut num_free_space) = get_initial_drive(input);

    let mut left_free_index = 0;
    while num_free_space > 0 {
        let last = drive_vec.pop();
        if last == Some(-1) {
            num_free_space -= 1;
            continue;
        }
        while left_free_index < drive_vec.len() && drive_vec[left_free_index] != -1 {
            left_free_index += 1;
        }
        drive_vec[left_free_index] = last.unwrap();
        num_free_space -= 1;
    }
    Some(
        drive_vec
            .into_iter()
            .enumerate()
            .fold(0, |acc, (index, num)| (acc + index as u64 * num as u64)),
    )
}

pub fn part_two(input: &str) -> Option<u64> {
    let (mut drive_vec, _num_free_space) = get_initial_drive(input);

    let mut right_index = drive_vec.len() - 1;

    while right_index > 0 {
        //while drive_vec[left_index] != -1 {
        //    left_index += 1;
        //}
        //let mut free_space_here = 1;
        //let start_free_space = left_index;
        //while drive_vec[left_index] == -1 {
        //    left_index += 1;
        //    free_space_here += 1;
        //}
        //
        while drive_vec[right_index] == -1 && right_index > 0 {
            right_index -= 1;
        }
        let current_right_num = drive_vec[right_index];
        let mut length_to_move = 0;
        while right_index > 0 && drive_vec[right_index] == current_right_num {
            right_index -= 1;
            length_to_move += 1;
        }
        let mut scanning_left_index = 0;
        let mut current_free_block_size = 0;
        while scanning_left_index <= right_index {
            if drive_vec[scanning_left_index] == -1 {
                current_free_block_size += 1;
                if current_free_block_size == length_to_move {
                    for i in 0..length_to_move {
                        drive_vec[scanning_left_index - i] = drive_vec[right_index + 1 + i];
                        drive_vec[right_index + 1 + i] = -1;
                    }
                    scanning_left_index = right_index + 1; // trigger a break here
                }
            } else {
                current_free_block_size = 0;
            }
            scanning_left_index += 1;
        }
    }
    Some(
        drive_vec
            .into_iter()
            .enumerate()
            .map(|(i, num)| if num == -1 { (i, 0) } else { (i, num) })
            .fold(0, |acc, (index, num)| (acc + index as u64 * num as u64)),
    )
}

fn get_initial_drive(input: &str) -> (Vec<i32>, u32) {
    let mut drive_vec = Vec::new();
    let mut current_id: i32 = 0;
    let mut num_free_space = 0;
    for (i, char) in input
        .lines()
        .next()
        .unwrap()
        .chars()
        .map(|num| {
            let num_converted = num.to_digit(10);
            match num_converted {
                Some(n) => n,
                None => panic!("what is {:?}", num),
            }
        })
        .enumerate()
    {
        if i % 2 == 0 {
            for _j in 0..char {
                drive_vec.push(current_id);
            }
            current_id += 1;
        } else {
            for _j in 0..char {
                drive_vec.push(-1);
            }
            num_free_space += char;
        }
    }

    (drive_vec, num_free_space)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(1928));
    }

    #[test]
    fn test_part_two() {
        let result = part_two(&advent_of_code::template::read_file("examples", DAY));
        assert_eq!(result, Some(2858));
    }
}
