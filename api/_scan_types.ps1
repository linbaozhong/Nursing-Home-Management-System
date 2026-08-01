$dir = 'internal/model/dto'
Get-ChildItem $dir -Filter *.go | ForEach-Object {
    $lines = Get-Content -Path $_.FullName -Encoding UTF8
    $inReq = $false
    for ($i = 0; $i -lt $lines.Count; $i++) {
        $l = $lines[$i]
        if ($l -match '^\s*//\s*@request') { $inReq = $true; continue }
        if ($l -match '^\s*//\s*@response') { $inReq = $false; continue }
        # entering a struct resets
        if ($l -match '^\s*type\s+\w+\s+struct') { if (-not $inReq) { continue } }
        if ($inReq -and $l -match '^\s*([A-Z]\w*)\s+(\*?\w+(?:\.\w+)?)\s+`') {
            $fname = $matches[1]; $ftype = $matches[2]
            if ($ftype -notmatch '^\*') {
                Write-Host ($_.Name + ' :: ' + $ftype)
            }
        }
        # leaving request block when a new comment-group or non-struct appears
        if ($inReq -and $l -match '^\s*//\s*@') { }
        if ($inReq -and $l -match '^\s*// =+') { $inReq = $false }
    }
}
