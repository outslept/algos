package perfectnumber

func isPerfectNumber(n int) bool {
	sum := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if i*i != n {
				sum += n / i
			}
		}
	}
	return sum == n && n != 1
}
