package data

import (
	"errors"
	"log"
	"os"

	"github.com/objectbox/objectbox-go/objectbox"
	dErr "github.com/timreimherr/jhelp/internal/errors"
	"github.com/timreimherr/jhelp/internal/model"
)

func init() {
	_, err := os.Open("./storage/data.mdb")
	if errors.Is(err, os.ErrNotExist) {
		createDefaultRecords()
	}
}

func initObjectBox() *objectbox.ObjectBox {
	objectBox, err := objectbox.NewBuilder().Directory("./storage").Model(model.ObjectBoxModel()).Build()
	if err != nil {
		log.Fatal(err)
	}

	return objectBox
}

// Sections

func CreateSection(name string) error {
	db := initObjectBox()
	defer db.Close()

	box := model.BoxForSection(db)

	sections, err := box.Query(model.Section_.Name.Equals(name, false)).Find()
	if err != nil {
		log.Fatal(err)
	}

	if len(sections) > 0 {
		return dErr.ErrSectionAlreadyExists
	}

	_, err = box.Put(&model.Section{
		Name: name,
	})

	return err
}

func GetSectionByName(name string) []*model.Section {
	db := initObjectBox()
	defer db.Close()

	box := model.BoxForSection(db)

	sections, err := box.Query(model.Section_.Name.Equals(name, false)).Find()
	if err != nil {
		log.Fatal(err)
	}

	return sections
}

func GetSectionById(id uint64) []*model.Section {
	db := initObjectBox()
	defer db.Close()

	box := model.BoxForSection(db)

	sections, err := box.Query(model.Section_.Id.Equals(id)).Find()
	if err != nil {
		log.Fatal(err)
	}

	return sections
}

func GetSections() []*model.Section {
	db := initObjectBox()
	defer db.Close()

	box := model.BoxForSection(db)

	sections, err := box.Query(model.Section_.Id.OrderAsc()).Find()
	if err != nil {
		log.Fatal(err)
	}

	return sections
}

// Infos

func AddInfoToSection(sectionID uint64, key string, value string) (sectionName string, err error) {
	db := initObjectBox()
	defer db.Close()

	box := model.BoxForSection(db)

	section, err := box.Get(sectionID)
	if err != nil {
		return "", err
	}

	if section == nil {
		return "", dErr.ErrSectionDoesNotExist
	}

	infoBox := model.BoxForInfo(db)

	_, err = infoBox.Put(&model.Info{
		Section: section,
		Key:     key,
		Value:   value,
	})

	return section.Name, err
}

func GetInfosBySectionId(sectionId uint64) []*model.Info {
	db := initObjectBox()
	defer db.Close()

	infoBox := model.BoxForInfo(db)

	infos, err := infoBox.Query(model.Info_.Section_Id.Equals(sectionId), model.Info_.Id.OrderAsc()).Find()
	if err != nil {
		log.Fatal(err)
	}

	return infos
}

func GetInfos() []*model.Info {
	db := initObjectBox()
	defer db.Close()

	infoBox := model.BoxForInfo(db)

	infos, _ := infoBox.Query(model.Info_.Section_Id.OrderAsc()).Find()

	return infos
}
