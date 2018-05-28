package main
import (
  "io/ioutil"
  "fmt"
  "os"
  "strings"
)

func main() {
  counts := make(map[string]int)
  files := make(map[string]bool)

  for _, filename := range os.Args[1:] {
    data, err := ioutil.ReadFile(filename)

    if err != nil {
      fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
      continue
    }

    strings := strings.Split(string(data), "\n")

    for _, line :=  range strings {
      if line == " " {
        continue
      }

      counts[line]++

      if files[filename] != true {
          if counts[line] > 1 {
            files[filename] = true
          }
      }
    }
  }

  fmt.Printf("%s", "В следующих файлах найдены повторения:\n")
  for fileName, haveRepeats := range files {
    if haveRepeats {
      fmt.Println(fileName)
    }
  }

  fmt.Println("=======================")

  for line, number := range counts {
    if(number > 1) {
      fmt.Printf("%s - %d раз\n", line, number)
    }
  }

}
