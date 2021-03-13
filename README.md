## Learn Kafka
Learning kafka using Go and library shopify sarama

### Running Snapshot

#### Kafka Publisher
![Kafka-Publisher](./snapshot/kafka-producers.png)

#### Kafka Consumer
![Kafka-Consumer](./snapshot/kafka-consumer.png)

### Publisher
* Using wordlist `word.txt` for dummy publish message

### Running
* Setup docker container
  ```
  bash run-docker.sh
  ```

* Setup topic name
  ```
  export KAFKA_TOPIC="kafka-topic"
  ```

* Setup broker & zookeeper connection
  ```
  export KAFKA_BROKER_URL="kafka:9092"
  export ZOOKEEPER_URL="zookeeper:2081"
  ```

* Setup hosts
  ```
  vi /etc/hosts
  ----
  127.0.0.1  zeroc0d3_kafka   # container_name
  ```

* Install Dependencies
  ```
  make setup
  ```

* Build Producer
  ```
  make build-producer
  ```

* Build Consumer
  ```
  make build-consumer
  ```

* Run Producer
  ```
  make run-producer
  ```

* Run Consumer
  ```
  make run-consumer
  ```

### Copyright
* Author : **Dwi Fahni Denni (@zeroc0d3)**
* License: **Apache ver-2**