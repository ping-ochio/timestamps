package list

import (
	"fmt"
	"strings"
	"timestamps/htime"
)

//========================== FINAL LIST =============================

func Final_list(isomatch []string, clean_raw string, offset_int int) {

	first_pos := htime.Clock_to_calc(isomatch[0]) - offset_int
	pos_video := make([]string, 0)

	for i := range isomatch {
		pos := htime.Clock_to_calc(isomatch[i]) - first_pos
		pos_video = append(pos_video, htime.Sec_to_HMS(pos))
	}

	for i, value := range isomatch {

		clean_raw = strings.Replace(clean_raw, value, pos_video[i], 1)
	}

	fmt.Println(clean_raw)
}
