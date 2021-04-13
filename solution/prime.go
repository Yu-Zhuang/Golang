package solution

/*
if N = 10000
result show:
FindPrimeInNum spend time:  45.724425ms
FindPrimeInNum2 spend time:  2.645978885s
FindPrimeInNum3 spend time:  4.788341021s
in my MacBook Air
*/

import "math"

func FindPrimeInNum(num int) (ret []int) {
	if num < 2 {
		return
	}
	sq := int(math.Sqrt(float64(num))) + 1
	mp := []bool{}
	for i := 0; i <= num; i++ {
		mp = append(mp, true)
	}
	for i := 2; i < sq; i++ {
		for k := i; k*i < num; k++ {
			mp[k*i] = false
		}
	}
	for i := 2; i <= num; i++ {
		if mp[i] == true {
			ret = append(ret, i)
		}
	}
	return
}

func FindPrimeInNum2(num int) (ret []int) {
	if num < 2 {
		return ret
	}
	ret = append(ret, 2)
	for i := 3; i <= num; i += 2 {
		for j := 3; j <= i; j += 2 {
			if i == j {
				ret = append(ret, i)
			} else if i%j == 0 {
				break
			}
		}
	}
	return
}

func FindPrimeInNum3(num int) (ret []int) {
	if num < 2 {
		return ret
	}
	for i := 2; i <= num; i++ {
		for j := 2; j <= i; j++ {
			if i == j {
				ret = append(ret, i)
			} else if i%j == 0 {
				break
			}
		}
	}
	return
}
