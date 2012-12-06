package villa

import(
    "testing"
//    "sort"
    "math/rand"
    "fmt"
)

func TestIntPriorityQueue(t *testing.T) {
    fmt.Println("== Begin TestIntPriorityQueue...");
    defer fmt.Println("== End TestIntPriorityQueue.");
    
    pq := NewIntPriorityQueue(func(e1, e2 int) bool {
        return e1 < e2
    })
    for i := 0; i < 1000; i ++ {
        pq.Push(rand.Int())
    } // for i
    
    AssertEquals(t, "pq.Len()", pq.Len(), 1000)

    peek := pq.Peek()
    last := pq.Pop()
    AssertEquals(t, "pg.Peek()", peek, last)
    for i := 1; i < 1000; i ++ {
        cur := pq.Pop()
        if cur < last {
            t.Errorf("%d should be larger than %d", cur, last)
        } // if
        last = cur
    } // for i
    fmt.Println(pq)
}
 