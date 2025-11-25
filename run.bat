@echo off
setlocal

echo Building Go application...

:: 1. Build โค้ด (สร้าง main.exe ในโฟลเดอร์ tmp)
:: ใช้คำสั่งที่พิสูจน์แล้วว่าทำงานได้
go build ./cmd/api
go build -o ./tmp/main.exe ./cmd/api

:: ตรวจสอบว่า Build สำเร็จหรือไม่
if %errorlevel% neq 0 (
echo.
echo ❌ BUILD FAILED! Please check your Go code for errors.
exit /b %errorlevel%
)

echo.
echo ✅ Build successful.
echo Running application from tmp\main.exe...
echo ---------------------------------------

:: 2. Run โค้ดที่ Build แล้ว
tmp\main.exe

endlocal