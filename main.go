package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=db user=postgres password=passw0rd dbname=postgres sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// INSERT
	user := User{
		Name:    "Fujiwara Takumi",
		RoleIDs: RoleIDs{1, 2, 3},
		Resume: Resume{
			Summary: "Hogehoge",
			Experiences: []string{
				"Experience 1",
				"Experience 2",
			},
			Skills: []string{
				"Skill 1",
				"Skill 2",
			},
		},
	}
	result := db.Create(&user)
	if result.Error != nil || result.RowsAffected != 1 {
		log.Println("create user error:", result.Error)
		return
	}
	insertedID := user.ID

	// SELECT
	var selectedUser User
	if err := db.First(&selectedUser, insertedID).Error; err != nil {
		log.Println("select user error:", err)
		return
	}

	log.Printf("%+v\n", selectedUser)
}

type User struct {
	gorm.Model
	Name    string
	RoleIDs RoleIDs
	Resume  Resume
}

type RoleID int

type RoleIDs []RoleID

type Resume struct {
	Summary     string
	Experiences []string
	Skills      []string
}

// RoleIDs Value Marshal
func (ids RoleIDs) Value() (driver.Value, error) {
	return json.Marshal(ids)
}

// RoleIDs Scan Unmarshal
func (ids *RoleIDs) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("User: type assertion to []byte failed")
	}
	return json.Unmarshal(b, &ids)
}

// Resume Value Marshal
func (resume Resume) Value() (driver.Value, error) {
	return json.Marshal(resume)
}

// Resume Scan Unmarshal
func (resume *Resume) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("Resume: type assertion to []byte failed")
	}
	return json.Unmarshal(b, &resume)
}
