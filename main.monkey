let a = 10;

let b = [1, 2, 3];

let x = fn(a, b) {
    return a * b;
}

let y = x(a, b[1])

puts(x)

let hasher = {y:10, "foo":20};

puts(hasher["foo"])
/#fsdfsd des fde puts(hasher[y]) #/ 
/#puts(hasher[y], y, a, b[2])#/

let z = x(10, x(5, 5))
puts(z, x(6, 6))
