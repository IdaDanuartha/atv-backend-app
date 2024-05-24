package models

// Customer Model
type Customer struct {
	Base
	Name 	string `gorm:"size:100" json:"name"`
	UserID  string    `gorm:"primaryKey;foreignKey:UserID"`
  	User    User   `gorm:"foreignKey:UserID"`
}

// TableName method sets table name for Customer model
func (Customer *Customer) TableName() string {
	return "customers"
}

// ResponseMap -> response map method of Customer
func (Customer *Customer) ResponseMap() map[string]interface{} {
	resp := make(map[string]interface{})
	resp["id"] = Customer.ID
	resp["name"] = Customer.Name
	resp["created_at"] = Customer.CreatedAt
	resp["updated_at"] = Customer.UpdatedAt
	resp["deleted_at"] = Customer.DeletedAt

	return resp
}
