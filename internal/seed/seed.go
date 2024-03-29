package seed

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
	"hp-hotel-rest/internal/model"
)

func Load(db *gorm.DB) {
	if err := db.AutoMigrate(&model.Hotel{}, &model.Location{}); err != nil {
		log.Fatalf("Unable to migrate: %v", err)
	}

	var count int64
	db.Model(&model.Hotel{}).Count(&count)
	if count > 0 {
		return
	}

	amenities := []model.Amenity{
		{Name: "Wi-Fi"},
		{Name: "Havuz"},
		{Name: "Spa"},
		{Name: "Fitness Merkezi"},
		{Name: "Ücretsiz Kahvaltı"},
		{Name: "Restoran"},
	}

	for _, ame := range amenities {
		db.Create(&ame)
	}

	locations := []model.Location{
		{City: "Antalya", District: "Kumluca", Address: "Kumluca No 1. Antalya"},
		{City: "Antalya", District: "Belek", Address: "Belek No 1. Antalya"},
		{City: "Antalya", District: "Merkez", Address: "Merkez No 1. Antalya"},
	}

	for _, loc := range locations {
		db.Create(&loc)
		hotels := []model.Hotel{
			{Name: gofakeit.Company() + " Hotel", Stars: gofakeit.Number(0, 5), Type: "Şehir Oteli", Rating: 0, LocationID: loc.ID},
			{Name: gofakeit.Company() + " Hotel", Stars: gofakeit.Number(0, 5), Type: "Resort", Rating: 0, LocationID: loc.ID},
			{Name: gofakeit.Company() + " Hotel", Stars: gofakeit.Number(0, 5), Type: "Tatil Köyü", Rating: 0, LocationID: loc.ID},
		}

		for _, hotel := range hotels {
			db.Create(&hotel)

			for i := 1; i < gofakeit.Number(2, 10); i++ {
				review := model.Review{
					UserName:  gofakeit.Name(),
					UserEmail: gofakeit.Email(),
					Text:      gofakeit.LoremIpsumSentence(10),
					Rating:    gofakeit.Number(0, 10),
				}
				db.Model(&hotel).Association("Reviews").Append(&review)
			}

			for i := 1; i < gofakeit.Number(2, 10); i++ {
				photos := model.Photo{
					URL: gofakeit.URL(),
				}
				db.Model(&hotel).Association("Photos").Append(&photos)
			}

			roomTypes := []string{"Tek Kişilik", "Çift Kişilik", "Superior", "Suit"}

			for i := 1; i < gofakeit.Number(2, 4); i++ {
				index := gofakeit.Number(0, len(roomTypes)-1)
				roomType := roomTypes[index]

				roomTypes = append(roomTypes[:index], roomTypes[index+1:]...)

				room := model.Room{
					Name:  roomType,
					Price: gofakeit.Price(500, 2500),
				}

				db.Model(&hotel).Association("Rooms").Append(&room)
			}

			var selectedAmenities []model.Amenity
			db.Find(&selectedAmenities, "name IN (?)", []string{"Wi-Fi", "Havuz", "Spa"})

			for _, amenity := range selectedAmenities {
				db.Model(&hotel).Association("Amenities").Append(&amenity)
			}

		}

	}

	log.Debugw("Mock data loaded successfully")
}
