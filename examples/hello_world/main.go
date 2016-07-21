package main

import (
	"fmt"
	"log"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	"github.com/cayleygraph/cayley/quad"
)

func main() {
	// Create a brand new graph
	store, err := cayley.NewMemoryGraph()
	if err != nil {
		log.Fatalln(err)
	}

	store.AddQuad(quad.Quad{quad.String("phrase of the day"), quad.String("is of course"), quad.String("Hello World!"), quad.String("demo graph")})

	// Now we create the path, to get to our data
	p := cayley.StartPath(store, quad.String("phrase of the day")).Out(quad.String("is of course"))

	// Now we get an iterator for the path (and optimize it, the second return is if it was optimized,
	// but we don't care for now)
	it, _ := p.BuildIterator().Optimize()
	// remember to cleanup after yourself
	defer it.Close()

	// Now for each time we can go to next iterator
	nxt := graph.AsNexter(it)
	// remember to cleanup after yourself
	defer nxt.Close()

	// While we have items
	for nxt.Next() {
		token := it.Result()          // get a ref to a node
		value := store.NameOf(token)  // get the value in the node
		nativeValue := value.Native() // this converts nquad escaped string to normal type

		fmt.Println(nativeValue) // print it!
	}
}
