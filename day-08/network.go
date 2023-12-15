package main

type Network struct {
	nodes map[string]NodeLinks
}

type NodeLinks struct {
	left  string
	right string
}
