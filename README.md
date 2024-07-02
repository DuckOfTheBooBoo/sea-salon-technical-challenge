# SEA Salon

> This repository contains the codebase for the SEA Salon, a full-stack web application developed for the COMPFEST Software Engineering Academy Technical Challenge.

## Features
- **Authentication**
  - Login
  - Sign Up
- **Customer**
  - Reservation
  - Review form (customers can add reviews about SEA Salon)
- **Admin**
  - Add new branch
  - Modify existing branch
  - Delete a branch

## Deployment
### Requirements
- Docker (with Docker Compose)

### Steps

1. **Clone this repository**

   Open your terminal and clone the repository:
   ```bash
   git clone https://github.com/DuckOfTheBooBoo/sea-salon-technical-challenge
   cd sea-salon-technical-challenge

2. Modify the credentials in `docker-compose.yaml` (optional)
```yaml
name: sea-salon
services:
  db:
    image: mysql
    environment:
      MYSQL_ROOT_PASSWORD: password # change this
      MYSQL_DATABASE: sea_salon

  backend:
    build:
      context: ./backend
      dockerfile: Dockerfile
    environment:
      DB_NAME: sea_salon
      DB_USER: root
      DB_PASS: password # change this
      DB_HOST: db
      MAX_RETRIES: 10 # attempts to connect to database on startup
      RETRY_DELAY: 10 # seconds before retrying
    depends_on:
      - db

  frontend:
    build:
      context: ./frontend
      dockerfile: Dockerfile
    ports:
      - "8080:80"
```

3. Build and run the container

Use Docker Compose to build and run the conainers in detached mode.
```bash
docker compose up -d
```

4. Check the logs

```bash
docker compose logs -f
```

Monitor the logs to ensure all services start correctly:
- Wait for the backend service to successfully connect to the database. This may take a few moments as MySQL needs time to initialize.

- Look for the log entry [GIN-debug] Listening and serving HTTP on :3000. This indicates that the backend service is running successfully.

5. Access the application

Open your web browser and navigate to http://localhost:8080 (or your desired port) to use the SEA Salon web application.


### NOTE
- If the backend logs show `connect: connection refused`, it means the backend is trying to connect to the database before it is ready. The backend will retry connecting based on the MAX_RETRIES and RETRY_DELAY environment variables.

- Ensure that the docker-compose.yaml file has the correct credentials and configurations as per your setup requirements.

- Additionally, if the backend stop retrying after the amount of attempts has reached, you could just restart the container, the name usually goes by `sea-salon-backend-1` (kindly check your container's name `docker ps -a`!)
```bash
docker container restart sea-salon-backend-1
```
