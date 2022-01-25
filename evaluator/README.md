# 내장함수 목록

1. __[len](#len "len 함수")__
2. __[first](#first "first 함수")__
3. __[last](#last "last 함수")__
4. __[rest](#rest "rest 함수")__
5. __[push](#push "push 함수")__
6. __[puts](#puts "puts 함수")__
7. ___[pop](#pop "pop 함수")___
8. ___[integer](#integer "integer 함수")___
9. ___[float](#float "float 함수")___
10. ___[range](#range "range 함수")___
11. ___[Itoa](#Itoa "Itoa 함수")___
12. ___[typeOf](#typeOf "typeOf 함수")___

---
_이탤릭체 함수_ 는 책에 없는</br>
실습자가 따로 추가한 함수입니다.</br>

## len
len(ARRAY) : return length of Array</br>
배열의 길이를 반환합니다.
```Go
let a = [1, 2, 3]
len(a) //return 3
```

## first
first(ARRAY) : return first element of Array</br>
배열의 첫번째 원소를 반환합니다.
```Go
let a = [1, 2, 3]
first(a) //return 1
```

## last
last(ARRAY) : return last element of Array</br>
배열의 마지막 원소를 반환합니다.
```Go
let a = [1, 2, 3]
last(a) //return 3
```

## rest
rest(ARRAY) : return new Array except for first element</br>
첫번째 원소를 제외한 배열을 반환합니다.
```Go
let a = [1, 2, 3]
rest(a) //return [2, 3]
```

## push
push(ARRAY, var) : your variable push in Array</br>
배열의 마지막에 변수를 추가합니다.
```Go
let a = [1, 2, 3]
push(a, 4) //return [1, 2, 3, 4]
```

push(ARRAY, var, INTEGER) : your variable push in Array at INTEGER</br>
원하는 위치에 변수를 추가합니다.
```Go
let a = [1, 2, 3]
push(a, 10, 0) //return [10, 1, 2, 3]
push(a, 10, 1) //return [1, 10, 2, 3]
push(a, 10, 2) //return [1, 2, 10, 3]

push(a, 10, 3) //return error
```

## puts
puts(args...) : Standard Print Function</br>
표준 출력함수. 마지막에 null을 출력.
```Go
let a = [1, 2, 3]
let b = {"one": 1, "two": 2}
let c = fn(x, y) { return x + y }

puts(a) 
/*print

[1, 2, 3]
null

*/

puts(b)
/*print

{one: 1, two: 2}
null

*/

puts(c)
/*

fn(x, y) {
return (x + y);
}
null

*/

let x = "hello world"
let y = 12
puts(x, y)
/*print

hello world
12
null
```

## _pop_
pop(ARRAY, INTEGER) : return new Array except for array[index]</br>
배열에서 입력받은 위치의 원소를 제거합니다.
```Go
let a = [1, 2, 3]
pop(a, 2) //return [1, 2]
pop(a, 1) //return [1, 3]
pop(a, -1) //return error

let b = []
pop(b, "other (integer)value") //return error
```

## _integer_
integer(FLOAT) : wrapper to integer from float</br>
실수를 정수로 변환.
```Go
let b = 10.1 //float
integer(b) //wrap to integer : 10
```

## _float_
float(INTEGER) : wrapper to float from integer</br>
정수를 실수로 변환.
```Go
let a = 10 //integer
float(a) //wrap to float : 10.000000
```

## _range_
range(INTEGER) : return Array. 0 to var. var is Integer</br>
0부터 1씩 증가하는 오름차순 정수배열을 입력받은 길이만큼 반환합니다.
```Go
range(5) //return [0, 1, 2, 3, 4]
range(Other types than Integer) //return error
```

## _Itoa_
Itoa(INTEGER) : return string from integer</br>
입력받은 정수를 문자열로 변환합니다.
```Go
Itoa(10) //return "10" (string type)
```

## _typeOf_
typeOf(var) : return type of value</br>
입력받은 변수의 타입을 출력합니다.
```Go
let a = 1 //integer type
typeOf(a) //return INTEGER

let b = "hello monkey" //string type
typeOf(b) //return STRING
//typeOf function is return string
```
