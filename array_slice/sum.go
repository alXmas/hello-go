package array_slice

func Sum(numbers [5]int) int {
  var result int
  for i := 0; i < 5; i++ {
    result += numbers[i]
  }

  return result
}
