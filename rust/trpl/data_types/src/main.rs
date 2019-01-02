fn main() {
    // Scalar
    // integer
    let _x: u8 = b'A';
    let _x: u16 = 0xFFFF;
    let _x = 1000; //default i32
    let _x: i32 = 1000;
    let _x: i64 = 100__000;
    let _x: i128 = 1_000_000;
    let _x: isize = 0o777;
    let _x: usize = 0b1111_1111;

    // floating point
    let _x: f32 = 1.0;
    let _x = 3.0; // default f64
    let _x: f64 = 2.0;

    // boolean
    let _x = false;
    let _x: bool = true;

    // character
    let _x = 'A';
    let _x: char = '😻';

    // compound
    // tuple
    let tup: (i32, f64, u8) = (500, 6.4, 1);
    let (x, y, z) = tup;
    println!("The value of tup is: ({}, {}, {})", x, y, z);
    println!("The value of tup is: ({}, {}, {})", tup.0, tup.1, tup.2);

    // array
    let _x: [i32; 5] = [1, 2, 3, 4, 5];
    let a = [1, 2, 3];
    println!("The value of a is: [{}, {}, {}]", a[0], a[1], a[2]);
    //let bad = a[100];
    //println!("The value of bad is: {}", bad);
}
