This repo has some gorp tips.

- unit test with mysql.
  - store and reset data in each cases.
- execute sql written in a go-template.

# Try

## Setup

```
$ docker-compose up # -d
$ docker exec -it gorp /bin/bash
# cd src/
```

## Exec
```
~/src# go run main_without_template.go
~/src# go run main_with_template.go
~/src# go run main_with_template_in_bin.go
```

## Test

```
~/src# go test repositories/jet_repository_test.go
ok      command-line-arguments  0.030s
```
