package main

const (
	TypeToggle int = 0
)

// Module represents a physical module connected to an Arduino
type Module struct {
	ID          int
	Name        string
	Description string
	Pin         int
	Type        int
}
