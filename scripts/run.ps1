#!/usr/bin/env pwsh
Set-StrictMode -Version Latest

if (-not (Test-Path bin)) {
    Write-Host "Binary not found, building first..."
    .\build.ps1
}

if (-not (Test-Path logs)) { New-Item -ItemType Directory logs }

Write-Host "Running myapp.exe, logs -> .\\logs\\app.log"
.\bin\myapp.exe 2>&1 | Tee-Object -FilePath .\logs\app.log

