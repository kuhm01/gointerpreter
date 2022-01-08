# gointerpreter
## Interpreter written in Golang
본 내용은 책의 내용을 실습한 프로그램입니다.

## 책정보 
제목 : 밑바닥부터 만드는 INTERPRETER in Go

ISBN : 978-89-6626-316-5

정보 : http://www.riss.kr/search/detail/DetailView.do?p_mat_type=75f99de66db18cf6&control_no=bf9574aa039969c6ffe0bdc3ef48d419

책에서 제공되는 코드 링크 : https://interpreterbook.com/waiig_code_1.3.zip

## 정보
사용언어 : Go (Version : 1.17.5)

### Commit name rule
n장 <챕터명> - <부챕터명> : 완료

## 구현하는 언어
명칭 : Monkey

### Monkey
Monkey는 REPL 환경에서 작동합니다.

## 실행해보기
test : go test ./(package name that you want to test)

main : go run main.go

## Monkey 언어 명세
Monkey는 정수 객체 비교 시, 포인터 비교를 허용하지않습니다. 따라서 정수 비교가 Boolean 비교보다 느립니다.

Monkey의 비교식에서 0은 정수로서 true로 취급됩니다.

Monkey의 조건식에서 `<consequence>`는 truthy인 경우 평가됩니다.

## Monkey 사용해보기
이름에 값 바인딩하기
```Monkey
let age = 1;
let name = "Monkey";
let result = 10 * (20 / 2);
```

정수 배열 바인딩하기
```Monkey
let myArray = [1, 2, 3, 4, 5];
```

해시타입. 각 값은 해당 키와 연결
```Monkey
let thorsten = {"name": "Thorsten", "age"L 28};
```

배열과 해시 요소는 인덱스 표현식으로 접근
```Monkey
myArray[0] // = 1
thorsten["name"] // = "Thorsten"
```

let 문은 함수 바인딩도 가능
```Monkey
let add = fn(a, b) { return a + b; };
let add = fn(a, b) { a + b; };  // = return a + b. 암묵적인 return값 허용

add(1, 2); // = 3. 함수 호출
```


재귀호출
```Monkey
let fibonacci = fn(x) {
  if (x == 0) {
    0
  } else {
    if (x == 1) {
      1
    } else {
      fibonacci(x - 1) + fibonacci(x - 2);
    }
  }
};
```

고차 함수
```Monkey
let twice = fn(x, y) {
  return f(f(x));
};

let addTwo = fn(x) {
  return x + 2;
};

twice(addTwo, 2); // = 6
```

출처</br>
토르슈텐 발, 밑바닥부터 만드는 인터프리터 INTERPRETER in go (n.p.: 도서출판 인사이트, n.d.), 13-15.
