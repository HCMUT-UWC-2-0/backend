# Gin template

## Before start the project

- Name of module is placed in `go.mod`, you need rename it.
- Fix all imports in all files with "re-named" module
- This project use **viper** for loading configs, make sure that you checked them
  - `app.env`
  - `.docker.env`
  - `docker-compose.dev`
  - `docker-compose`
  - `Makefile`
- It uses **sqlc** tool for generating **.go** files at `db/sqlc` directory. It makes easier for interacting with database layer.
  > Just add sql scripts in `db/query` directory, each file corresponding one entity.
- It uses **paseto** for token generation, you can add another token generator that you want, it has an interface named **Maker** at `token/maker.go`

## Project structure

```null
📦 api                      # Business layer
 ┣ 📜 server.go             # Init routing
 ┗ 📜 [sample].go         # Sample API   
 ┣ 📜 middleware.go         # Middleware
 ┗ 📜 validator.go          # Validator for Gin
 ┣ 📜 main_test.go          # Test all apis
📦 db                       # Database layer
 ┣ 📂 initdb                  # Init sql (optional)
 ┃ ┣ 📂 func_proc             # function and procedure for database
 ┃ ┣ 📂 trigger               # trigger for database
 ┃ ┗ 📜 [sample].sql             
 ┣ 📂 migration              # Contains migration files
 ┣ 📂 mock                   # Mock db (auto generating) 
 ┣ 📂 query                  # Define entities here
 ┃ ┗ 📜 [sample].sql               
 ┣ 📂 sqlc                   # Database caller (auto-generating)
 ┃ ┣ 📜 db.go, models.go, querier.go  # (auto-generating)          
 ┃ ┣ 📜 [sample].go                   # (auto-generating)               
 ┃ ┗ 📜 store.go             # Database storage (tx logics)
 ┣ 📂 token                  # token
 ┣ 📂 utils                  
 ┣ 📜 main.go                # Init sever and db, start here
 ┣ 📜 .docker.env            
 ┗ 📜 app.env              
```

## How to start this project?

1. Install dependencies

  ```bash
  go get -u -v  all
  ```

2. Init postgres

  ```bash
  make bootstrap
  ```

3. Init schema

  ```bash
  make migrateup
  ```

4. Create sample data data (optional)

  ```bash
  make seed
  ```

5. Start the app:

  ```bash
  make server
  ```

6. Access [adminer](http://localhost:8080/) and login follows info in **.docker.env** file.

## Some helpful commands

- Drop all db:
  `make migratedown`
- Sql command line
  `make psql`

## After added new entities, you must run these commands for re-generating **models**

```bash
make sqlc && make mock
```

## Simple rules when using Git

1. For development, you have to create a new branch like this: `feat/your_feature`. Eg: `feat/customer`
2. Before merging your branch to `main`,

- Ensure that your ticket passes the definition of done
- Check that you’ve added the necessary tests
- Finally, create pull requests. ([ref here](https://docs.github.com/en/desktop/contributing-and-collaborating-using-github-desktop/working-with-your-remote-repository-on-github-or-github-enterprise/creating-an-issue-or-pull-request))

3. Commitment (`git commit -m`) follows this format: `feature|fix(branch_name): your message`.

- Eg: `feature(customer): new function`, `fix(customer): fix leak memory`,...

4. Before coding, you have to make sure that the source code is the latest version, use `git checkout main && git pull main`.

## Merge request locally

1. `git fetch origin`

2. `git checkout -b origin/feature feature`

3. `git fetch origin`

4. `git checkout origin/main`

5. `git merge --no-ff feature`

6. `git push origin main`
