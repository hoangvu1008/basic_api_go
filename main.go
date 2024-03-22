package main

func main() {
	store := NewStore(db)
	api := NewApiServer(":3000", nil)
}
