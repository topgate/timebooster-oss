@echo off

CD /d %~dp0\..\

SET GOPATH=%CD%\.gopath\windows
SET PATH=%MINGW64_PATH%\bin;%PATH%

echo CD=%CD%
echo GOROOT=%GOROOT%
echo GOPATH=%GOPATH%
echo TIMEBOOSTER_API_KEY=%TIMEBOOSTER_API_KEY%
echo TIMEBOOSTER_ENDPOINT=%TIMEBOOSTER_ENDPOINT%
echo TIMEBOOSTER_PROJECT_ID=%TIMEBOOSTER_PROJECT_ID%

SET EXE_NAME="timebooster"

echo "########################"
echo "## Build for Linux"
echo "########################"
SET GOOS=linux
SET GOARCH=amd64

go build -o %EXE_NAME%
