fn main() {
    // The rule of reference
    // 1. At any given time, you can have either one mutable reference or any number of immutable
    //    references
    // 2. References must always valid
    let s1 = String::from("hello");

    let len = calculate_length(&s1);

    println!("The length of '{}' is {}.", s1, len);

    let mut s = String::from("hello");
    {
        let r = &mut s;
    } // r1 goes out of scope here, so we can make a new reference with no problem
    let r = &mut s;
    change(&mut s);
    //r.push_str("!");
    println!("The value of s is: {}", s);
}

fn calculate_length(s: &String) -> usize {
    s.len()
}

fn change(some_string: &mut String) {
    some_string.push_str(", world");
}
