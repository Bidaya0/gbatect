package converter

import (
	"github.com/compose-spec/compose-go/loader"
	"testing"
)

var dockercomposefile = []byte(`version: '3'

services:

  jobaggrv-mysql:
    image: mysql:8-debian
    extra_hosts:
      - "somehost:162.242.195.82"
      - "otherhost:50.31.209.229"
    environment:
      MYSQL_ROOT_PASSWORD: jobaggrv
      MYSQL_USER: jobaggrv
      MYSQL_PASSWORD: jobaggrv
      MYSQL_DATABASE: jobaggrv
    ports:
      - 33060:3306
    cap_add:
      - ALL
    cap_drop:
      - NET_ADMIN
      - SYS_ADMIN
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost"]
      interval: 1m30s
      timeout: 10s
      retries: 3
      start_period: 40s

  jobaggrv-scrapy:
    build:
      dockerfile: ./Dockerfile
      context: .
      target: prod
      args:
        GIT_COMMIT: cdc3b19
    command: echo "I'm running ${COMPOSE_PROJECT_NAME}"
    devices:
      - "/dev/ttyUSB0:/dev/ttyUSB0"
      - "/dev/sda:/dev/xvda:rwm"
    init: true
    shm_size: 1024kb
    working_dir: /opt

  jobaggrv-api:
    build:
      dockerfile: ./Dockerfile
      context: .
      args:
        - GIT_COMMIT=cdc3b19
    entrypoint: gunicorn  --worker-class gevent -w 4 --worker-connections 1000 -b 0.0.0.0:8000 'jobaggrv_api.app:app'
    ports: 
      - 8000:8000
    depends_on:
      - jobaggrv-scrapy
    logging:
      driver: syslog
      options:
        syslog-address: "tcp://192.168.0.42:123"
`)

func TestLoadServices(t *testing.T) {
	k1, _ := loader.ParseYAML(dockercomposefile)
	tmpk := k1["services"]
	tmp3, _ := tmpk.(map[string]interface{})
	_, err := LoadServices(tmp3)
	if err != nil {
		t.Error("Expected success, got ", err)
	}
	return
}

func TestTransServicesToContainer(t *testing.T) {
	k1, _ := loader.ParseYAML(dockercomposefile)
	tmpk := k1["services"]
	tmp3, _ := tmpk.(map[string]interface{})
	services, err := LoadServices(tmp3)
	_, err = TransServicesToContainer(services)
	if err != nil {
		t.Error("Expected success, got ", err)
	}
	return
}
