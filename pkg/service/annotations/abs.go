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
	return context.WithValue(parent, ctxValue, annotated)
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
