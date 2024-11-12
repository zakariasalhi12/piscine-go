package piscine

func TrimAtoi(s string) int {
	resault := 0
	counter := 0
	sign := 1

	for _, char := range s {
		if char == '-' && resault == 0 && counter == 0 {
			counter++
			sign = -1
		}
		if char >= '0' && char <= '9' {
			resault = resault*10 + int(char-'0')
		}
	}
	return resault * sign
}
