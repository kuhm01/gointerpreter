# GOINTERPRETER
## Interpreter written in Golang
본 내용은 책의 내용을 실습한 프로그램입니다.

---
목차
1. [책정보](#책정보 "책정보")
2. [구현 정보](#구현-정보 "구현 정보")
3. [Monkey](#Monkey "몽키")
4. [Monkey 실행](#실행해보기 "Monkey 실행해보기")
5. [Monkey 2.0](#Monkey2.0 "Monkey 2.0!")



---
## 책정보 
제목 : 밑바닥부터 만드는 INTERPRETER in Go

ISBN : 978-89-6626-316-5

정보 : http://www.riss.kr/link?id=M15913813

책에서 제공되는 코드 링크 : https://interpreterbook.com/waiig_code_1.3.zip

---
## 구현 정보
구현에 사용한 언어 : Go (Version : 1.17.5)

### 커밋 명칭의 규칙
n장 <챕터명> - <부챕터명> : 완료

---
## Monkey
이 책을 통해 우리가 구현하는 언어는 Monkey입니다.</br>
다음에서 Monkey에 대한 정보를 제공

+ 문법
+ 내장함수
+ Monkey 실행옵션
+ REPL 지원

### 문법
Monkey의 문법에 대해 소개합니다.</br>
Monkey에는 책에 정의된 문법이 있고 이를 표준문법이라 칭하겠습니다.</br>
Monkey에는 표준문법 외에 실습자가 추가한 기타 문법들이 존재합니다.</br>
[문법 열람](../main/GRAMMAR_README.md "Monkey의 문법 열람")
</br>

### 내장함수
Monkey는 사용자의 편의를 위해 다양한 함수를 미리 구현해두었습니다.</br>
[내장함수 목록](../main/evaluator/README.md "Monkey가 지원하는 기본 내장함수")
</br>

### Monkey 실행옵션
Monkey는 다양한 실행옵션을 제공합니다.</br>
옵션없이 기본모드로 실행할 경우 Monkey는 REPL환경을 제공합니다.</br>
</br>
옵션 내역
* -c, start : interpreting monkey file(monkey파일을 읽어 실행합니다.)
* -h, help : printing option list

### REPL 지원
Monkey는 옵션없이 실행할 경우, REPL환경을 기본으로 제공합니다.</br>

---
## 실행해보기
### 우선 Monkey 컴파일하기
Go언어로 작성된 Monkey를 컴파일하세요!</br>
다음 과정으로 컴파일
```Go
go build
```
Monkey는 Go언어로 완전히 작성되었기때문에 특별한 컴파일 옵션이 필요하지않습니다.</br>
또는 컴파일하지않고 다음 과정으로 Monkey를 실행할 수 있습니다.
```Go
go run main.go
```
물론 위의 두 과정 전부 Monkey의 main.go 파일이 있는 주 디렉토리에서 실행하십시오.</br>

### REPL환경에서 작성
어려울 것이 있습니까! REPL 환경에서 자유롭게 Monkey를 활용하십시오.</br>
다만, Monkey의 REPL은 여러 줄을 인식하지 못 합니다.</br>
즉, 엔터키를 누르는 순간 문장을 인식하고 실행하므로 모든 문장은 한 줄에 기술하십시오.</br>

### Monkey 파일을 작성
REPL환경이 싫다면 Monkey파일을 작성하여 읽어 실행하도록 할 수 있습니다!</br>
Monkey파일은 반드시 확장자가 monkey여야합니다. (ex. main.monkey)</br>
또한 당신이 실행할 Monkey파일은 반드시 Monkey실행파일이 있는 주 디렉토리에 위치해야 합니다.</br>

### 잘 작동하나요?
표준문법은 그렇습니다! 그리고 실습자가 추가한 기타 문법도 물론 그럴 것입니다.</br>
본 레포지토리에는 책에서 기술한 표준 test파일들이 있습니다.</br>
다음 과정으로 test 파일 실행
```Window
go test ./(package name that you are want to test)
```
테스트할 수 있는 패키지들은 다음과 같습니다.
```Go
monkey/lexer, monkey/parser, monkey/ast, monkey/object, monkey/evaluator
```

### exit(exit)
Monkey의 REPL에서 exit(exit)를 입력해보십시오.</br>
그러면 원숭이(Monkey)가 도망가고 Monkey는 종료됩니다!</br>
이 것은 언어의 함수가 아니라 프로그램 자체가 인식하는 명령문입니다.</br>
사실 Monkey를 종료하는 효율적인 방법은 Ctrl + Z 이므로 취향대로 사용하십시오.</br>

---
## Monkey2.0
실습자가 머리를 오지게 굴리다보니 Monkey에겐 꽤나 획기적인</br>
업데이트가 이루어졌습니다. 지금부터 이 것을 Monkey2.0이라고 부를 것입니다!</br>
일단 다음이 추가되었습니다.
+ for-loop
+ := 연산자 추가
+ = 기호를 통한 대입가능
+ var
+ 새로운 여럿 내장함수

이 외에도 계속 많은 것이 추가될 것입니다!</br>
다음에서 상세한 사항 확인</br> 
[새로운 문법](../main/GRAMMAR_README.md#Monkey2.0 "Monkey2.0 문법")</br>
[새로운 함수](../main/evaluator/README.md#Monkey2.0 "Monkey2.0 함수")</br>

### for-loop
Monkey가 드디어 대부분의 언어에 있는 for-loop를 지원합니다!</br>

### := 
':=' 연산자가 추가되었습니다!</br>
따라서 지금 이 순간부터 Monkey는 let 없이도</br>
변수의 선언과 값의 할당을 이루어낼 수 있습니다!</br>

### = 기호를 통한 대입가능
Monkey가 미리 선언된 변수에 한해서</br>
이제 꼭 let문이나 위의 := 없이 평범하게 = 을 통한</br>
대입으로도 값의 할당이 가능해집니다! 답답했는 데 잘 됐어요!</br>

### var
이제 Monkey는 var 키워드를 통한 변수의 선언만이 가능해졌습니다.</br>
그러면 let 이랑 := 은 필요없는 것 아닌가요?</br>
물론 그럴 수 있는 데! 그만큼 Monkey는 훨씬 더 많은 선택지를 자유롭게</br>
선택할 수 있게되었다고 할 수 있는 것이죠! 큰 발전입니다</br>

### 새로운 내장함수들
Monkey에 많은 내장함수들이 새롭게 추가되었습니다!</br>
이 함수들이 여러분들이 Monkey를 이용함에 있어 사소한 도움이 되리라 자부합니다</br>

### Monkey2.0은 잘 작동하나요?
대대적으로 큰 변경이 이루어진만큼</br>
어쩌면 표준문법에서조차 버그가 발생할 수 있습니다. 버그를 발견했다면 빠르게 알려주세요!</br>
하지만 표면적인 test는 대게 통과하였기때문에 문제없이 작동할 것입니다!</br>
실습자가 디버깅한다고 parser_test.go나 evaluator_test.go파일에 의미없는</br>
test case가 추가되었는 데, 관심있으면 다양한 경우로 디버깅해보세요!</br>