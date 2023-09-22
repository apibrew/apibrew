package annotations

import (
	"golang.org/x/net/context"
	"strings"
)

type Annotated interface {
	GetAnnotations() map[string]string
}

const ctxValue = "annotationsCtx"

func WithContext(parent context.Context, annotated Annotated) context.Context {
	existingAnnotations := FromCtx(parent).GetAnnotations()
	return context.WithValue(parent, ctxValue, merge(existingAnnotations, annotated.GetAnnotations()))
}

func merge(annotations map[string]string, annotations2 map[string]string) Annotated {
	if annotations == nil {
		annotations = make(map[string]string)
	}

	if annotations2 == nil {
		annotations2 = make(map[string]string)
	}

	for key, value := range annotations2 {
		annotations[key] = value
	}

	return &annotated{
		annotations: annotations,
	}
}

func SetWithContext(parent context.Context, name, value string) context.Context {
	val := FromCtx(parent).GetAnnotations()
	if val == nil {
		val = make(map[string]string)
	}
	val[name] = value

	return WithContext(parent, &annotated{
		annotations: val,
	})
}

type annotated struct {
	annotations map[string]string
}

func (a *annotated) GetAnnotations() map[string]string {
	return a.annotations
}

//goland:noinspection GoUnusedExportedFunction
func FromCtx(ctx context.Context) Annotated {
	if annotations, ok := ctx.Value(ctxValue).(Annotated); ok {
		return annotations
	}

	return &annotated{
		annotations: make(map[string]string),
	}
}

func IsEnabled(resource Annotated, name string) bool {
	return resource.GetAnnotations() != nil && resource.GetAnnotations()[name] == "true"
}

func IsEnabledOnCtx(ctx context.Context, name string) bool {
	return IsEnabled(FromCtx(ctx), name)
}

func Enable(resource Annotated, names ...string) {
	for _, name := range names {
		resource.GetAnnotations()[name] = "true"
	}
}

func EnableWith(annotations map[string]string, names ...string) map[string]string {
	if annotations == nil {
		annotations = make(map[string]string)
	}
	for _, name := range names {
		annotations[name] = "true"
	}
	return annotations
}

func Set(resource Annotated, name, value string) {
	resource.GetAnnotations()[name] = value
}

func Get(resource Annotated, name string) string {
	return resource.GetAnnotations()[name]
}

//goland:noinspection GoUnusedExportedFunction
func Disable(resource Annotated, names ...string) {
	for _, name := range names {
		resource.GetAnnotations()[name] = "false"
	}
}

func ToString(resource Annotated) string {
	var parts []string

	for key, value := range resource.GetAnnotations() {
		parts = append(parts, key+"="+value)
	}

	return strings.Join(parts, ";")
}

func IsSame(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}

	for key := range a {
		if a[key] != b[key] {
			return false
		}
	}

	return true
}
