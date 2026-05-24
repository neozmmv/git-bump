$BASE_URL = "https://github.com/neozmmv/git-bump/releases/latest/download"
$BINARY = "git-bump.exe"
$INSTALL_DIR = "$env:USERPROFILE\AppData\Local\Microsoft\WindowsApps"

Write-Host "Downloading git-bump..."
Invoke-WebRequest -Uri "$BASE_URL/$BINARY" -OutFile "$INSTALL_DIR\$BINARY"

Write-Host "Installed! Run: git bump patch"