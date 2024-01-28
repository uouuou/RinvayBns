SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=386
SET GCC=C:\mingw64\bin\gcc.exe
SET CXX=C:\mingw64\bin\g++.exe

set Version=%date:~0,4%%date:~5,2%%date:~8,2%%h%%time:~3,2%%time:~6,2%
set APP_VERSION=BSN%Version%
::set APP_VERSION=ShuHeCrmDll202312202812
rsrc -ico ./utils/icon.ico -o bns_x86.syso -manifest .\bns_x86.exe.mainfest
rsrc -ico ./utils/icon.ico -o bns_amd64.syso -manifest .\bns_amd64.exe.mainfest

echo now the CGO_ENABLED:
 go env CGO_ENABLED

echo now the GOOS:
 go env GOOS

echo now the GOARCH:
 go env GOARCH
  go build -ldflags  "-H windowsgui -s -w -X 'bns/utils.Version=%APP_VERSION%'" -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o bns_x86.exe .\main.go
upx bns_x86.exe

SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=amd64


echo now the CGO_ENABLED:
 go env CGO_ENABLED

echo now the GOOS:
 go env GOOS

echo now the GOARCH:
 go env GOARCH
  go build -ldflags  "-H windowsgui -s -w -X 'bns/utils.Version=%APP_VERSION%'" -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o bns_amd64.exe .\main.go
upx bns_amd64.exe

echo %APP_VERSION%