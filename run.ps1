# run.ps1
$script:process = $null

# Запускаем процесс
Write-Host "Run searcher.exe..." -ForegroundColor Green
$script:process = Start-Process -FilePath ".\searcher.exe" -PassThru -NoNewWindow

Write-Host "Searcher started. PID: $($process.Id)" -ForegroundColor Cyan
Write-Host "Close this window, for stop searcher.exe" -ForegroundColor Yellow

# Ждём закрытия процесса
try {
    $process.WaitForExit()
    Write-Host "Searcher.exe is done." -ForegroundColor Gray
}
catch {
    Write-Host "Error: $_" -ForegroundColor Red
}