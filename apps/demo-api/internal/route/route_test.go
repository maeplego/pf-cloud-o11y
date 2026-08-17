package route

import "testing"

func TestPatternNormalizesWorkID(t *testing.T) {
	got := Pattern("GET", "/work/abc-123")
	want := "GET /work/:id"
	if got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}

func TestPatternRoot(t *testing.T) {
	if Pattern("GET", "/") != "GET /" {
		t.Fatal()
	}
}
