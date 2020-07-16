package main
type room struct {
	forward chan []byte
	join chan *client
	leave chan *client
	clinents map[*client]bool
}
