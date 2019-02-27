# Slice添加(add)删除(remove)元素

tags： golang code

---

切片算法支持整型
```golang
type Slice []int

func NewSlice() Slice {
    return make(Slice, 0)
}

func (this* Slice) Add(elem int) error {
    for _, v := range *this {
        if v == elem {
            fmt.Printf("Slice:Add elem: %v already exist\n", elem)
            return ERR_ELEM_EXIST
        }
    }
    *this = append(*this, elem)
    fmt.Printf("Slice:Add elem: %v succ\n", elem)
    return nil
}

func (this* Slice) Remove(elem int) error {
    found := false
    for i, v := range *this {
        if v == elem {
            if i == len(*this) - 1 {
                *this = (*this)[:i]

            } else {
                *this = append((*this)[:i], (*this)[i+1:]...)
            }
            found = true
            break
        }
    }
    if !found {
        fmt.Printf("Slice:Remove elem: %v not exist\n", elem)
        return ERR_ELEM_NT_EXIST
    }
    fmt.Printf("Slice:Remove elem: %v succ\n", elem)
    return nil
}
```
切片算法支持字符串
```golang
type Slice []interface{}

func NewSlice() Slice {
    return make(Slice, 0)
}

func (this* Slice) Add(elem interface{}) error {
    for _, v := range *this {
        if v == elem {
            fmt.Printf("Slice:Add elem: %v already exist\n", elem)
            return ERR_ELEM_EXIST
        }
    }
    *this = append(*this, elem)
    fmt.Printf("Slice:Add elem: %v succ\n", elem)
    return nil
}

func (this* Slice) Remove(elem interface{}) error {
    found := false
    for i, v := range *this {
        if v == elem {
            if i == len(*this) - 1 {
                *this = (*this)[:i]

            } else {
                *this = append((*this)[:i], (*this)[i+1:]...)
            }
            found = true
            break
        }
    }
    if !found {
        fmt.Printf("Slice:Remove elem: %v not exist\n", elem)
        return ERR_ELEM_NT_EXIST
    }
    fmt.Printf("Slice:Remove elem: %v succ\n", elem)
    return nil
}
```
<!--stackedit_data:
eyJoaXN0b3J5IjpbMzI5MzcxMDcxXX0=
-->