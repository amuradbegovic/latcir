package latcir

import (
	"testing"
)

func TestJednaRijec(t *testing.T) {

	rijec := []rune("slovo")
	trazim := "слово"
	rezultat := Latcir(rijec)

	if rezultat != trazim {
		t.Fatalf(`Latcir("slovo") = %q, tražio sam %q`, rezultat, trazim)
	}

}

func TestViseRijeci(t *testing.T) {

	rijeci := []rune("Riba ribi grize rep.")
	trazim := "Риба риби гризе реп."
	rezultat := Latcir(rijeci)

	if rezultat != trazim {
		t.Fatalf(`Latcir("Riba ribi grize rep.") = %q, tražio sam %q`, rezultat, trazim)
	}

}

func TestLNaKraju(t *testing.T) {

	rijec := []rune("val")
	trazim := "вал"
	rezultat := Latcir(rijec)

	if rezultat != trazim {
		t.Fatalf(`Latcir("val") = %q, trazio sam %q`, rezultat, trazim)
	}
}
