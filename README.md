## About
fpl-find-a-manager is a simple tool that collects manager's data from official Fantasy Premier League API, and allows you to find teams of your favorite managers.

## Prerequisites
* Docker https://docs.docker.com/get-docker/
* Mage https://github.com/magefile/mage

## Building
```sh
mage build cli
```

## Running
```sh
docker-compose -f docker-compose-dev.yaml up -d
./app
```

## License
fpl-find-a-manager is licensed under MIT License. See LICENSE file.