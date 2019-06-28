package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
  point := &l.Head.Data
  i := l.Head
  for i != nil {
    if comp(i, ref) {
      point = &i.Data
      return point
    }

    i = i.Next
  }

  return nil
}
