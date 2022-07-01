/*
  Algorithms based on the book Aditya Bhargava <<Grokking Algorithms>>
*/
package main

import ("fmt")

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

func bfs(dict map[string][]string) {
  // nothing
}

func main() {
  s := []int{4, 2, 3, 1, 7, 9, 12, 16, 15}
  graph :=  map[string][]string {}
  graph["you"] = []string {"alice", "ЬоЬ", "claire"}
  graph["bob"] = []string {"anuj", "peggy"}
  graph["alice"] = []string {"peggy"}
  graph["claire"] = []string {"thom", "jonny"}
  graph["anuj"] = []string {}
  graph["peggy"] = []string {}
  graph["thom"] = []string {}
  graph["jonny"] = []string {}


  fmt.Println(qsort(s))
}
