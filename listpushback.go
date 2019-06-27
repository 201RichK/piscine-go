package piscine

type t struct {
	Data interface{}
	Next *t
}

type l struct {
	Head *t
}

func ListPushBack(l *l, data interface{}) {
	n := &t{Data: data}
	if l.Head == nil {
		l.Head = n
	} else {
		listItem := l.Head

		for listItem.Next != nil {
			listItem = listItem.Next
		}

		listItem.Next = n
	}
}
