package utils

import (
	"regexp"
	"strings"
)

// EmailRegex เป็นรูปแบบ Regular Expression สำหรับตรวจสอบอีเมลเบื้องต้น
var EmailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.(com|net|org|info)$`)
var plateFormat = regexp.MustCompile(`^[ก-ฮ]{1,2}\s[0-9]{1,4}$`)

// IsValidEmail ตรวจสอบว่ารูปแบบอีเมลถูกต้องหรือไม่
func IsValidEmail(email string) bool {
	if email == "" {
		return false
	}
	return EmailRegex.MatchString(strings.ToLower(email))
}

// IsValidID ตรวจสอบว่า ID ที่ส่งมาเป็นค่าบวก (> 0) หรือไม่
func IsValidID(id int) bool {
	return id > 0
}

// IsPlateValid ตรวจสอบรูปแบบทะเบียนรถเบื้องต้น (เช่น ต้องไม่ว่าง และมีความยาวที่เหมาะสม)
func IsPlateValid(plate string) bool {
	// 1. นำช่องว่างหน้าหลังออก
	trimmedPlate := strings.TrimSpace(plate)

	// 2. ตรวจสอบความว่างเปล่าหลัง Trim (ถ้าว่างให้ return false ทันที)
	if len(trimmedPlate) == 0 {
		return false
	}

	// 3. ตรวจสอบรูปแบบด้วย Regexp
	return plateFormat.MatchString(trimmedPlate)
}

// IsRequired ตรวจสอบว่า string ที่ส่งมาไม่ใช่ค่าว่างหลังจาก trim space แล้ว
func IsRequired(input string) bool {
	return strings.TrimSpace(input) != ""
}
