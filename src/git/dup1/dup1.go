package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

type Student struct {
    name string
    age int
    school string
}

func (s Student) String() string {
	return s.name + "," + strconv.Itoa(s.age) + "," + s.school
}

func GetStudent(name string, age int, school string) (s *Student) {
	return &Student{name, age, school}
}
func main() {
  john := Student{"john", 25, "Santa Clara University"}
  mary := GetStudent("mary", 22, "Stanford Univeristy")
  fmt.Println(john)
  fmt.Println(*mary)
  counts := make(map[string]int)
  temp := make(map[string]int)
  temp["a"] = 0
  temp["b"] = 1
  fmt.Println(temp)
  input := bufio.NewScanner(os.Stdin)
  for input.Scan() {
    counts[input.Text()]++
  }
  for line, n := range counts {
    if n > 1 {
      fmt.Println(n, line)
    }
  }
}

