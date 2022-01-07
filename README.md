# gointerpreter
## Interpreter written in Golang
본 내용은 책의 내용을 실습한 프로그램입니다.

## 책정보 
제목 : (밑바닥부터 만드는) 인터프리터 in Go

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
