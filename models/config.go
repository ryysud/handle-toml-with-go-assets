package models

type Config struct {
	Sample Sample
}

type Sample struct {
	String string
	Int    int
	Float  float32
	Bool   bool
	Array  []string
}
