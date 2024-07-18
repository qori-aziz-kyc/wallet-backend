# wallet-backend

## Migrations

Written in cmd/api/migration/db/migrations
use [golang-migrate](https://github.com/golang-migrate) as CLI

How to install

- brew install golang-migrate (Macos)
- Create a migration:

  `migrate create -ext sql -dir ./cmd/migrations -seq {filename}`

- run migration
  `migrate -path ./cmd/migrations -database mysql://"{username}:{password}@tcp({host}:3306)/wallet?charset=utf8mb4&parseTime=True&loc=Local" up`
