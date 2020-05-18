package core

type PropertyResolver interface {
	ContainsProperty(name string) bool
	GetProperty(name string, defaultValue string) string
}
