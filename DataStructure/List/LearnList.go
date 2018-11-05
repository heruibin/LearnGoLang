package List

import (
	list "container/list"
	"fmt"
)

func main() {
	var l list.List

	for i := 0; i <= 5; i++ {
		l.PushBack(i)
	}

	l.PushBack(6)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println()

	l.MoveAfter(l.Front(),l.Front().Next())

	for e := l.Front(); e != nil; e = e.Next(){
		fmt.Println(e.Value)
	}

	for e := l.Front() ; e != nil; e = e.Next(){
		if e.Value == 3 {
			l.Remove(e)
		}
	}
	fmt.Println()
	for e := l.Front(); e != nil; e = e.Next(){
		fmt.Println(e.Value)
	}

	for e := l.Front();e != nil ;e = e.Next()  {
		if e.Value == 4 {
			l.MoveAfter(e,e.Next())
		}
	}
	fmt.Println()
	for e := l.Front(); e != nil; e = e.Next(){
		fmt.Println(e.Value)
	}
}
