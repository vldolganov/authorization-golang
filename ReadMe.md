### Docker compose start up
- Copy the environment keys from `cmd/.example.env` to `cmd/.env`
- Set values in .env file (default values are acceptable)
- If the database name in .env has been changed, then `./sql/dump.sql` must be changed

Create build
```bash
$ docker-compose build
```

Run build
```bash
$ docker-compose up
```
