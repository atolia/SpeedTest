package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func main() {
  f, _ := os.Open("BIGDATA2")
  sum := 0
  x := 0
  reader := bufio.NewReaderSize(f, 1024*1024)
  scanner := bufio.NewScanner(reader)

  scanner.Split(bufio.ScanLines)

  for scanner.Scan() {
    x, _ = strconv.Atoi(scanner.Text())
    sum += x
  }

  fmt.Println(sum)
}

