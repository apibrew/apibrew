package unstructured

type Unstructured map[string]interface{}

func (u Unstructured) MergeInto(un Unstructured) {
	for key, value := range un {
		u[key] = value
	}
}
