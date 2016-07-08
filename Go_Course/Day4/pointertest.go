package main

import (
  "fmt"
  "errors"
)

type intListNode struct {
  val int
  next *intListNode
}

func createIntList(a []int) *intListNode {
    if len(a) == 0 {
        return nil
    }
    var prev *intListNode
    for i := len(a) - 1; i >= 1; i-- {
        node := intListNode{val: a[i], next: prev}
        prev = &node
    }
    return &intListNode{val: a[0], next: prev}
}

func (head *intListNode) index(n int) (int, error) {
    if n < 0 {
        return 0, errors.New("Out of bounds")
    }
    for i := 0; i < n; i++ {
        if head == nil {
            return 0, errors.New("Out of bounds")
        }
        head = head.next
    }
    if head == nil {
        return 0, errors.New("Out of bounds")
    }
    return head.val, nil
}

func (li *intListNode) append(val int) *intListNode {
  head := li
  item := intListNode{val:val, next: nil}
  if li == nil {
    return &item
  }
  for {
    if li.next == nil {
      li.next = &item
      return head
    }
    li = li.next
  }
  return nil
}

func (li *intListNode) remove(idx int) *intListNode {
  head := li
  if head == nil {
    return head
  }
  if idx == 0 {
    return head.next
  }
  targetNode := li
  var prev *intListNode

  for i := 0 ; i < idx ; i++ {
    prev = targetNode
    targetNode = targetNode.next

    if targetNode.next == nil && i != idx {
      return head
    }
  }

  after := targetNode.next

  if after == nil {
    prev.next = nil
    return head
  } else {
    prev.next = after
  }

  return head
}

func main () {
  list := createIntList([]int{7,2,5,-8})
  list = list.append(12)
  a, err := list.index(0)
  if err != nil {
    fmt.Println("Out of bounds.")
    return
  }
  fmt.Println(a)

  list = list.remove(0)
  a, err = list.index(0)
  if err != nil {
    fmt.Println("Out of bounds.")
    return
  }

  fmt.Println(a)

}
