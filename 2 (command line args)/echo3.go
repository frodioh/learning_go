//Вывод аргументов командной строки

package main
import (
  "fmt"
  "os"
)

func main() {

  fmt.Println(os.Args[1:])

}
