# leetcode_tournament
Соревнование внутри алема на платформе литкод

### Install `migrate` command-tool:
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

### Create new migration:
```
migrate create -ext sql -dir migrations mg_name
```

### Apply migration:
```
migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/leetcode?sslmode=disable" up
```


### Dependencies 
```azure
.env
```

### Run project
```
docker build -t leetcode .
docker run -p 3000:3000 leetcode
```