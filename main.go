package main

import (
	"fmt"
	"sort"
)

type TimeValue struct {
	Value string
	Timestamp int
}

type TimeMap map[string] []TimeValue


/** Initialize your data structure here. */
func Constructor() TimeMap {
	return make(TimeMap)
}


func (tm *TimeMap) Set(key string, value string, timestamp int)  {
	arr, seen := (*tm)[key]

	if !seen {
		arr = make([]TimeValue, 0)
	}

	arr = append(arr, TimeValue{value, timestamp})
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Timestamp < arr[j].Timestamp
	})
	(*tm)[key] = arr
}


func (tm *TimeMap) Get(key string, timestamp int) string {
	tL, in := (*tm)[key]

	if in {
		i := sort.Search(len(tL), func(i int) bool {
			return tL[i].Timestamp >= timestamp
		})

		if i < len(tL) && tL[i].Timestamp == timestamp {
			return tL[i].Value
		} else {
			i--
			return tL[i].Value
		}
	} else {
		return ""
	}
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

func main() {
	obj := Constructor()
	obj.Set("foo", "bar", 1)
	obj.Set("foo", "bar2", 4)
	fmt.Printf("%s\n", obj.Get("foo", 5))
}
