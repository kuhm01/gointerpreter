let a = 10;

let b = [1, 2, 3];

let x = fn(a, b) {
    return a * b;
}

puts(10)

let y = x(a, b[1])

let hasher = {y:10, "foo":20};

puts(hasher["foo"])

puts(hasher[y], y, a, b[2])

puts(f(1,2))

let f = fn(a, b) {
    return a + b;
}