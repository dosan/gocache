```
package main

import (
	"fmt"

	"github.com/dosan/gocache"
)


func main(){
	cache  := gocache.New()
	cache.Set("userId",  300)
	fmt.Println(cache.Get("userId"))
	cache.Delete("userId")
	fmt.Println(cache.Get("userId"))
}
```
