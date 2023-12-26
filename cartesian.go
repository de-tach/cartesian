package cartesian

// Iter takes interface-slices and returns a channel, receiving cartesian products
func Iter[T any](params ...[]T) chan []T {
	// create channel
	c := make(chan []T)
	if len(params) == 0 {
		close(c)
		return c // Return a safe value for nil/empty params.
	}
	go func() {
		iterate(c, params[0], []T{}, params[1:]...)
		close(c)
	}()
	return c
}

func iterate[T any](channel chan []T, topLevel, result []T, needUnpacking ...[]T) {
	if len(needUnpacking) == 0 {
		for _, p := range topLevel {
			channel <- append(append([]T{}, result...), p)
		}
		return
	}
	for _, p := range topLevel {
		iterate(channel, needUnpacking[0], append(result, p), needUnpacking[1:]...)
	}
}
