# URL Shortener from Free Mastering Go With GoLand Course
## Build
```code
 go run  cmd/web-app/main.go
```
Add some urls says it

1.```https://github.com```

2.```https://Github.com```


You will see the generate hash and for looking to the SQLite3 database you will have the same result

## SQLite3

```code
sqlite3 db.sqlite 

sqlite> select * from urls ;
1|996e1f71|https://github.com
2|47032296|https://Github.com
3|f9190e43|https://
```
