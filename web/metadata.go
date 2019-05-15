package web

type MD map[string][]string

func NewMetadata() MD {
	return make(MD)
}

func (md MD) Pair(metadata map[string]string) {
	for key, val := range metadata {
		if m, ok := md[key]; !ok {
			md[key] = []string{val}
		} else {
			md[key] = append(m, val)
		}
	}
}