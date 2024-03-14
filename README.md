# Clone Repository
```
https://github.com/dev-nichola/go-crud.git
```

# Set Up Mysql Table
```
// create database
CREATE database crud_go;

// use the database
use crud_go

// create table
create table employe (id int primary key auto_increment, name varchar(255), npwp varchar(255), address(255))

```
# Set Up Mysql Driver
```
go get -u github.com/go-sql-driver/mysql
```

# Run GO
```
go run main.go
```