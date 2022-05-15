package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	a := []int64{1, 2, 5, 15}

	//создание копии

	b := make([]int64, len(a))
	c := make([]int64, len(a))

	copy(b, a)

	//реверз копии
	for i := len(b) - 1; i >= 0; i-- {
		c = append(c, a[i])
	}

	return c
}
