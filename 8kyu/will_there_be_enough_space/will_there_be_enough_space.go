package will_there_be_enough_space

func Enough(cap, on, wait int) int {
	if wait <= cap-on {
		return 0
	}
	return wait - (cap - on)

	// My first attempt
	//capable := cap - on - wait
	//
	//if capable < 0 {
	//	if on > wait {
	//		return on - wait
	//	} else if wait > on {
	//		return wait - on
	//	}
	//	return wait
	//}
	//
	//return 0
}
