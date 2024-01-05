
# Dota2 Scheduler

Dota 2 Scheduler is a pet project based on retrieving information from Dota servers about the current leaderboard of the best players and providing this information to users from the database

## API Reference (coming soon)

All API routes are available in Swagger documentation

#### Get all items

```http
  GET /api/swagger
```




## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`HOST`

`PORT`

`DRIVERNAME`

`DATASOURCENAME`

`SQL="./sql/init.sql"`

SQL cannot be changed in any way. it is required for creating all tables.

## Run Locally

To run this project locally, you need the following:

Clone the repo from github

```bash
  git clone https://github.com/tgkzz/dota2-scheduler.git
```

Go to the project directory

```bash
  cd dota2-scheduler
```

Create database using sqlite

```bash
  sqlite3 DatabaseName.db
```

After completing all the steps, you can run your project.

```bash
  go run . **config_path**
```

By default, the config is specified in the root directory with the file name `.env`
