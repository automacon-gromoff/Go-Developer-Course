package main

func main() {
	s := "abc"
	change(&s)
}

func change(s *string) {
	*s = "cba"
}
