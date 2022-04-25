package data

import (
	"log"

	"github.com/objectbox/objectbox-go/objectbox"
	dErr "github.com/timreimherr/dhelp/internal/errors"
	"github.com/timreimherr/dhelp/internal/model"
)

func initObjectBox() *objectbox.ObjectBox {
	objectBox, err := objectbox.NewBuilder().Model(model.ObjectBoxModel()).Build()
	if err != nil {
		log.Fatal(err)
	}

	return objectBox
}

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

func AddInfoToSection(sectionID uint64, key string, value string) error {
	db := initObjectBox()
	defer db.Close()

	box := model.BoxForSection(db)

	section, err := box.Get(sectionID)
	if err != nil {
		return err
	}

	if section == nil {
		return dErr.ErrSectionDoesNotExist
	}

	infoBox := model.BoxForInfo(db)

	_, err = infoBox.Put(&model.Info{
		Section: section,
		Key:     key,
		Value:   value,
	})

	return err
}

func GetInfos() []*model.Info {
	db := initObjectBox()
	defer db.Close()

	infoBox := model.BoxForInfo(db)

	infos, _ := infoBox.Query(model.Info_.Section_Id.OrderAsc()).Find()

	return infos
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
