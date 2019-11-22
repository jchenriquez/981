package main

import (
  "math"
	"fmt"
)

type Node struct {
  parent *Node
  left *Node
  right *Node
  height int
  timeStamp int
  val interface {}
}


func updateTreeHeight(node *Node) {
  if node.left == nil && node.right == nil {
    node.height = 1
  } else if node.right == nil {
    node.height = node.left.height+1
  } else if node.left == nil {
    node.height = node.right.height+1
  } else {
    node.height = int(math.Max(float64(node.left.height), float64(node.right.height)))+1
  }
}

func shiftLeft(node *Node) *Node {
  rightSub := node.right
  parent := node.parent

  node.right = rightSub.left
  rightSub.left = node

  if parent != nil {
    if parent.left == node {
      parent.left = rightSub
    } else {
      parent.right = rightSub
    }
  }
  
  node.parent = rightSub
  rightSub.parent = parent

  updateTreeHeight(node)
  updateTreeHeight(rightSub)

  return rightSub
}

func shiftRight(node *Node) *Node {
  leftSub := node.left
  parent := node.parent

  node.left = leftSub.right
  leftSub.right = node

  if parent != nil {
    if parent.left == node {
      parent.left = leftSub
    } else {
      parent.right = leftSub
    }
  }

  node.parent = leftSub
  leftSub.parent = parent

  updateTreeHeight(node)
  updateTreeHeight(leftSub)

  return leftSub
}

func balance (tree *Node) *Node {

  leftSub := tree.left
  rightSub := tree.right

  if leftSub == nil  {
    if rightSub.height >= 2 {
      tree = shiftLeft(tree)
    } else {
      updateTreeHeight(tree)
    }
  } else if rightSub == nil {
    if leftSub.height >= 2 {
      tree = shiftRight(tree)
    } else {
      updateTreeHeight(tree)
    }
  } else {
    if math.Abs(float64(leftSub.height - rightSub.height)) >= 2 {
      if leftSub.height > rightSub.height {
        tree = shiftRight(tree)
      } else {
        tree = shiftLeft(tree)
      }
    } else {
      updateTreeHeight(tree)
    }
  }
  if tree.parent == nil {
    return tree
  }
  return balance(tree.parent)
}

func (tr *Node) Insert (val interface{}, timeStamp int) *Node {
  curNode := tr

  for curNode != nil {
    switch {
      case timeStamp > curNode.timeStamp:
        if curNode.right == nil {
          curNode.right = &Node{curNode, nil, nil, 1, timeStamp, val}
          return balance(curNode)
        }
        curNode = curNode.right
      case timeStamp < curNode.timeStamp:
        if curNode.left == nil {
          curNode.left = &Node{curNode, nil, nil, 1, timeStamp, val}
          return balance(curNode)
        }  
        curNode = curNode.left
      default:
        curNode.val = val
        return tr
    }
  }

  return tr
}

func (tr *Node) Search (timeStamp int) interface{} {
  currNode := tr

  for currNode != nil {
    switch {
      case currNode.timeStamp < timeStamp:
        if currNode.right == nil {
          return currNode.val
        }
        currNode = currNode.right
      case currNode.timeStamp > timeStamp:
        if currNode.left == nil {
          parent := currNode.parent

          for parent != nil {

            if parent.right == currNode {
              return parent.val
            }

            currNode = parent
            parent = currNode.parent
          }

          return nil
        }
        currNode = currNode.left
      default:
        return currNode.val
    }
  }

  return nil
}

type TimeMap map[string]*Node 


/** Initialize your data structure here. */
func Constructor() TimeMap {
	return make(TimeMap)
}


func (tm *TimeMap) Set(key string, value string, timeStamp int)  {
	tree, seen := (*tm)[key]

	if !seen {
		(*tm)[key] = &Node{nil, nil, nil, 1, timeStamp, value}
    return
	}

	(*tm)[key] = tree.Insert(value, timeStamp)
}


func (tm *TimeMap) Get(key string, timeStamp int) string {
	tree, in := (*tm)[key]
  if in {
    val := tree.Search(timeStamp)

    if val == nil {
      return ""
    }

    return val.(string)
  }

  return ""
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
  obj.Set("love", "low2", 30)
  obj.Set("love", "low3", 40)
  obj.Set("love", "low4", 50)
  obj.Set("love", "low5", 60)
	obj.Set("foo", "bar2", 4)
	fmt.Printf("%s\n", obj.Get("foo", 5))
	fmt.Printf("%s\n", obj.Get("love", 10))
	fmt.Printf("%s\n", obj.Get("love", 15))
	fmt.Printf("%s\n", obj.Get("love", 20))
	fmt.Printf("%s\n", obj.Get("love", 25))
	fmt.Printf("%s\n", obj.Get("foo", 5))
}
