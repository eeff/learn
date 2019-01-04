// definition
struct User {
    username: String,
    email: String,
    sign_in_count: u64,
    active: bool,
}

#[allow(unused_variables)]
fn main() {
    // instantiation
    let user = User {
        email: String::from("someone@example.com"),
        username: String::from("someusername123"),
        active: true,
        sign_in_count: 1,
    };

    let mut user1 = build_user(
        String::from("someone@example.com"),
        String::from("someusername123"),
    );
    // the entire instance have to be mutable
    user1.email = String::from("anotheremail@example.com");
    println!(
        "username: {}, email: {}, active: {}, sign_in_count: {}",
        user1.username, user1.email, user1.active, user1.sign_in_count
    );

    // struct upate syntax
    let user2 = User {
        email: String::from("another@example.com"),
        username: String::from("anotherusername567"),
        ..user1 // other fields same as user1
    };
    println!(
        "username: {}, email: {}, active: {}, sign_in_count: {}",
        user2.username, user2.email, user2.active, user2.sign_in_count
    );

    // tuple struct, destructuring, access with dot index
    // each struct define its own type
    struct Color(i32, i32, i32);
    struct Point(i32, i32, i32);
    let black = Color(0, 0, 0);
    let Color(x, y, z) = black;
    println!("({}, {}, {})", x, y, z);
    let origin = Point(0, 0, 0);
    let Point(x, y, z) = origin;
    println!("({}, {}, {})", x, y, z);
}

fn build_user(email: String, username: String) -> User {
    // field init shorthand
    User {
        email,
        username,
        active: true,
        sign_in_count: 1,
    }
}
