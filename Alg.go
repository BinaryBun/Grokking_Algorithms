/*
  Algorithms based on the book Aditya Bhargava <<Grokking Algorithms>>
*/
package main

import ("fmt"
)

type deque struct {
  data []string
}

func (d *deque) write (elim string)  { d.data = append(d.data, elim) }

func (d *deque) read () (string) {
  tmp := d.data[0]
  d.data = d.data[1:]
  return tmp
}

func binary_search(data []int, num int) (int) {
  low := 0
  high := len(data)-1

  for low <= high {
    mid_index := (low + high) / 2
    midle_num := data[mid_index]

    if midle_num == num {
      return mid_index
    } else if midle_num > num {
      high = mid_index - 1
    } else {
      low = mid_index + 1
    }
  }
  return -1
}

func findSmallest(data []int) (int) {
  smallest := data[0]
  smallest_index := 0

  for i := 1; i < len(data); i++ {
    if data[i] < smallest {
      smallest = data[i]
      smallest_index = i
    }
  }

  return smallest_index
}

func selectionSort(data []int) {
  new_data := []int {}
  _len_ := len(data)

  for i := 0; i < _len_; i++ {
    smallest := findSmallest(data)
    new_data = append(new_data, data[smallest])

    if smallest != len(data)-1      { data = append(data[:smallest], data[smallest+1:]...)
    } else if len(data) != 1        { data = data[:len(data)-1] }

  }
  fmt.Println(new_data)
}

func pop(data []int, index int) {
  data = append(data[:index], data[index+1:]...)
}

func qsort(data []int) ([]int) {
  if len(data) < 2 {
    return data
  } else {
    support := data[0]
    less := []int{}
    more := []int{}

    for _, i := range(data[1:]) {
      if i <= support { less = append(less, i) }
      if i > support { more = append(more, i) }
    }

    return append(append(qsort(less), support), qsort(more)...)
  }
}

func bfs(dict map[string][]string, start, stop string) (bool) {
  d := deque{}
  // neighbors record
  for _, i := range(dict[start]) { d.write(i) }

  // run 2 deque
  for len(d.data) != 0 {
    current := d.read()
    if current != stop {
      for _, i := range(dict[current]) { d.write(i) }
    } else {
      return true
    }
  }

  return false
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func dijkstr(child map[string]map[string]int, start, stop string) {
  costs_v := map[string]int {start: 0}

  check := []string {}
  d := deque{}
  d.write(start)
  for len(d.data) != 0 {
    now := d.read()
    check = append(check, now)
    // add 2 deque
    for i := range(child[now]) {
      if !contains(d.data, i) && !contains(check, i) { d.write(i) }
    }

    // main code
    for point := range(child[now]) {
      _, flag := costs_v[point]
      if flag {
        if costs_v[point] > child[now][point] + costs_v[now] {
          costs_v[point] = child[now][point] + costs_v[now]
        }
      } else {
        costs_v[point] = child[now][point] + costs_v[now]
      }
    }
  }
  fmt.Println(costs_v)
}

func main() {
  //s := []int{4, 2, 3, 1, 7, 9, 12, 16, 15}
  /*graph :=  map[string][]string {}
  graph["you"] = []string {"alice", "bob", "claire"}
  graph["bob"] = []string {"anuj", "peggy"}
  graph["alice"] = []string {"peggy"}
  graph["claire"] = []string {"thom", "jonny"}
  graph["anuj"] = []string {}
  graph["peggy"] = []string {}
  graph["thom"] = []string {}
  graph["jonny"] = []string {}*/

  graph := map[string]map[string]int {}
  graph["A"] = map[string]int {"B": 3, "C": 10, "D": 1}
  graph["B"] = map[string]int {"A": 3, "F": 6}
  graph["C"] = map[string]int {"A": 10, "D": 2}
  graph["D"] = map[string]int {"A": 1, "C": 2, "E": 5, "F": 4}
  graph["E"] = map[string]int {"D": 5, "F": 2}
  graph["F"] = map[string]int {"B": 6, "D": 4, "E": 2}

  dijkstr(graph, "A", "E")
}
