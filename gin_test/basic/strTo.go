package main  

import (
  "fmt"
  "github.com/Unknwon/com"

)


func main() {
	state :=15
	a  := com.StrTo(state).MustInt()
	fmt.Println(a)
}