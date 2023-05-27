package unstructured

func ParseRef(un Unstructured, path string) (Unstructured, error) {
	data, err := JsonPathRead(un, path)

	if err != nil {
		return nil, err
	}

	return data.(Unstructured), nil
}
