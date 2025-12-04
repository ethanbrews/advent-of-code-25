package main

func greatestTwoDigitCombination(s string) int {
	// Scan backwards starting with b = s[-1] and a = s[-2]
	a, b := s[len(s)-2], s[len(s)-1]

	for i := len(s) - 3; i >= 0; i-- {
		if s[i] >= a {
			if a > b {
				b = a
			}
			a = s[i]
		}
	}

	return (int(a-'0') * 10) + int(b-'0')
}

func greatest12DigitCombination(s string) int {
	const k = 12
	n := len(s)
	if k > n {
		return 0
	}

	result := make([]byte, 0, k)
	lastPos := -1

	for digitsNeeded := k; digitsNeeded > 0; digitsNeeded-- {
		maxDigit := byte('0')
		maxPos := lastPos + 1

		for j := lastPos + 1; j <= n-digitsNeeded; j++ {
			if s[j] > maxDigit {
				maxDigit = s[j]
				maxPos = j
			}
		}

		result = append(result, maxDigit)
		lastPos = maxPos
	}

	val := 0
	for _, digit := range result {
		val = val*10 + int(digit-'0')
	}
	return val
}

func day03Part1() int {
	return day03(greatestTwoDigitCombination)
}

func day03Part2() int {
	return day03(greatest12DigitCombination)
}

func day03(extractor func(string) int) int {
	puzzleData, err := ReadLines("data/day03.txt", "\n")
	if err != nil {
		panic(err)
	}

	var total = 0

	for _, line := range puzzleData {
		total += extractor(line)
	}

	return total
}
