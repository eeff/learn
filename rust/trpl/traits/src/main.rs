use std::fmt::Debug;
use std::fmt::Display;

// definition
pub trait Summary {
    fn summarize(&self) -> String {
        // default implementation
        // default implementations can call method in the same trait
        format!("(Read more from {}...)", self.summarize_author())
    }
    fn summarize_author(&self) -> String;
}

pub struct NewsArticle {
    pub headline: String,
    pub location: String,
    pub author: String,
    pub content: String,
}

// we can implement a trait on a type only if
// either the trait or the type is local to our crate
impl Summary for NewsArticle {
    // It's NOT possible to call the default implementation
    // from an overriding implementation of that same method
    fn summarize(&self) -> String {
        format!("{}, by {} ({})", self.headline, self.author, self.location)
    }
    fn summarize_author(&self) -> String {
        self.author.clone()
    }
}

pub struct Tweet {
    pub username: String,
    pub content: String,
    pub reply: bool,
    pub retweet: bool,
}

impl Summary for Tweet {
    fn summarize(&self) -> String {
        format!("{}: {}", self.username, self.content)
    }
    fn summarize_author(&self) -> String {
        format!("@{}", self.username)
    }
}

// trait bounds
pub fn notify<T: Summary + Display>(item: T) {
    // syntanx sugar: pub fn notify(item: impl Summary + Display)
    println!("Breaking news! {}", item.summarize());
}

#[allow(dead_code)]
#[allow(unused_variables)]
fn some_function<T, U>(t: T, u: U)
where
    T: Display + Clone,
    U: Clone + Debug,
{
    // where clauses for clearer code
}

#[allow(dead_code)]
fn return_summarizable() -> impl Summary {
    // only works for single type returning
    Tweet {
        username: String::from("horse_ebooks"),
        content: String::from("of course, as you probably already know, people"),
        reply: false,
        retweet: false,
    }
}

fn main() {
    let tweet = Tweet {
        username: String::from("horse_ebooks"),
        content: String::from("of course, as you probably already know, people"),
        reply: false,
        retweet: false,
    };

    println!("1 new tweet: {}", tweet.summarize());

    let article = NewsArticle {
        headline: String::from("Penguins win the Stanley Cup Championship!"),
        location: String::from("Pittsburgh, PA, USA"),
        author: String::from("Iceburgh"),
        content: String::from(
            "The Pittsburgh Penguins once again are the best hockey team in the NHL.",
        ),
    };

    println!("New article available! {}", article.summarize());
}
