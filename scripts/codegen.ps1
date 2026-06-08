#!/usr/bin/env pwsh
Set-StrictMode -Version Latest

$oapi = Join-Path (go env GOPATH) 'bin\oapi-codegen'
if (-not (Test-Path $oapi)) {
    Write-Host "oapi-codegen not found in GOPATH/bin. Please run: go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest"
    exit 1
}

& "$oapi" --config .\openapi\oapi-codegen.yaml .\openapi\openapi.yaml

