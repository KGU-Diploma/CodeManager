# CodeManager

# How to run service
- clone repo and go to project dir
```bash
git clone git@github.com:KGU-Diploma/CodeManager.git
cd CodeManager
```

- install dependencies 
```bash
go mod download
```

 - go to cmd dir and execute this script
```bash
go run main.go
 ```

 - call api for get liveness probe
  ```bash
 curl --location 'http://localhost:8001/health/ping'
 ```