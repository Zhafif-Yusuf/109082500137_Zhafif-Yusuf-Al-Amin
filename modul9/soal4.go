package main
import "fmt"

const NMAX int = 127
type tabel [NMAX]rune
func isiArray(t *tabel, n *int) {
	var huruf rune
	*n = 0

	for {
		fmt.Scanf("%c", &huruf)
		if huruf == ' ' || huruf == '\n' {
			continue
		}
		if huruf == '.' || *n >= NMAX {
			break
		}
		t[*n] = huruf
		*n = *n + 1
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]), " ")
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	var temp rune

	for i := 0; i < n/2; i++ {
		temp = t[i]
		t[i] = t[n-1-i]
		t[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var data tabel
	var n int

	fmt.Print("teks: ")
	isiArray(&data, &n)

	if palindrom(data, n) {
		fmt.Println("palindrom ? true")
	} else {
		fmt.Println("palindrom ? false")
	}
	fmt.Print("reverse teks: ")
	balikanArray(&data, n)
	cetakArray(data, n)
}