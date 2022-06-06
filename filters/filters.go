package filters

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"timestamps/htime"
	"timestamps/list"
)

//======================== FILTER GIT ===============================

func Filter_github(str string) {

	fmt.Println()
	fmt.Println("**************** GITHUB ***************")
	fmt.Println("***************************************")
	section := regexp.MustCompile(`START[\S\s]*\(\d*\)`)
	block := section.FindAllString(str, -1)

	rep := strings.NewReplacer(
		"<li>", "-   ",
		"</li>", "\r",
		"<p dir=\"auto\">", "",
		"<ul dir=\"auto\">", "",
		"<ul>", "",
		"</ul>", "",
		"<code>", "",
		"</code>", "",
		"</p>", "",
	)
	clean_raw := rep.Replace(block[0])

	var validID = regexp.MustCompile(`[[:digit:]]{14}`)
	isomatch := validID.FindAllString(clean_raw, -1)

	offset := htime.Offset_validation()
	list.Final_list(isomatch, clean_raw, offset)

}

//======================== FILTER YOUTUBE ===========================

func Filter_youtube(str string) {

	fmt.Println()
	fmt.Println("**************** YOUTUBE ***************")
	fmt.Println("****************************************")

	first := regexp.MustCompile(`descriptionBodyText":{(\S\w.*)\((\d{14})\)`)
	second := regexp.MustCompile(`START.*`)
	sa := first.FindAllString(str, -1)
	if sa == nil {
		fmt.Println()
		fmt.Println("*** It seems that the information we are looking for is not found ***")
		fmt.Println()
		os.Exit(0)
	}
	sb := second.FindAllString(sa[0], -1)

	rep := strings.NewReplacer(
		"\\n\\n", "\n\n",
		"\\n", "\n",
	)
	clean_raw := rep.Replace(sb[0])

	var validID = regexp.MustCompile(`[[:digit:]]{14}`)
	isomatch := validID.FindAllString(clean_raw, -1)

	offset := htime.Offset_validation()
	list.Final_list(isomatch, clean_raw, offset)

}

//===================================================================
//===================================================================
