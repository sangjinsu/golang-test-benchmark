# Golang Test and Benchmark

- Turker 의 Go 언어 프로그래밍

- [Udemy] [Learn How To Code: Google's Go (golang) Programming Language](https://www.udemy.com/course/learn-how-to-code/)

  ## 정리

  - `[file name]_test.go` 으로 끝나는 파일을 생성해야 한다

  - testing 패키지를 사용한다

  - `func TestXxx(t *testing.T)` 형식으로 함수를 작성한다

  -  `func TestXxx(b *testing.b)` 형식으로 함수를 작성한다

  - 실행 명령어

    ```bash
    $ go test
    $ go -run test [test function name] // 지정 실행
    $ go test -bench . // 벤치마크 테스트
    $ go test -bench [test function name] // 지정 실행
    ```

  ## 외부 테스팅 패키지 사용

  - `strectchr/testify` 패키지

    ```bash
      $ go get -u github.com/stretchr/testify
    ```
  
    패키지 사용
  
    ```go
    func TestSquare1(t *testing.T) {
    	assertions := assert.New(t)
    	assertions.Equal(81, square(9), "square(9) should be 81")
    }
    ```
  
    패키지 미사용
  
    ```go
    func TestSquare3(t *testing.T) {
       result := square(3)
       if result != 9 {
          t.Errorf("square(9) should be 9 but returns %d", result)
       }
    }
    ```
  
  - assert 패키지
    - `Equal(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool` : 두 값이 같을 경우 테스트를 통과한다
    - `Greater(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) bool` : e2가 e1보다 크면 테스트를 통과한다
    - `Len(object interface{}, length int, msgAndArgs ...interface{}) bool` : object 항목 개수가 length 가 일치해야 테스트를 통과한다
    - `NotNil(object interface{}, msgAndArgs ...interface{}) bool` : object 가 nil이 아니면 테스트를 통과한다
    - `NotEqual(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool`: 두 값이 다르면 테스트를 통과한다