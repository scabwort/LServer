@echo off

setlocal

if exist install.bat goto ok
echo install.bat must be run from its folder
goto end

: ok

set OLDGOPATH=%GOPATH%
set GOPATH=%GOPATH%;%~dp0

gofmt -w src

go install %1

cd bin
%1

:end
