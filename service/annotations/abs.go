package annotations

import "data-handler/model"

func IsEnabled(resource *model.Resource, name string) bool {
	return resource.Annotations != nil && resource.Annotations[name] == "true"
}

func Enable(resource *model.Resource, names ...string) {
	if resource.Annotations == nil {
		resource.Annotations = make(map[string]string)
	}

	for _, name := range names {
		resource.Annotations[name] = "true"
	}
}

func Disable(resource *model.Resource, names ...string) {
	if resource.Annotations == nil {
		resource.Annotations = make(map[string]string)
	}

	for _, name := range names {
		resource.Annotations[name] = "false"
	}
}
