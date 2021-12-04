package test

type T interface {
	Fatal(args ...interface{})
	Log(args ...interface{})
}

func Check(t T, got, example, input int) {
	want := example
	if *useInput {
		want = input
	}

	t.Log("got: ", got, " want: ", want, " use_input: ", *useInput)

	if got != want {
		t.Fatal(got, want)
	}
}
