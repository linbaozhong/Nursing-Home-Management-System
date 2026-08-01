$dir = 'internal/service'
$re1 = '^\s*db := DB\(\)\s*$'
$re2 = '^\s*_ = db\s*$'
$total1 = 0; $total2 = 0
Get-ChildItem $dir -Filter *.go | ForEach-Object {
    $file = $_.FullName
    $text = [System.IO.File]::ReadAllText($file, [System.Text.Encoding]::UTF8)
    $lines = $text -split "`n"
    $out = @()
    $skipNext = $false
    for ($i = 0; $i -lt $lines.Count; $i++) {
        $l = $lines[$i]
        if ($l -match $re1) { $total1++; continue }
        if ($l -match $re2) { $total2++; continue }
        $out += $l
    }
    if ($total1 -gt 0 -or $total2 -gt 0) {
        $newText = $out -join "`n"
        [System.IO.File]::WriteAllText($file, $newText, (New-Object System.Text.UTF8Encoding($false)))
        Write-Host ($_.Name + ' removed db:=DB()' + $total1 + ' _=db' + $total2)
    }
}
Write-Host "TOTAL db:=DB()=$total1 ; _=db=$total2"
