#!/usr/bin/env pwsh
Set-StrictMode -Version Latest

Write-Host "Downloading modules..."
go mod download

if (-not (Test-Path bin)) { New-Item -ItemType Directory bin }

Write-Host "Building binary..."
go build -o .\bin\myapp.exe .\cmd\app\main.go

