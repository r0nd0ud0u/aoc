use std::path::Path;

use anyhow::Result;

pub fn star1<P: AsRef<Path>>(path: P) -> Result<i32> {
    let lines = rust_utils::utils::read_lines(path)?;
    let (_, output) = lines.iter().fold((50, 0), |(dial, output), item| {
        let mut chars = item.chars();

        let sign = match chars.next() {
            Some('R') => 1,
            Some('L') => -1,
            _ => return (dial, output),
        };

        let value: i32 = match chars.as_str().parse() {
            Ok(v) => v,
            Err(_) => return (dial, output),
        };

        let mut new_dial = dial + sign * value;
        new_dial = ((new_dial % 100) + 100) % 100;
        let new_output = if new_dial == 0 { output + 1 } else { output };

        (new_dial, new_output)
    });
    Ok(output)
}
