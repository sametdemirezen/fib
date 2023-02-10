package fib

        import "testing"

        func TestFib(t *testing.T) {
          type test struct {
            input int
            want int
          }

          var tests = []test {
            {0,0},
            {1,1},
            {2,1},
            {3,2},
            // her kan vi bare legge på nye test cases
          }

          // tc - test case
          for _, tc := range tests {
              got := Fib(tc.input)
              if got != tc.want {
                  t.Errorf("Fib(%d): want %d, got %d", tc.input, tc.want, got)
              }
          }
	   /* input := 0
          got := Fib(input)
          want := 0
          if got != want {
              t.Errorf("Fib(%d): want %d, got %d", input, want, got)
          }
      RESPONSE#1 FAIL github.com/sametdemirezen/fib/fibb [build failed]
      *************************************************

        RESPONSE#2
     === RUN   TestFib
--- PASS: TestFib (0.00s)
PASS
ok  	github.com/sametdemirezen/fib/fibb	0.027s
********************************************************
input := 0
          got := Fib(input)
          want := 1
          if got != want {
              t.Errorf("Fib(%d): want %d, got %d", input, want, got)
          }     

     RESPONSE#3
     === RUN   TestFib
    fib_test.go:10: Fib(0): want 1, got 0
--- FAIL: TestFib (0.00s)
FAIL
FAIL	github.com/sametdemirezen/fib/fibb	0.004s
FAIL
*******************************************************
input = 1 // legg merke til operatøren =
          got = Fib(input)
          want = 1
          if got != want {
              t.Errorf("Fib(%d): want %d, got %d", input, want, got)
          }
        RESPONSE#4
          === RUN   TestFib
    fib_test.go:10: Fib(0): want 1, got 0
    fib_test.go:17: Fib(1): want 1, got 0
--- FAIL: TestFib (0.00s)
FAIL
FAIL	github.com/sametdemirezen/fib/fibb	0.004s
FAIL
********************************************************
        RESPONSE#5
    === RUN   TestFib
--- PASS: TestFib (0.00s)
PASS
ok  	github.com/sametdemirezen/fib/fibb	(cached)
*********************************************************
        }
package fib

 import "testing"


func TestFib(t *testing.T) {
          type test struct {
            input int
            want int
          }

          var tests = []test {
            {0,0},
            {1,1},
            {2,1},
            // her kan vi bare legge på nye test cases
          }

          // tc - test case
          for _, tc := range tests {
              got := Fib(tc.input)
              if got != tc.want {
                  t.Errorf("Fib(%d): want %d, got %d", tc.input, tc.want, got)
              }
          }
	  }
	  
        RESPONSE#6 
    === RUN   TestFib
--- PASS: TestFib (0.00s)
PASS
ok  	github.com/sametdemirezen/fib/fibb	(cached)
*********************************************************
package fib

        import "testing"

        func TestFib(t *testing.T) {
          type test struct {
            input int
            want int
          }

          var tests = []test {
            {0,0},
            {1,1},
            {2,1},
            {3,2},
            // her kan vi bare legge på nye test cases
          }

          // tc - test case
          for _, tc := range tests {
              got := Fib(tc.input)
              if got != tc.want {
                  t.Errorf("Fib(%d): want %d, got %d", tc.input, tc.want, got)
              }
          }
        }
        RESPONSE#7
    === RUN   TestFib
    fib_test.go:23: Fib(3): want 2, got 1
--- FAIL: TestFib (0.00s)
FAIL
FAIL	github.com/sametdemirezen/fib/fibb	0.004s
FAIL
     
*/
          }
        