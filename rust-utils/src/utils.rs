use std::io::BufRead;
use std::{fs::File, io::BufReader, path::Path};

use anyhow::Result;

pub fn read_lines<P: AsRef<Path>>(path: P) -> Result<Vec<String>> {
    let file = File::open(path)?;
    let reader = BufReader::new(file);

    let lines = reader.lines().collect::<std::io::Result<Vec<_>>>()?;
    Ok(lines)
}
