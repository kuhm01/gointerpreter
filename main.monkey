let a = 10;

let b = [1, 2, 3];

let x = fn(a, b) {
    return a * b;
}

let y = x(a, b[1])

let hasher = {y:10, "foo":20};

puts(hasher[y], y, a, b[2])