# 문법
### REPL은 여러 줄을 인식하지 못 함

---

1. __[특이사항](#특이사항 "특이사항")__
2. __[문장의 끝](#문장의-끝 "문장의 끝")__ 
3. __[변수 이름 규칙](#변수이름-규칙 "변수이름 규칙")__
4. __[변수에 값 할당](#변수에-값-할당 "변수에 값 할당")__
5. __[자료형](#자료형 "자료형")__
6. __[함수](#함수 "함수")__
7. __[조건문](#조건문 "조건문")__
8. __[Return문](#Return문 "Return문")__
9. __[연산자](#연산자 "연산자")__
10. __[주석](#주석 "주석")__
11. __[차례성](#차례성 "차례성")__
12. __[Monkey2.0](#monkey20 "Monkey2.0!")__
13. __[for-loop](#for-loop "for-loop")__
14. __[=로 대입](#로-대입 "=로 대입")__
15. __[변수 선언](#변수-선언 "변수 선언")__
16. __[Macro](#macro "Macro")__

---

## 특이사항
Monkey는 정수 객체 비교 시, 포인터 비교를 허용하지않습니다. </br>
따라서 정수 비교가 Boolean 비교보다 느립니다.

Monkey의 비교식에서 0은 정수로서 true로 취급됩니다.

Monkey의 조건식에서 `<consequence>`는 truthy인 경우 평가됩니다.

---
## 문장의 끝
Monkey에서 문장의 끝에는 ; 을 표기합니다.</br>
예시
```Go
let a = 10;
let b = "hello world";
b * 2;
```
반드시 표기하지않아도 무방하나</br>
문법 표준은 ; 을 표기하는 것임을 인지하시기 바랍니다.
</br>
</br>


## 변수이름 규칙
변수이름은 영 대/소문자 및 "_"을 사용할 수 있습니다.

```Go
hello
_hello //_은 처음 또는
hel_lo //중간
hello_ //끝에 사용가능
```
또한 Monkey는 다수개의 내장함수를 기본으로 지원하는 데,</br>
변수이름으로 내장함수의 명칭을 사용할 수 없습니다.</br>
```Go
let puts = 1 // error
```
내장함수에 대한 자세한 사항은 다음 참조.</br>
[내장함수 목록](../main/evaluator/README.md "내장함수 목록")</br>
</br>



## 변수에 값 할당
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

### 재할당
미리 선언되고 값이 할당된 변수에</br>
새로운 값을 할당하면 해당 변수는 새로운 값으로</br>
할당되어집니다. ~~당연한 사실~~
```Go
let a = 10
a // 10
let a = 20
a // 20
```
</br>


## 자료형
Monkey는 다음의 자료형을 제공합니다. </br>
+ INTEGER (정수)
+ FLOAT (실수)
+ STRING (문자열)
+ ARRAY (배열)
+ HASH (해시)
+ BOOLEAN (부울대수)

### 정수와 실수 사용
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
integer, float 함수를 통한 명시적 형변환이 필요합니다.
다음 참고 [링크](https://github.com/kuhm01/gointerpreter/blob/main/evaluator/README.md#integer "내장함수 목록")

### 문자열 사용
```Go
let a = "Hello World"
```

문자열은 다음의 연산을 제공합니다.</br>
STRING + STRING</br>
STRING * INTEGER (경고. INTEGER * STRING 은 불가. 오류를 발생하므로 주의)
```Go
//STRING + STRING
"hello " + "world" // return "hello world"

//STRING * INTEGER
"hello" * 2 // return "hellohello"
"hello" * 0 // return ""
"hello" * a negative number // return error
```

### 배열 사용
```Go
let a = [1, 2.1, "string", [1, 2], {"one": 1}, true, fn(x, y) { return x + y; }]
```
Monkey의 배열에는 공식적으로 지원하는 모든 자료형을 원소로 취급할 수 있습니다.</br>
위의 배열에 대해 다음이 가능합니다.
```Go
integer(a[1]) // return 2
a[2] + " hello" // return "string hello"
a[3][0] // 1
a[4]["one"] // 1
a[5] * false // return false
a[6](1, 2) // return 3
```
단, Monkey에서 배열의 요소에 직접 값을 대입하는 것은 불가능합니다.
```Go
let a = [1, 2]
a[0] = 2 // parsing error
let a[0] = 2 //parsing error
```
__Monkey에서는 index를 통한 변수로의 접근자체가 구현이 안 되어있음을 참고하십시오.__

### 해시 사용
해시는 {키: 값}의 구조를 가집니다.
```Go
let a = {"one": 1, 2: 10.1, true: fn(x, y) { return x + y; }, "two": [1, 2]}
```
해시의 값에는 배열과 마찬가지로 Monkey가 지원하는 모든 자료형이 취급됩니다.</br>
단, 해시의 키에는 정수, 부울대수, 문자열만 가능함을 인지하십시오.</br>
위의 해시에 대해
```Go
a["one"] + 2 // return 3
a[1 > 0](1, 2) // return 3
a["two"][0] // 1
```
위의 강조문과 마찬가지로 index를 통한 변수로의 접근자체는 구현되어 있지않습니다.</br>

### 부울대수
Monkey는 공식적으로 부울대수를 지원합니다. 다음의 값들이 있음
```Go
let a = true
let b = false
```
Monkey에서는 부울대수에 대한 다음의 연산을 지원합니다.</br>
"+" : or, "*" : and
```Go
true + true // return true
true + false // return true
false + true // return true
false + false // return false

true * true // return true
true * false // return false
false * true // return false
false * false // return false
```
</br>

## 함수
Monkey는 당연히 함수를 사용할 수 있습니다.</br>
다음과 같이 선언하여 사용
```Go
let a = fn(x, y) {
  return x + y;
}
```
함수는 fn 키워드를 활용합니다. fn(매개변수) { 표현식 } 의 모양을 가집니다.</br>
함수도 당연히 변수에 할당하여 사용해야합니다.</br>
위의 함수 변수를 다음과 같이 사용가능
```Go
a(1, 2) // return 3
```

Monkey는 재귀함수의 작성을 공식적으로 지원합니다.</br>
따라서 다음의 함수를 작성가능
```Go
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
놀랍게도 순수한 Monkey의 언어로 작성되었습니다.</br>
</br>

또한 고차함수의 작성을 지원합니다.</br>
따라서 다음의 함수를 작성가능
```Go
let twice = fn(x, y) {
  return f(f(x));
}

let addTwo = fn(x) {
  return x + 2;
}

twice(addTwo, 2); // 6
```

함수 내부에서 변수가 선언될 수 있습니다.</br>
단, 해당 변수는 함수 외부에선 존재하지않습니다.
```Go
let a = fn(x, y) {
  let b = 10 + x * y
  return b + 10;
}

b // error
```
</br>

## 조건문
Monkey는 당연히 조건문을 지원합니다.</br>
다음과 같이 사용가능
```Go
if (1 != 10) {
  return 1;
} else {
  return 2;
}
```
아쉽게도 else if, elif는 Monkey에서 지원하지않습니다.</br>
물론 조건문 내부에 조건문을 활용하는 것은 가능합니다.
```Go
if (1 != 10) {
  if (1 != 5) {
    return 1;
  } else {
    return 2;
  }
} else {
  return 3;
}
```
else문은 작성 시, 반드시 if문이 끝나는 } 뒤에 이어 쓰십시오.</br>
또한 else문은 반드시 작성할 필요 없으나, 대신 조건분기에 의해 거짓일 경우</br>
else문이 없다면 Monkey는 null을 반환합니다.</br>
</br>

## Return문
위의 몇몇 코드를 보다보면 으레 return이 있어야 할 자리에 없는 것을 볼 수 있을 겁니다.</br>
그렇습니다. Monkey의 표준문법에서는 return을 생략할수도 있습니다!
다음 참고
```Go
if (1 != 10) {
  1;
} else {
  2;
}
```
위 코드에서 if, else문에서 각각 1과 2를 반환한다는 의미로 앞에 return을 작성하는 것이</br>
당연하게 보입니다. 그러나 Monkey에서는 저런 경우 묵시적으로 return이 있다고 인식하므로</br>
return을 작성하지않아도 값을 반환합니다.</br>
</br>

## 연산자
Monkey는 당연하게도 연산자를 제공합니다.</br>

+ 사칙연산
"+", "-", "*", "/"

+ 대소비교
">", "<"

+ 논리비교
"==", "!="</br>
</br>

## 주석
Monkey는 놀랍게도 주석도 지원합니다.</br>
주석은 반드시 주석문의 앞에 "/#"와 "#/"을 작성하여 사용합니다.</br>
다음 예시
```Go
let a = fn(x, y) { return x + y; }
/# that is standard adding function #/
```
주석문은 Monkey의 여러 과정 중</br>
처음 어휘분석단계에서부터 인식되지않습니다.</br>
</br>

## 차례성
Monkey는 사용자가 입력하는대로 위에서부터 아래로</br>
코드를 차례대로 인식합니다.</br>
따라서 다음의 경우, 오류발생
```Go
add(1, 2)

let add = fn(x, y) {
  return x + y;
}
```
위 코드에서 식별자 add를 호출할 때, Monkey에 add는 없는 것과 같습니다.</br>
변수 add를 선언하고 값을 할당하는 것이 호출문 아래에 있기때문입니다.</br>
따라서 위의 코드는 오류발생.</br>
</br>

# Monkey2.0
Monkey에 획기적인 변화가 일어났습니다!</br>
실습자는 꽤나 머리를 굴려 Monkey에 문법을 새로이 추가했고</br>
이는 Monkey를 더욱 빛나게 할 것입니다. 아마도?</br>
아래에서 업데이트된 문법을 소개하겠습니다.</br>

## for-loop
Monkey가 드디어 for-loop를 지원합니다!</br>
다음과 같이 사용</br>
```Go
for i := range(5) {
  puts(i)
}
```
핵심은 for와 "{" 중간에 loop를 통제할 표현식이 필요하다는 점입니다!</br>
기본사용법은 위와같이 <변수> := range(반복할 횟수) 로 표현합니다.</br>
다만 range를 사용하지않고 미리 정의된 배열을 사용할 수도 있습니다!</br>
다음과 같음</br>
```Go
let a = [1, 2.1, "string"]
for i := a {
  puts(i)
}
/*
print
1
2.1
string
*/
```

## :=
Monkey에 새롭게 := 연산자가 추가되었습니다! 눈치 채셨겠지만 위의 for-loop를 위한 연산자이기도 합니다.</br>
Monkey는 이제 := 연산자를 통해서 let 없이도 변수를 선언하고 값을 할당할 수 있습니다.</br>
let 과 똑같이 기능하므로 할당할 수 있는 값은 전부 할당할 수 있습니다.
다음과 같이 사용</br>
```Go
a := 1 
puts(a) // print 1
```

## =로 대입
Monkey는 이제 =을 통해서 간단하게 변수에 값을 할당할 수 있습니다.</br>
다음과 같이 사용</br>
```Go
a = 12 // error
let b = 12
c := 12
b = 13
c = 13
```
위의 코드에 집중해주십시오. 처음 a변수에 값을 할당하는 것은 error인데</br>
그 이유는 식별자 a가 처음 let 이나 := 연산자를 통해서 선언되지않았기 때문입니다.</br>
즉, =를 통한 대입은 미리 선언되고 값이 할당된 변수에 대해 사용할 수 있습니다.</br>

### 추가로 중요한 Point
Monkey는 변수를 선언하고 값을 할당할 때에 그 것이 저장될 환경을 내부적으로 구현합니다.</br>
즉, Monkey에서 선언된 여러 변수는 환경에 저장이됩니다. 여느 언어의 block범위와 유사합니다.</br>
__[함수](#함수 "함수")__ 문법에 이와 비슷한 내용이 있습니다.</br>
하여튼 Monkey에는 환경이 있는 데, let 과 := 연산자를 통한 값의 대입은 환경을 엄밀히 논하지않습니다.</br>
반면 =을 통한 대입은 환경을 엄밀히 논합니다.</br>
즉, 이 말은 =을 통한 대입은 다른 대입과 달리 변수가 어느 환경에 위치해있는 지 정확히 검토하고</br>
해당 환경에 값을 대입하고자 한다는 이야기입니다.</br>
다음의 예시를 참고</br>
```Go
let a = 1
for i := range(5) {
  let a = a + i
}
puts(a) // print 1
```
위와 같은 결과가 나오는 이유는 let 은(그리고 := 연산자는) 변수가 위치한 환경을 엄밀히 논하지않기때문에</br>
for-loop 내부의 환경에 값을 대입하게되는 것입니다.</br>
따라서 for-loop가 끝난 후의 변수 a는 for-loop 외부의 변수 a이므로 값이 그대로인 것입니다.</br>
</br>

반면, 다음의 예시에서는
```Go
let a = 1
for i := range(5) {
  a = a + i
}
puts(a) // print 11
```
무언가 우리가 원한 결과가 나왔다고 할 수 있습니다. 이 것은 =을 통한 대입은 변수가 위치한 환경을</br>
엄밀히 논하기때문입니다.</br>
</br>

즉, 대략 다음과 같이 이해할 수 있습니다.</br>
+ let, := 을 통한 대입 : 변수의 신규 선언 및 값 할당
+ =을 통한 대입 : 기존 변수에 값 할당
</br>
이러한 차이를 염두에 두고 신중히 사용하기 바랍니다.</br>
</br>

## 변수 선언
어라? 변수 선언은 위에 있는 거 아닌가요?</br>
물론 맞습니다. 그런데 Monkey는 이제 변수의 선언만이 가능해집니다!</br>
var `<ident>` objectType 과 같은 모양으로 선언하는 것으로 Monkey는 이제 변수만을 선언할 수 있게되었습니다!</br>
다음과 같이 사용
```Go
var value string // string("")으로 초기화
var fv float // float(0.0)으로 초기화
var iv integer // Integer(0)으로 초기화
var bv boolean // Boolean(true)으로 초기화
```
var 를 사용함으로써 Monkey는 이제 변수의 선언만이 가능해지고 해당 자료형의 초깃값으로 자동 초기화되는 기술을 누릴 수 있습니다!</br>
var 를 사용한 변수의 선언은 뒤에 자료형을 명시해주어야합니다.</br>
명시 가능한 자료형은 다음과 같음.</br>

+ string (문자열 : ""로 초기화)
+ integer (정수 : 0으로 초기화)
+ float (실수 : 0.0으로 초기화)
+ boolean (불린대수 : true로 초기화)

### 주의
위의 for-loop 설명 중 [추가로 중요한 Point](#추가로-중요한-Point "참고")와 같이</br>
var 도 let 이나 := 연산자와 같은 "변수를 신규 선언"하는 방식이기때문에</br>
변수의 환경을 엄밀하게 조사하지않습니다. 따라서 해당 문단에서 기술한 상황이 발생함을 인지하십시오.</br>

## Macro
잃어버린 챕터 5장. Macro시스템을 구현하였습니다.</br>
상세한 내용은 아래 링크 확인</br>
https://interpreterbook.com/lost/

### 비고
Monkey 내부에서 Macro를 선언하고 변수에 할당할 때에는</br>
무조건 let 을 사용해야합니다. 기타 대입 방법은 사용불가</br>
</br>


##### 이 외에도 실습자가 Monkey에 class 또는 구조체 등의 다양한 무언가를 추가하기 위해 열심히 머리를 굴리고 있습니다.
