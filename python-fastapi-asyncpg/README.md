# Python + FastAPI + asyncpg Template

## Getting started

Requirements:
- docker
- poetry
- httpie

install dependencies:

```sh
poetry install
```

build images:

```sh
docker compose build
```

run containers locally:

```sh
docker compose up
```

verify server is running by hitting the status endpoint:

```sh
http GET http://localhost:8000/status
```

Run DB migrations:
```sh
docker compose exec app alembic upgrade head
```

## Try out the "items" example API

POST an items:

```sh
http POST http://127.0.0.1:8000/api/items item:='{"name": "foo", "price": 3.14}'
```

GET a single item:

```sh
http GET http://127.0.0.1:8000/api/items/1
```

GET multiple items:

```sh
http GET 'http://127.0.0.1:8000/api/items' item_ids==1 item_ids==2
```

### Running the docker containers will spin-up Swagger docs and Adminer

- Visit Swagger docs here:

    ```sh
    http://127.0.0.1:8000/docs
    ```

- Visit Adminer DB management tool here:

    ```sh
    http://127.0.0.1:8080/?pgsql=db&username=user&db=example_db&ns=public
    ```

## Database migrations

Alembic is used to manage raw SQL migrations. Migrations are not automatically
run when doing local development, but _are_ run automatically when a production
container is started.

The process for creating new migration files, dry-run testing, and application
of a new migration is as follows.

make sure you have the containers up and running in a terminal tab:
```sh
docker compose up
```

open a poetry shell:

```sh
poetry shell
```

create a new migration file:

```sh
alembic revision -m "my new migration"
# and also write your `upgrade` and `downgrade` queries
```

dry run your migration:
```sh
docker compose exec app alembic upgrade head --sql
```

apply your new migration:
```sh
docker compose exec app alembic upgrade head
```

(optional) roll-back your migration:
```sh
docker compose exec app alembic downgrade -1
```

## Running loadtests

Locust is used for loadtesting. The python file for managing tests is [here](./locustfile.py)

run a loadtest method or set of methods by tag:
```sh
poetry run locust -f locustfile.py --headless --tag=<my_tag> \
--users=<num_users> --spawn-rate=<spawn_rate> --run-time=<run_duration> \
--host=http://localhost:8000
```

run all loadtests:
```sh
make loadtest
```

## Other dev commands

See the [Makefile](./Makefile)

if you get isort errors, run the command alone to fix:

```sh
poetry run isort .
```
