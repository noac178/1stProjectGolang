package pkg

import "strings"

func createPagePath(cate string) string {
	var cate_eng string
	replacer := strings.NewReplacer(" - ", "-", " ", "-")
	cate_eng = replacer.Replace(cate)

	return strings.ToLower(cate_eng)
}
