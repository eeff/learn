fn main() {
    let number = 3;

    // condition must be type bool
    if number < 5 {
        println!("condition was true");
    } else {
        println!("condition was false");
    }

    // if is an expression, each arm must have same type
    let condition = true;
    let number = if condition { 5 } else { 6 };
    println!("The value of number is: {}", number);
}
