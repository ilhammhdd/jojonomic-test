this repo is a technical test for Jojonomic Backend Go(Golang) position.

## required tools to review this repo
1. visual studio code
2. docker
3. insomnia
4. DB manager, e.g. DBeaver

## fire up
1. clone repo
2. open repo as workspace in visual studio code
3. run script ./misc/setup.sh
4. run script ./misc/docker_compose_up.sh
5. run go mod tidy and go mod vendor for each services
6. run debug for each services except misc_test
7. open insomnia and import ./misc/Insomnia-Jojonomic-Test-Backend-Go_2022-12-04.json
