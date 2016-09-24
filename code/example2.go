type tree struct {
    v    int
    l, r *tree
}

func (t *tree) Sum() int {
    sum := t.v
    
    if t.l != nil {
        sum += t.l.Sum()
    }
    
    if t.r != nil {
        sum += t.r.Sum()
    }
    
    return sum
}