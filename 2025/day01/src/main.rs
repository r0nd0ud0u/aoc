use std::path::PathBuf;

use anyhow::Result;
use day01::{star1::star1, star2::star2};

fn main() -> Result<()> {
    assert_eq!(
        star1(PathBuf::from(env!("CARGO_MANIFEST_DIR")).join("puzzles/example.txt"))?,
        3
    );
    println!(
        "{}",
        star1(PathBuf::from(env!("CARGO_MANIFEST_DIR")).join("puzzles/puzzle.txt"))?
    );
    assert_eq!(
        star2(PathBuf::from(env!("CARGO_MANIFEST_DIR")).join("puzzles/example.txt"))?,
        6
    );
    println!(
        "{}",
        star2(PathBuf::from(env!("CARGO_MANIFEST_DIR")).join("puzzles/puzzle.txt"))?
    );
    Ok(())
}
