package main

type Sample struct {
	S1      int
	S2      int
	S3      string
	private float64
}

func main() {
	var smp Sample
	smp.S1 = 10
	smp.S2 = 20
	smp.S3 = "Hello World"
}
