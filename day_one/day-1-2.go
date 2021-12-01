package day_one

func DoDayOnePartTwo() int {
	input := readDayOneInput("day_one/input_1-1.txt")

	var measurementSlides []int

	for i := 0; i < len(input)-2; i++ {
		measurementSlide := 0
		for j := 0; j < 3; j++ {
			measurementSlide += input[i+j]
		}

		measurementSlides = append(measurementSlides, measurementSlide)
	}

	increases := 0

	for i := 0; i < len(measurementSlides); i++ {
		if i == 0 {
			continue
		}

		if measurementSlides[i-1] < measurementSlides[i] {
			increases++
		}
	}

	return increases
}
