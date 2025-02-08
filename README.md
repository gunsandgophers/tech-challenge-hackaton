# Tech Challenge Hackaton

Process videos into snapshots

## Architecture

![tech-challenge-hackaton](https://github.com/user-attachments/assets/a5b32f10-7b41-4561-a33b-d11b4893a9fd)

## Run project
To run the application it is necessary to execute the command ```make start```

### Migration
All migrations are executed as soon as the make start or make build command is executed

#### Create
To create a migration, you need to run the make migrate/create command passing the file name

**example:**

```make migrate/create name=add_user```
to create a migration to add a user

### Swagger
URL to access running Swagger is ```/api/v1/swagger/index.html```
