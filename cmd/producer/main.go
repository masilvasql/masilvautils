package main

import "github.com/masilvasql/masilvautils/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	err = rabbitmq.Publish(ch, "Hello World", "amq.direct")
	if err != nil {
		panic(err)
	}

}
