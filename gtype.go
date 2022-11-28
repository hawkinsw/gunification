package gunification

type GType interface {
	IsCanonical() bool
	IsConstructor() bool
	Equals(GType) bool
	Constructor() GType
	SubTerms() []GType
	Identifier() string
}
