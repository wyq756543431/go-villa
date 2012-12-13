package villa

import(
    "testing"
    "sort"
    "math/rand"
    "fmt"
)

func TestFloatSlice(t *testing.T) {
    fmt.Println("== Begin TestFloatSlice...");
    defer fmt.Println("== End TestFloatSlice.");
    
    var s FloatSlice
    for i := 0; i < 1000; i ++ {
        s.Add(float64(i))
    } // for i
    
    AssertEquals(t, "len", len(s), 1000)
    //fmt.Println(s)
    s.Clear()
    AssertEquals(t, "len", len(s), 0)
    
    s = FloatSlice{}
    s.Add(1)
    s.Insert(0, 2)
    s.Insert(1, 3)
    fmt.Println(s)
    AssertEquals(t, "len", len(s), 3)
    AssertStringEquals(t, "s", s, "[2 3 1]")
    
    sort.Sort(s.NewSortList(FloatValueCompare))
    AssertStringEquals(t, "s", s, "[1 2 3]")
}

func TestFloatSliceRemove(t *testing.T) {
    fmt.Println("== Begin TestFloatSliceRemove...");
    defer fmt.Println("== End TestFloatSliceRemove.");
    var s FloatSlice
    s.Add(1, 2, 3, 4, 5, 6, 7)
    AssertEquals(t, "len", len(s), 7)
    AssertStringEquals(t, "s", s, "[1 2 3 4 5 6 7]")
    
    s.RemoveRange(2, 5)
    AssertEquals(t, "len", len(s), 4)
    AssertStringEquals(t, "s", s, "[1 2 6 7]")
    
    s.Remove(2)
    AssertEquals(t, "len", len(s), 3)
    AssertStringEquals(t, "s", s, "[1 2 7]")
}

func TestFloatSliceSort(t *testing.T) {
    var s FloatSlice
    for i := 0; i < 100; i ++ {
        s.Add(rand.Float64())
    } // for i
    
    //fmt.Println(s)
    
    adp := s.NewSortList(FloatValueCompare)
    sort.Sort(adp)
    
    //fmt.Println(s)
    for i := 1; i < len(s); i ++ {
        if s[i - 1] > s[i] {
            t.Errorf("s[%d](%v) is supposed to be less or equal than s[%d](%v)", i - 1, s[i - 1], i, s[i])
        } //  if
    } //  if
    
    for i := range(s) {
        p, found := adp.BinarySearch(s[i])
        AssertEquals(t, fmt.Sprintf("%d found", i), found, true)
        if found {
            AssertEquals(t, fmt.Sprintf("%d found element", i), s[p], s[i])
        } // if
    } // for i
    
    for i := range(s) {
        e := rand.Float64()
        p, found := adp.BinarySearch(e)
        if found {
            AssertEquals(t, fmt.Sprintf("found element", i), s[p], e)
        } else {
            beforeOk := p == 0 || s[p - 1] <= e;
            afterOk := p == len(s) || s[p] >= e;
            
            if !beforeOk || !afterOk {
                t.Errorf("Wrong position %d for %v", p, e)
            } // if
        } // else
    } // for i
}

func BenchmarkFloatSliceInsert(b *testing.B) {
    b.StopTimer()
    s := make(FloatSlice, 100000, 100000)
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        s.Insert(1, float64(i))
    } // for i
}

func BenchmarkFloatSliceInsertByAppend(b *testing.B) {
    b.StopTimer()
    s := make([]float64, 100000, 100000)
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        s = append(s[:1], append([]float64{float64(i)}, s[1:]...)...)
    } // for i
}

func BenchmarkFloatSliceInsertByCopy(b *testing.B) {
    b.StopTimer()
    s := make([]float64, 100000, 100000)
    b.StartTimer()
    
    for i := 0; i < b.N; i ++ {
        s = append(s, 0)
        copy(s[2:], s[1:])
        s[1] = float64(i)
    } // for i
}
  