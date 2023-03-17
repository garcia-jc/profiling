package functional

func forEach[In, Out any](in []In, mapFn func(In) Out) []Out {
	out := make([]Out, len(in))
	for i := range in {
		out[i] = mapFn(in[i])
	}
	return out
}

func filter[T any](in []T, filterFn func(T) bool) []T {
	out := make([]T, 0, len(in))
	for _, v := range in {
		if filterFn(v) {
			out = append(out, v)
		}
	}
	return out
}

func reduce[In, Out any](in []In, initial Out, reductor func(Out, In) Out) Out {
	end := initial
	for _, v := range in {
		end = reductor(end, v)
	}
	return end
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
