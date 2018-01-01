package main


func isOneBitCharacter(bits []int) bool {
	if len(bits) == 0 {
		return false
	}

	if bits[len(bits)-1] != 0 {
		return false
	}

	i := 0
	for i<len(bits)-1 {
		if bits[i] == 1 {
			i+=2
		}else{
			i+=1
		}
	}

	if i == len(bits)-1 {
		return true
	}
	return false
}