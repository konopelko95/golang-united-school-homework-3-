package homework

func average(input [15]float32) (result float32) {
	//Place your code here

	//объявление массива
	var a [15]float32
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5
	a[5] = 6

	// посчитать сумму массива и //поссчитать кол-во элементов не равных 0
	var sum float32 = 0
	var i float32 = 0
	for _, value := range a {
		if value != 0 {
			i++
		}
		sum += value
	}

	//посчитать среднее
	result = (sum / i)

	return result
}
