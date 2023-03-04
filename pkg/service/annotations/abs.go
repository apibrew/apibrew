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

type annotated struct {
	annotations map[string]string
}

func (a *annotated) GetAnnotations() map[string]string {
	return a.annotations
}

func FromCtx(ctx context.Context) Annotated {
	return &annotated{
		annotations: ctx.Value(ctxValue).(map[string]string),
	}
}

func IsEnabled(resource Annotated, name string) bool {
	return resource.GetAnnotations() != nil && resource.GetAnnotations()[name] == "true"
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
