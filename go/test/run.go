package test

func Run(tf func()) {
	prev := filename
	defer func() {
		filename = prev
	}()

	tf()
	filename = "input.txt"
	tf()
}
