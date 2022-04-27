package service

import (
	"fmt"

	"github.com/timreimherr/jhelp/internal/data"
	"github.com/timreimherr/jhelp/internal/log"
	"github.com/timreimherr/jhelp/internal/model"
)

func printSection(sections []*model.Section) {
	fmt.Println()
	for _, s := range sections {
		infos := data.GetInfosBySectionId(s.Id)
		log.Section(s.Name)
		for _, i := range infos {
			log.Info(i.Key, i.Value)
		}
		fmt.Println()
	}
}

func PrintAll() {
	sections := data.GetSections()
	printSection(sections)
}

func PrintSection(name string) {
	section := data.GetSectionByName(name)
	printSection(section)
}
