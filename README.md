## Dependencies
- docker
- kind 
- kubectl
- kustomize

## Git Commit Message
https://www.conventionalcommits.org/en/v1.0.0/

# Not Complete
## Setup
get docker image for postgres
```
docker pull vincent0426/tuber-postgres
```
1. create k8s cluster with Kind
```sh
make dev-up
```
2. create docker image for main server named `tuber/tuber-api`
```sh
make service
```
3. load tuber/tuber-api to kind container
```sh
make dev-load
```
4. 
```sh
make dev-apply
```

if update code only
```sh
make dev-update
```
if update config and code
```sh
make dev-update-apply
```

## Database

### Migrations
```sh
make db-migrations-up

make db-migrations-down
```

### Seed
```sh
make db-seed-up

make db-seed-down
```