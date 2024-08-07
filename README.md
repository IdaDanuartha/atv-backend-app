# ATV Backend System with Golang

An simple app for booking atv ticket


## Run Locally

Clone the project

```bash
  git clone https://github.com/IdaDanuartha/atv-backend-app.git
```

Go to the project directory

```bash
  cd atv-backend-app
```

Duplicate .env.example to .env

```bash
  cp .env.example .env
```

Adjust the .env configuration according to your settings
```bash
  APP_NAME="ATV System"
  APP_PORT=8000
  APP_PREFIX=/api/v1
  APP_URL=http://localhost:${APP_PORT}${APP_PREFIX}

  DB_USERNAME=root
  DB_PASSWORD=
  DB_HOST=localhost
  DB_PORT=3306
  DB_DATABASE=atv_system

  SECRET_KEY=atv_system_s3cr3T_k3Y
```

Install golang

```bash
  Download & Install golang https://go.dev/doc/install
```

Check the golang version

```bash
  go version
```

Start the server

```bash
  go run main.go
```

