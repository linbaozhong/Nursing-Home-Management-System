$dir = 'internal/model/dto'
$types = @('int','int64','string','float64')
Get-ChildItem $dir -Filter *.go | ForEach-Object {
    $lines = Get-Content -Path $_.FullName -Encoding UTF8
    $inReq = $false
    for ($i = 0; $i -lt $lines.Count; $i++) {
        $l = $lines[$i]
        if ($l -match '^\s*//\s*@request') { $inReq = $true; continue }
        if ($l -match '^\s*//\s*@response') { $inReq = $false; continue }
        if ($l -match '^\s*// =+') { $inReq = $false; continue }
        if ($inReq -and $l -match '^(\s*[A-Z]\w*\s+)(' + ($types -join '|') + ')(\s+`)') {
            $t = $matches[2]
            $new = $l -replace '^(\s*[A-Z]\w*\s+)(' + ($types -join '|') + ')(\s+`)', ('$1*' + $t + '$3')
            Write-Host ($_.Name + ' | ' + $l.Trim() + ' => ' + $new.Trim())
        }
    }
}
Write-Host 'PREVIEW DONE'
