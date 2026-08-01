$dir = 'internal/model/dto'
Get-ChildItem $dir -Filter *.go | ForEach-Object {
    $lines = Get-Content -Path $_.FullName -Encoding UTF8
    $inReq = $false
    for ($i = 0; $i -lt $lines.Count; $i++) {
        $l = $lines[$i]
        if ($l -match '^\s*//\s*@request') { $inReq = $true; continue }
        if ($l -match '^\s*//\s*@response') { $inReq = $false; continue }
        if ($l -match '^\s*// =+') { $inReq = $false; continue }
        if ($inReq -and $l -match '^\s*([A-Z]\w*)\s+(\*?\w+(?:\.\w+)?)\s+`') {
            $ftype = $matches[2]
            if ($ftype -notmatch '^\*') { Write-Host ($_.Name + ' :: ' + $ftype) }
        }
    }
}
