#!/usr/bin/env pwsh
Set-StrictMode -Version Latest

function Run-Command($cmd) {
    Write-Host "$cmd"
    $res = & pwsh -Command $cmd
    return $LASTEXITCODE
}

if (-not (Get-Command git -ErrorAction SilentlyContinue)) {
    Write-Error "git not found in PATH. Install git and retry."
    exit 1
}

$branch = "feat/split-commits"
Write-Host "Creating branch: $branch"
git checkout -b $branch

function CommitFiles($files, $message, $body) {
    Write-Host "Staging: $files"
    git add $files
    if (-not (git diff --cached --quiet)) {
        Write-Host "Committing: $message"
        git commit -m "$message" -m "$body"
        if ($LASTEXITCODE -ne 0) {
            Write-Error "git commit failed"
            exit 1
        }
        # optional quick build/test
        Write-Host "Running go build..."
        & go build ./... 2>&1 | Write-Host
    } else {
        Write-Host "No changes to commit for: $message"
    }
}

CommitFiles "pkg/logger/logger.go cmd/app/main.go" "feat(logger): add zap-based structured logger with LOG_LEVEL support" "Introduce pkg/logger using zap; select production/dev config based on ENV and allow setting explicit LOG_LEVEL (debug/info/warn/error). Initialize logger in main and defer Sync()."

CommitFiles "internal/middleware/request_id.go" "feat(middleware): add X-Request-ID middleware and register in server" "Add middleware to generate/preserve X-Request-ID, set header and put request_id into Echo context. Wire middleware in main and include request_id in request logs."

CommitFiles "internal/repositories/user_repository.go" "fix(repo): correct column name and use pointer for DB connection; improve logging" "Fix SQL column naming (createdat -> created_at), change NewDatabaseConn to accept *pgx.Conn, and update repository logging to log errors as Error and successes as Debug."

CommitFiles "internal/handlers/user_handler/user_handler.go" "feat(handler): add structured error logging and remove debug prints" "Handlers now log internal errors using pkg/logger before returning safe HTTP responses to clients. Removed stray fmt/print debug statements."

CommitFiles "migrations/0001_create_users.up.sql migrations/0001_create_users.down.sql scripts/migrate.ps1" "feat(migrations): add initial create_users migration and migrate script" "Add initial SQL migration to create users table and PowerShell script to run migrations using golang-migrate. Encourage running migrations in CI and staging before app start."

CommitFiles "scripts/codegen.ps1 scripts/build.ps1 scripts/run.ps1" "feat(scripts): add PowerShell codegen/build/run helpers" "Provide scripts for Windows development: codegen (oapi-codegen), build (go build), and run (tee logs). These make it easy to work without GNU make on Windows."

CommitFiles ".github/workflows/ci.yml" "ci: add GitHub Actions workflow (codegen, vet, test, build)" "Add CI workflow to run oapi-codegen, go vet, go test and build on push/pull_request to main/master. Ensures generated code alignment and basic checks in CI."

CommitFiles ".ai-factory/ARCHITECTURE.md .ai-factory/RULES.md .ai-factory/SKILLS.md .env.example" "docs(onboarding): add AI Factory architecture, rules, skills and .env.example" "Add documentation in .ai-factory/ describing architecture, rules and required skills for contributors/agents. Add .env.example for local dev env variables."

# Final catch-all commit for any remaining changes
git add -A
if (-not (git diff --cached --quiet)) {
    git commit -m "chore(repo): remaining changes" -m "Catch-all commit for any leftover modifications after thematic commits."
}

Write-Host "All commits created on branch: $branch"
Write-Host "To push branch: git push -u origin $branch"

