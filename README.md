## Learn Kafka
Learning kafka using Go and library shopify sarama

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