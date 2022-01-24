# 문법
### REPL은 여러 줄을 인식하지 못 함

---

1. __[특이사항](#특이사항 "특이사항")__
2. __[변수 이름 규칙](#변수이름규칙 "변수이름규칙")__
3. __[변수에 값 할당](#변수에값할당 "변수에 값 할당")__
4. __[자료형](#자료형 "자료형")__

---

## 특이사항
Monkey는 정수 객체 비교 시, 포인터 비교를 허용하지않습니다. </br>
따라서 정수 비교가 Boolean 비교보다 느립니다.

Monkey의 비교식에서 0은 정수로서 true로 취급됩니다.

Monkey의 조건식에서 `<consequence>`는 truthy인 경우 평가됩니다.

---

## 변수이름규칙
변수이름은 영 대/소문자 및 "_"을 사용할 수 있습니다.

```Go
hello
_hello //_은 처음 또는
hel_lo //중간
hello_ //끝에 사용가능
```
</br>

## 변수에값할당
### 선언
변수를 선언할 때는 변수이름 앞에 let 을 사용합니다.
```Go
let hello
let star
let monkey
```

### 할당
변수를 let 을 사용해 선언하면 바로 값을 할당해야합니다.
```Go
let hello = "hello"
let a = 1
let result = 10 + (2 * 10) 

let a //이렇게 사용하는 경우, 구문분석 오류발생
```
</br>



## 자료형
Monkey는 다음의 자료형을 제공합니다. </br>
+ INTEGER (정수)
+ FLOAT (실수)
+ STRING (문자열)
+ ARRAY (배열)
+ HASH (해시)

정수와 실수 사용
```Go
let a = 10; let b = 20;
a + b // 30

let c = 30.1; let d = 10.1;
c + d // 40.2
```

Monkey는 정수와 실수의 묵시적 형변환을 지원하지않습니다.</br>
다음의 경우 오류발생
```Go
let a = 10 // integer
let b = 10.1 // float

a + b // return error
```

