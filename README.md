# Golang RabbitMQ Example

```
$ go get github.com/lightstaff/go-rabbitmq-example/protocol
$ go get github.com/streadway/amqp
$ docker pull rabbitmq
$ docker run -d --hostname my-rabbit --name some-rabbit -p 5672:5672 rabbitmq
```

```
$ cd consumer1 && go run main.go -rabbitmqUrl="127.0.0.1:5672"
$ cd consumer2 && go run main.go -rabbitmqUrl="127.0.0.1:5672"
$ cd publisher && go run main.go -rabbitmqUrl="127.0.0.1:5672"
```
