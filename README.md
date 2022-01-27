# GOINTERPRETER
## Interpreter written in Golang
본 내용은 책의 내용을 실습한 프로그램입니다.

---
목차
1. [책정보](#책정보 "책정보")
2. [구현 정보](#구현-정보 "구현 정보")
3. [Monkey](#Monkey "몽키")
4. [Monkey 실행](#실행해보기 "Monkey 실행해보기")



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

### 잘 작성되어있나요?
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
