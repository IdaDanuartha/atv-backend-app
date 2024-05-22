package models

// Mandatory Luggage Model
type MandatoryLuggage struct {
	Base
	Name string `gorm:"size:100" json:"name"`
}

// TableName method sets table name for Bus model
func (MandatoryLuggage *MandatoryLuggage) TableName() string {
	return "mandatory_luggages"
}

// ResponseMap -> response map method of Entertainment Category
func (MandatoryLuggage *MandatoryLuggage) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = MandatoryLuggage.ID
	resp["name"] = MandatoryLuggage.Name
	resp["created_at"] = MandatoryLuggage.CreatedAt
	resp["updated_at"] = MandatoryLuggage.UpdatedAt
	resp["deleted_at"] = MandatoryLuggage.DeletedAt

	return resp
}
