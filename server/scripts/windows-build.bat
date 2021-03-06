@echo off

CD /d %~dp0\..\

SET PATH=%GAE_GO_SDK_HOME%;%GAE_GO_ROOT%;%GAE_GO_ROOT%\bin;%GAE_PYTHON_PATH%;%PATH%
SET GOROOT=%GAE_GO_ROOT%
SET GOPATH=%CD%\.gopath\windows;%CD%


REM Setup Environment
SET WORKSPACE=%CD%\src
SET GCP_PROJECT_ID=%TIMEBOOSTER_PROJECT_ID%

REM echo PATH=%PATH%

echo CD=%CD%
echo GOROOT=%GOROOT%
echo GOPATH=%GOPATH%
echo GCP_PROJECT_ID=%GCP_PROJECT_ID%
echo WORKSPACE=%WORKSPACE%

echo TIMEBOOSTER_API_KEY=%TIMEBOOSTER_API_KEY%
echo TIMEBOOSTER_ENDPOINT=%TIMEBOOSTER_ENDPOINT%
echo TIMEBOOSTER_PROJECT_ID=%TIMEBOOSTER_PROJECT_ID%

cd src
call goapp build ./...
