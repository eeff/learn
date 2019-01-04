fn main() {
    // creation
    let v: Vec<i32> = Vec::new();
    let v = vec![1, 2, 3];
    let mut v = Vec::new();

    // update
    v.push(5);
    v.push(6);
    v.push(7);
    v.push(8);

    // element access
    let third: &i32 = &v[2];
    println!("The third element is {}", third);
    match v.get(2) {
        Some(third) => println!("The third element is {}", third),
        None => println!("There is no third element"),
    }

    // iteration
    for i in &v {
        println!("{}", i);
    }
    for i in &mut v {
        *i += 50;
    }
}
