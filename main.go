package main

import (
	"fmt"
	"sort"
)

type TimeValue struct {
	Value string
	Timestamp int
}

type ValuedArr struct {
  TimeValues []TimeValue
  LastInserted bool
}

type TimeMap map[string] ValuedArr


/** Initialize your data structure here. */
func Constructor() TimeMap {
	return make(TimeMap)
}


func (tm *TimeMap) Set(key string, value string, timestamp int)  {
	valueArr, seen := (*tm)[key]

	if !seen {
		arr := make([]TimeValue, 0)
    valueArr = ValuedArr{arr, true}
	}

	valueArr.TimeValues = append(valueArr.TimeValues, TimeValue{value, timestamp})
	
	(*tm)[key] = valueArr
}


func (tm *TimeMap) Get(key string, timestamp int) string {
	vA, in := (*tm)[key]


	if in {
    if vA.LastInserted {
      sort.Slice(vA.TimeValues, func(i, j int) bool {
        return vA.TimeValues[i].Timestamp < vA.TimeValues[j].Timestamp
      })
    }

    vA.LastInserted = false
    (*tm)[key] = vA

		i := sort.Search(len(vA.TimeValues), func(i int) bool {
			return vA.TimeValues[i].Timestamp >= timestamp
		})

		if i < len(vA.TimeValues) && vA.TimeValues[i].Timestamp == timestamp {
			return vA.TimeValues[i].Value
		} else {
			i--

			if i < 0 {
				return ""
			}

			return vA.TimeValues[i].Value
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
	obj.Set("love", "high", 10)
	obj.Set("love", "low", 20)
	fmt.Printf("%s\n", obj.Get("love", 5))
	fmt.Printf("%s\n", obj.Get("love", 10))
	fmt.Printf("%s\n", obj.Get("love", 15))
	fmt.Printf("%s\n", obj.Get("love", 20))
	fmt.Printf("%s\n", obj.Get("love", 25))
	//obj.Set("foo", "bar2", 4)
	//fmt.Printf("%s\n", obj.Get("foo", 5))
}
