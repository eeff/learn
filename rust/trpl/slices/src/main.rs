fn main() {
    // A string slice is a reference to part of a String
    let s = String::from("hello world");

    // exclusive end index
    let hello = &s[0..5];
    let world = &s[6..11];
    // inclusive end index
    let hello = &s[0..=4];
    let world = &s[6..=10];

    println!("{} {}", hello, world);

    let s = String::from("hello");
    let slice = &s[..2]; // same as &s[0 .. 2]
    let slice = &s[3..]; // same as &s[3 .. s.len()]
    let slice = &s[..]; // whole String
    println!("slice is: {}", s);

    // string literals are slices
    // type is &str
    let my_string_literal = "hello world";
    let my_string = String::from("hello world");
    let word = first_word(&my_string[..]);
    println!("first word is: {}", word);
    let word = first_word(my_string_literal);
    println!("first word is: {}", word);
}

fn first_word(s: &str) -> &str {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[0..i];
        }
    }
    &s[..]
}
