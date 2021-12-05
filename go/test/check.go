package test

type T interface {
	Fatal(args ...interface{})
	Log(args ...interface{})
}

func Check(t T, got, example, input int) {
	want := example
	if filename == "input.txt" {
		want = input
	}

	t.Log("[", filename, "]", "got: ", got, " want: ", want, " ok: ", got == want)

	if got != want {
		t.Fatal(got, want)
	}
}
