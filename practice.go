
// GO basic practice
package main
import "fmt"

func main() {
  fmt.Println("Hello World!")
  for i := 1; i <= 5; i++ {
    fmt.Println(i)
  }
  for i := 1; i <= 5; i++ {
    fmt.Print(i, " ")
  }
  var sum = 0;
  for i := 1; i <= 5; i++ {
    sum += i;
  }
  fmt.Println();
  fmt.Println(sum);
}