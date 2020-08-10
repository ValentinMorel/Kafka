# Kafka

![Kafka](https://miro.medium.com/max/687/1*6SQiJ4tinE0p4sjoBexxuA.png)

Apache Kafka® is a distributed streaming platform. What exactly does that mean?
A streaming platform has three key capabilities:

 - Publish and subscribe to streams of records, similar to a message queue or enterprise messaging system.
 - Store streams of records in a fault-tolerant durable way.
 - Process streams of records as they occur.

Kafka is generally used for two broad classes of applications:

 - Building real-time streaming data pipelines that reliably get data between systems or applications
 - Building real-time streaming applications that transform or react to the streams of data

## Tech

This Project uses the following external libraries : 

* Echo - High Performance Web API
* Sarama - Kafka for Golang


## Installation

1. Download kafka package and start a standalone instance
```
# get a quick-and-dirty single-node ZooKeeper instance
bin/zookeeper-server-start.sh config/zookeeper.properties

# start the Kafka server
bin/kafka-server-start.sh config/server.properties
```

2. Create a topic "test" and list topics

```
# create a topic named "test" with a single partition and only one replica
bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic test

# list topic
bin/kafka-topics.sh --list --bootstrap-server localhost:9092
```


3. Send and receive messages

```
# Run the producer and then type a few messages into the console to send to the server
bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test

# command line consumer that will dump out messages to standard output
bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --from-beginning
```
## Features

Dillinger is currently extended with the following plugins. Instructions on how to use them in your own application are linked below.

| Feature | Action | State |
| ------ | ------ | ------ |
| Consumer | Listening on Kafka Broker | Done |
| Consumer Version | Zookeeper is the Newest, to use | TODO | 
| Producer | Triggering Message Sending on Kafka Broker | Done |
| RestAPI | Triggering actions with routes on API | TODO |
| RequestMethod | Requesting API with curl -X POST cmd | TODO |
| HealthChecker | Pinging the Broker to ensure his availability | TODO |


## Usage

After Installing the needed packages : 

```
go run consumer/main.go KAFKA_URL TOPIC
go run producer/main.go
```

expected output : 
```
Received messages  Something to test broker
```

