$dir = 'internal/model/dto'
$types = @('int','int64','string','float64')
$total = 0
Get-ChildItem $dir -Filter *.go | ForEach-Object {
    $file = $_.FullName
    $text = [System.IO.File]::ReadAllText($file, [System.Text.Encoding]::UTF8)
    $lines = $text -split "`n"
    $inReq = $false
    $count = 0
    for ($i = 0; $i -lt $lines.Count; $i++) {
        $l = $lines[$i]
        if ($l -match '^\s*//\s*@request') { $inReq = $true; continue }
        if ($l -match '^\s*//\s*@response') { $inReq = $false; continue }
        if ($l -match '^\s*// =+') { $inReq = $false; continue }
        if ($inReq -and $l -match '^(\s*[A-Z]\w*\s+)(' + ($types -join '|') + ')(\s+`)') {
            $t = $matches[2]
            $lines[$i] = $l -replace ('^(\s*[A-Z]\w*\s+)(' + ($types -join '|') + ')(\s+`)'), ('$1*' + $t + '$3')
            $count++
        }
    }
    if ($count -gt 0) {
        $newText = $lines -join "`n"
        [System.IO.File]::WriteAllText($file, $newText, (New-Object System.Text.UTF8Encoding($false)))
        $total += $count
        Write-Host ($_.Name + ' changed=' + $count)
    }
}
Write-Host "TOTAL_CHANGED=$total"
