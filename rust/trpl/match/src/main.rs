fn main() {
    let five = Some(5);
    let six = plus_one(five);
    let none = plus_one(None);
    println!("six: {:?}", six);
    println!("none: {:?}", none);

    // matches are exhausitive
    // Placeholder _ match all
    let some_u8_value = 0u8;
    match some_u8_value {
        1 => println!("one"),
        2 => println!("three"),
        5 => println!("five"),
        7 => println!("seven"),
        _ => (),
    }
}

fn plus_one(x: Option<i32>) -> Option<i32> {
    // match is also an expression
    match x {
        None => None,
        // pattern bind to vlaues
        Some(i) => Some(i + 1),
    }
}
