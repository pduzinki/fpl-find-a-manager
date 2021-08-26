## About
*fpl-find-a-manager* is a little tool that collects manager's data from official Fantasy Premier League API, and allows you to find teams of your favorite managers.

## Prerequisites
* Go (1.16 or higher) https://golang.org/doc/install
* Docker https://docs.docker.com/get-docker/
* Docker Compose https://docs.docker.com/compose/install/
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
*fpl-find-a-manager* is licensed under MIT License. See LICENSE file.
