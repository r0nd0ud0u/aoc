use std::path::PathBuf;

use anyhow::Result;
use day01::star1::star1;

fn main() -> Result<()> {
    let test1 = star1(PathBuf::from(env!("CARGO_MANIFEST_DIR")).join("puzzles/example.txt"))?;
    assert_eq!(test1, 3);
    println!(
        "{}",
        star1(PathBuf::from(env!("CARGO_MANIFEST_DIR")).join("puzzles/puzzle.txt"))?
    );
    Ok(())
}
