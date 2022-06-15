package htime

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//======================== OFFSET VALIDATION  =======================

func Offset_validation() int {

	offset := "00:00:00"
	offset_int := 0
	offset_int = offset_int + 0
	var sele string
	choice := ""
	re := regexp.MustCompile(`^(\d\d):(\d\d):(\d\d)$`)

	for choice != "s" {

		fmt.Println("Please enter the time offset ( in this format 00:00:00 ) if it exists: ")
		fmt.Scanln(&offset)

		valid_offset := re.MatchString(offset)

		if valid_offset {

			offset = strings.Replace(offset, ":", "", -1)

			offset_int = Clock_to_calc(offset)
			choice = "s"
		} else {

			fmt.Printf("Invalid offset format %s", offset)
			fmt.Println(" press \"s\" to exit another key to continue")
			fmt.Scanf("%s", &sele)
			if sele == "s" {
				os.Exit(0)

			}
		}
	}
	return offset_int
}

//====================== SPLIT H:M.S ================================
// Recive "isosec" parameter, split it and return: hour, min & sec

func Clock_to_calc(t string) int {

	if t != "" {
		prev_sec := t                           // 20220522170124
		sec := prev_sec[len(prev_sec)-2:]       //             24
		prev_min := prev_sec[:len(prev_sec)-2]  // 202205221701
		min := prev_min[len(prev_min)-2:]       //           01
		prev_hour := prev_min[:len(prev_min)-2] // 2022052217
		hour := prev_hour[len(prev_hour)-2:]    //         17

		time_format := All_in_seconds(hour, min, sec)

		return time_format

	} else {
		return 0
	}
}

//===================== H:M:S TO SEC ================================

func All_in_seconds(h string, m string, s string) int {

	htos, _ := strconv.Atoi(h)
	mtos, _ := strconv.Atoi(m)
	sec, _ := strconv.Atoi(s)

	htos = htos * 3600
	mtos = mtos * 60

	total_sec := htos + mtos + sec

	return total_sec
}

//==================== SECONDS TO H:M:S =============================

func Sec_to_HMS(inSec int) string {
	seconds := inSec % 60
	minutes := (inSec / (60)) % 60
	hours := inSec / 3600

	str := fmt.Sprintf("%d:%d:%d", hours, minutes, seconds)

	return str

}

//===================================================================
//===================================================================
