package main

func main() {
	e := createRouter()
	e.Logger.Fatal(e.Start(":1337"))
}
