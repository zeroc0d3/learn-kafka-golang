# -----------------------------------------------------------------------------
#  MAKEFILE RUNNING COMMAND
# -----------------------------------------------------------------------------
#  Author     : Dwi Fahni Denni (@zeroc0d3)
#  License    : Apache version 2
# -----------------------------------------------------------------------------
# Notes:
# use [TAB] instead [SPACE]

export PROJECT_NAME="learn-kafka-golang"
export PATH_SCRIPTS="./scripts/installer/"
export PATH_PROJECT="./src"
export PATH_BINARY="./bin"
export PATH_CMD="./cmd"
export BIN_PUBLISHER="kafka-producer"
export BIN_CONSUMER="kafka-consumer"
export MSG_FILE="./words.txt"
export KAFKA_TOPIC="kafka-topic"


install-ansible:
	@echo '-----------------------------'
	@echo ' Install Ansible '
	@echo '-----------------------------'
	@cd ${PATH_SCRIPTS} && bash ./install_ansible.sh

install-docker:
	@echo '-----------------------------'
	@echo ' Install Docker '
	@echo '-----------------------------'
	@cd ${PATH_SCRIPTS} && bash ./install_docker.sh

install-pyenv:
	@echo '-----------------------------'
	@echo ' Install Python '
	@echo '-----------------------------'
	@cd ${PATH_SCRIPTS} && bash ./install_python_env.sh
	@python ${PATH_SCRIPTS}/get-pip.py

install-golang:
	@echo '-----------------------------'
	@echo ' Install Golang '
	@echo '-----------------------------'
	@cd ${PATH_SCRIPTS} && bash ./install_golang.sh
	@echo ' - DONE - '
	@echo ''

clean:
	@echo '-----------------------------'
	@echo ' Cleanup Old Binary '
	@echo '-----------------------------'
	@rm -f ${PATH_BINARY}/*
	@rm -rf __vendor
	@echo ' - DONE - '
	@echo ''

clean-vendor:
	@echo '-----------------------------'
	@echo ' Cleanup Old Vendor '
	@echo '-----------------------------'
	@rm -rf __vendor
	@echo ' - DONE - '
	@echo ''

init:
	@make clean-vendor
	@echo '-----------------------------'
	@echo ' Initialization Dependencies '
	@echo '-----------------------------'
	@cd ${PATH_PROJECT} && dep init
	@echo ' - DONE - '
	@echo ''

setup:
	@make clean
	@echo '-----------------------------'
	@echo ' Installation Dependencies '
	@echo '-----------------------------'
	@cd ${PATH_PROJECT} && dep ensure -v
	@echo ' - DONE - '
	@echo ''

build-producer:
	@echo '-----------------------------'
	@echo ' Building Binary Producer '
	@echo '-----------------------------'
	@go build -o ${PATH_BINARY}/${BIN_PUBLISHER} ${PATH_CMD}/producer/main.go
	@echo ' - DONE - '
	@echo ''

build-consumer:
	@echo '-----------------------------'
	@echo ' Building Binary Consumer '
	@echo '-----------------------------'
	@go build -o ${PATH_BINARY}/${BIN_CONSUMER} ${PATH_CMD}/consumer/main.go
	@echo ' - DONE - '
	@echo ''

run-producer:
	@echo '-----------------------------'
	@echo ' Running Binary Producer '
	@echo '-----------------------------'
	@cp ${MSG_FILE} ${PATH_BINARY}
	@chmod +x ${PATH_BINARY}/${BIN_PUBLISHER}
	@cd ${PATH_BINARY} && ./${BIN_PUBLISHER}
	@echo ' - DONE - '
	@echo ''

run-consumer:
	@echo '-----------------------------'
	@echo ' Running Binary Consumer '
	@echo '-----------------------------'
	@chmod +x ${PATH_BINARY}/${BIN_CONSUMER}
	@cd ${PATH_BINARY} && ./${BIN_CONSUMER}
	@echo ' - DONE - '
	@echo ''