
## int
```
package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Ints(s)
	fmt.Println(s)
	// [1 2 3 4 5 6]
}
```

## float
```
package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []float64{5.4, 2.1, 3.5, 6.1, -10.5} // unsorted
	sort.Float64s(s)
	fmt.Println(s)
	// [-10.5 2.1 3.5 5.4 6.1]
}
```
## string
```
package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"X", "x", "a", "A", "G"} // unsorted
	sort.Strings(s)
	fmt.Println(s)
	// [A G X a x]
}
```
