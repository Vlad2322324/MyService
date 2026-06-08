#!/usr/bin/env pwsh
Set-StrictMode -Version Latest

if (-not $env:DATABASE_URL) {
    Write-Host "DATABASE_URL not set. Example: $env:DATABASE_URL = 'postgres://postgres:password@localhost:5432/mydb?sslmode=disable'"
    exit 1
}

if (-not (Get-Command migrate -ErrorAction SilentlyContinue)) {
    Write-Host "migrate tool not found. Install: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"
    exit 1
}

migrate -path .\migrations -database $env:DATABASE_URL up

