$dir = 'internal/service'
$re1 = '^\s*db := DB\(\)\s*$'
$re2 = '^\s*_ = db\s*$'
$total1 = 0; $total2 = 0
Get-ChildItem $dir -Filter *_service.go | ForEach-Object {
    $lines = Get-Content -Path $_.FullName -Encoding UTF8
    for ($i = 0; $i -lt $lines.Count; $i++) {
        if ($lines[$i] -match $re1) { $total1++; if ($total1 -le 5) { Write-Host ($_.Name + ' [L' + ($i+1) + '] ' + $lines[$i].Trim()) } }
        if ($lines[$i] -match $re2) { $total2++ }
    }
}
Write-Host "db:=DB()=$total1 ; _ = db=$total2"
