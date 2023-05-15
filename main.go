package main

func main() {
	e, err := InitializeEvent("Hello, World!")
	if err != nil {
		panic(err)
	}

	e.Start()
}
