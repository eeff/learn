fn main() {
    let some_u8_value = Some(0u8);
    // syntax candy for match
    if let Some(3) = some_u8_value {
        println!("three");
    } else {
        println!("not three");
    }
}
