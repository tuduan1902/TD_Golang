package src

import (
	"fmt"
)

func greeting(name string) string {
	result := fmt.Sprintf("Hello %s", name)
	return result
}

func demo() {
	w, h, isSquare := rectInfo(100, 200)
	if isSquare {
		fmt.Println("This is square")
	} else {
		fmt.Println("width=", w)
		fmt.Println("heigth", h)
	}
}

// Multiple return values
// bool isSquare = true when width = height
func rectInfo(w, h int) (width int, height int, isSquare bool) {
	isSquare = w == h
	return w, h, isSquare

}
