# Golang Test and Benchmark

- Turker 의 Go 언어 프로그래밍

- [Udemy] [Learn How To Code: Google's Go (golang) Programming Language](https://www.udemy.com/course/learn-how-to-code/)

  ## 정리

    - [test naming example](https://go.dev/blog/examples)

    - `[file name]_test.go` 으로 끝나는 파일을 생성해야 한다

    - testing 패키지를 사용한다

    - `func TestXxx(t *testing.T)` 형식으로 함수를 작성한다

    - `func TestXxx(b *testing.b)` 형식으로 함수를 작성한다

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
      func TestSquare(t *testing.T) {
          assertions := assert.New(t)
          assertions.Equal(81, square(9), "square(9) should be 81")
      }
      ```

      패키지 미사용

      ```go
      func TestSquare_third(t *testing.T) {
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

  ## DFS 깊이 우선 탐색 벤치 마크

    - 반복문

      ```go
      func dfs1(v int, visited []bool, graph [][]int) {
          stack := []int{v}
          for len(stack) > 0 {
              node := stack[len(stack)-1]
              stack = stack[:len(stack)-1]
              visited[node] = true
              // fmt.Println(node)
              for _, nd := range graph[node] {
                  if visited[nd] == false {
                      stack = append(stack, nd)
                  }
              }
          }
      }
      ```

    - 재귀문

      ```go
      func dfs2(v int, visited []bool, graph [][]int) {
          visited[v] = true
          // fmt.Println(v)
          for _, node := range graph[v] {
              if visited[node] == false {
                  dfs2(node, visited, graph)
              }
          }
      }
      ```

    - 벤치마크

      ```go
      func BenchmarkDFS(b *testing.B) {
          for i := 0; i < b.N; i++ {
              visited := make([]bool, n+1)
              dfs1(1, visited, graph)
          }
      }
      
      func BenchmarkDFSRecursive(b *testing.B) {
          for i := 0; i < b.N; i++ {
              visited := make([]bool, n+1)
              dfs2(1, visited, graph)
          }
      }
      
      // cpu: Intel(R) Core(TM) i5-8265U CPU @ 1.60GHz
      // BenchmarkDFS
      // BenchmarkDFS-8            	 3612478	       277.7 ns/op
      // BenchmarkDFSRecursive
      //BenchmarkDFSRecursive-8   	14021049	       107.1 ns/op
      ```

  ## Table Testing

  ```go
  func TestSquare_fourth(t *testing.T) {
  
  	type test struct {
  		data   int
  		answer int
  	}
  
  	tests := []test{
  		{5, 25},
  		{7, 49},
  		{8, 64},
  		{10, 100},
  		{2, 4},
  	}
  
  	assertions := assert.New(t)
  	for _, test := range tests {
  		assertions.Equalf(test.answer, square(test.data),
  			"square(%d) should be %d", test.data, test.answer)
  	}
  }
  ```

  ## 테스트가 여러개인 경우

  ```go
  func ExampleReverse()
  func ExampleReverse_second()
  func ExampleReverse_third()
  ```

  ## go vet

    - `vet` 은 go 소스 코드를 검사하고 형식 문자열과 일치하지 않는 인수를 가진 Printf 호출과 같은 의심스러운 구조를 보고한다

    - 모든 보고가 진짜 문제라는 것을 보장하지는 않지만 컴파일러에 의해 잡히지 않는 오류를 발견한다

      ```bash
      $ go vet
      $ go vet my/project/..
      ```

  ## go help testflag

    - [testflag](https://pkg.go.dev/cmd/go/internal/test)

        - `-bench regexp`
          > Run only those benchmarks matching a regular expression. By default, no benchmarks are run. To run all benchmarks, use '-bench .' or '-bench=.'

        - `-benchtime t`

          > Run enough iterations of each benchmark to take t, specified
          > as a time.Duration (for example, -benchtime 1h30s).
          > The default is 1 second (1s).

        - `-benchmem`

          >
          > Print memory allocation statistics for benchmarks.

>

## Coverage

- `-cover`

  > Enable coverage analysis.
  >
  > Note that because coverage works by annotating the source
  > code before compilation, compilation and test failures with
  > coverage enabled may report line numbers that don't correspond
  > to the original sources.

- `-coverprofile cover.out`

  >
  > 	    Write a coverage profile to the file after all tests have passed.
  > 	    Sets -cover.
>

- 브라우저에서 `cover.out` 보기

  ```bash
  $ go tool cover -html=cover.out
  ```

- `go tool cover` 사용법

  ```bash
  Usage of 'go tool cover':
  Given a coverage profile produced by 'go test':
          go test -coverprofile=c.out
  
  Open a web browser displaying annotated source code:
          go tool cover -html=c.out
  
  Write out an HTML file instead of launching a web browser:
          go tool cover -html=c.out -o coverage.html
  
  Display coverage percentages to stdout for each function:
          go tool cover -func=c.out
  
  Finally, to generate modified source code with coverage annotations
  (what go test -cover does):
          go tool cover -mode=set -var=CoverageVariableName program.go
  
  Flags:
    -V    print version and exit
    -func string
          output coverage profile information for each function
    -html string
          generate HTML representation of coverage profile
    -mode string
          coverage mode: set, count, atomic
    -o string
          file for output; default: stdout
    -var string
          name of coverage variable to generate (default "GoCover")
  
    Only one of -html, -func, or -mode may be set.
  ```
  
    

