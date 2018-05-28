package main
import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  counts := make(map[string]int)
  input := bufio.NewScanner(os.Stdin)

  for input.Scan() {
    if(input.Text() == "0") {
      break
    }
    counts[input.Text()]++
  }

  //Примечание: игнорируем потенциальные ошибки из input.Err()

  for line, n := range counts {
    if n > 1 {
      fmt.Printf("%s - %d раз\n", line, n)
    }
  }
}
