package model

//gorm libreria para ORM
type User struct {
	Id       int    `gorm:"primaryKey"`                 //clave
	Name     string `gorm:"type:varchar(350);not null"` //restricciones
	LastName string `gorm:"type:varchar(250);not null"`
	UserName string `gorm:"type:varchar(150);not null;unique"` //unique: no se puede repetir y no es PK
	Password string `gorm:"type:varchar(150);not null"`        //deberia estar encriptado
}

type Users []User
