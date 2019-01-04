use std::collections::HashMap;

#[allow(unused_variables)]
fn main() {
    let mut scores = HashMap::new();
    scores.insert(String::from("Blue"), 10);
    scores.insert(String::from("Yellow"), 50);

    let teams = vec![String::from("Blue"), String::from("Yellow")];
    let initial_scores = vec![10, 50];
    let scores: HashMap<_, _> = teams.iter().zip(initial_scores.iter()).collect();
    // key and value are move by default

    // get
    let team_name = String::from("Blue");
    let score = scores.get(&team_name);
    if let Some(i) = score {
        println!("Score is {}", i);
    }

    // iteration
    for (key, value) in &scores {
        println!("{}: {}", key, value);
    }

    // update
    // overwriting
    let mut scores = HashMap::new();
    scores.insert(String::from("Blue"), 10);
    scores.insert(String::from("Yellow"), 25);
    println!("{:?}", scores);
    // only inserting a value if the key has no value
    let mut scores = HashMap::new();
    scores.insert(String::from("Blue"), 10);
    scores.entry(String::from("Yellow")).or_insert(50);
    scores.entry(String::from("Blue")).or_insert(50);
    println!("{:?}", scores);
    //
    let text = "hello world wonderful world";
    let mut map = HashMap::new();
    for word in text.split_whitespace() {
        let count = map.entry(word).or_insert(0);
        *count += 1;
    }
    println!("{:?}", map);
}
