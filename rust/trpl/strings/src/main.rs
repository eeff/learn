#[allow(unused_variables)]
fn main() {
    // both String and string slice are UTF-8 encoded

    // creation
    let s = String::new; // empty string
    let s = "initial contents".to_string();
    let s = String::from("initial contents");

    // append
    let mut s = String::from("foo");
    let s2 = "bar";
    s.push_str(s2);
    println!("s is: {}", s);
    println!("s2 is: {}", s2);
    s = String::from("lo");
    s.push('l');
    println!("s is: {}", s);

    // concatenation
    let s1 = String::from("Hello, ");
    let s2 = String::from("world!");
    let s3 = s1 + &s2; // s1 has been moved and no longer be used
    println!("s3 is: {}", s3);
    let s1 = String::from("tic");
    let s2 = String::from("tac");
    let s3 = String::from("toe");
    let s = format!("{}-{}-{}", s1, s2, s3);
    println!("s is: {}", s);

    // iteration over individual unicode scalar value
    for c in "Здравствуйте".chars() {
        println!("{}", c);
    }
}
