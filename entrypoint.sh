#!/bin/sh
go run db/migrate/migrate.go  # マイグレーションを実行
exec go run main.go           # アプリケーションを実行