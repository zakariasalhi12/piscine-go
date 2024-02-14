package piscine

func Max(a []int) int {
	resault := 0
	for i := 0; i < len(a); i++ {
		if a[i] > resault {
			resault = a[i]
		}
	}
	return resault
}
