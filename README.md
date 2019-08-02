#c4-portfolio

```sh
$   CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/c4-portfolio
```

```sh
$   docker build --no-cache -t c4-portfolio .
```

```sh
$   docker run -d --name c4-portfolio -p 8080:8080 c4-portfolio
```