//Вывод аргументов командной строки

package main
import (
  "fmt"
  "os"
  "strconv"
  "time"
)

func main() {
  var s, sep string
  start := time.Now()

  for i, arg := range os.Args[0:] {
    s += sep + strconv.Itoa(i) + " - " + arg
    sep = ", "
  }

  fmt.Println(s)

  fmt.Println(time.Since(start).Seconds())
}
