advent_of_code::solution!(17);

static OP_CODES: [fn((u8, &mut Puter)); 8] = [
    (|(operand, puter): (u8, &mut Puter)| {
        // 0
        puter.a >>= combo_operand(operand, puter);
        puter.pointer += 2;
    }),
    (|(operand, puter): (u8, &mut Puter)| {
        // 1
        puter.b ^= literal_operand(operand);
        puter.pointer += 2;
    }),
    (|(operand, puter): (u8, &mut Puter)| {
        // 2
        puter.b = combo_operand(operand, puter) % 8;
        puter.pointer += 2;
    }),
    (|(operand, puter): (u8, &mut Puter)| {
        // 3
        if puter.a != 0 {
            puter.pointer = operand;
        } else {
            puter.pointer += 2;
        }
    }),
    (|(_, puter): (u8, &mut Puter)| {
        // 4
        puter.b ^= puter.c;
        puter.pointer += 2;
    }),
    (|(operand, puter): (u8, &mut Puter)| {
        // 5
        puter.output.push((combo_operand(operand, puter) % 8) as u8);
        puter.pointer += 2;
    }),
    (|(operand, puter): (u8, &mut Puter)| {
        // 6
        puter.b = puter.a >> combo_operand(operand, puter);
        puter.pointer += 2;
    }),
    (|(operand, puter): (u8, &mut Puter)| {
        // 7
        puter.c = puter.a >> combo_operand(operand, puter);
        puter.pointer += 2;
    }),
];

pub fn part_one(input: &str) -> Option<String> {
    let mut puter = setup(input);
    loop {
        if puter.pointer + 1 >= puter.instructions.len() as u8 {
            break;
        }
        let instruction = puter.instructions[puter.pointer as usize];
        let operand = puter.instructions[puter.pointer as usize + 1];
        OP_CODES[instruction as usize]((operand, &mut puter));
    }

    let mut s = String::new();
    for (i, val) in puter.output.iter().enumerate() {
        s += &val.to_string();
        if i != puter.output.len() - 1 {
            s += ",";
        }
    }
    Some(s)
}

pub fn part_two(_input: &str) -> Option<u32> {
    None
    //let first_puter = setup(input);
    //let mut i = 0;
    //loop {
    //    let mut puter = first_puter.clone();
    //    puter.a = i;
    //    loop {
    //        if puter.pointer + 1 >= puter.instructions.len() as u8 {
    //            break;
    //        }
    //        let instruction = puter.instructions[puter.pointer as usize];
    //        let operand = puter.instructions[puter.pointer as usize + 1];
    //        OP_CODES[instruction as usize]((operand, &mut puter));
    //        if instruction == 5 {
    //            let out_len = puter.output.len();
    //            if puter.output[out_len - 1] != puter.instructions[out_len - 1] {
    //                break;
    //            }
    //        }
    //    }
    //
    //    let mut s = String::new();
    //    for (i, val) in puter.output.iter().enumerate() {
    //        s += &val.to_string();
    //        if i != puter.output.len() - 1 {
    //            s += ",";
    //        }
    //    }
    //    //println!("{}", s);
    //    if puter.output == puter.instructions {
    //        return Some(i);
    //    } else {
    //        //println!("{}", i);
    //        i += 1;
    //    }
    //}
}

fn setup(input: &str) -> Puter {
    // input parsing is the worst sometimes
    let lines = input.lines();
    let reg_re = regex::Regex::new("Register [ABC]: ([0-9]+)").unwrap();
    let registers: Vec<u32> = lines
        .clone()
        .take(3)
        .map(|l| {
            reg_re.captures(l).unwrap().extract::<1>().1[0]
                .parse::<u32>()
                .unwrap()
        })
        .collect();
    let instructions_line = lines.collect::<Vec<&str>>()[4];
    let instructions_line = instructions_line.split(" ").collect::<Vec<&str>>()[1];
    let instructions: Vec<u8> = instructions_line
        .split(",")
        .map(|s| s.parse::<u8>().unwrap())
        .collect();

    Puter {
        instructions,
        pointer: 0,
        a: registers[0],
        b: registers[1],
        c: registers[2],
        output: Vec::new(),
    }
}

fn literal_operand(i: u8) -> u32 {
    i as u32
}

fn combo_operand(i: u8, puter: &Puter) -> u32 {
    match i {
        0..=3 => i as u32,
        4 => puter.a,
        5 => puter.b,
        6 => puter.c,
        _ => panic!("not a valid number"),
    }
}

#[derive(Debug, Clone)]
struct Puter {
    instructions: Vec<u8>,
    pointer: u8,
    a: u32,
    b: u32,
    c: u32,
    output: Vec<u8>,
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_one() {
        let result = part_one(&advent_of_code::template::read_file_part(
            "examples", DAY, 1,
        ));
        assert_eq!(result, Some("4,6,3,5,6,3,5,2,1,0".to_string()));
    }

    //#[test]
    //fn test_part_two() {
    //    let result = part_two(&advent_of_code::template::read_file_part(
    //        "examples", DAY, 2,
    //    ));
    //    assert_eq!(result, Some(117440));
    //}
}
