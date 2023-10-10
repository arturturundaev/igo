package model

type User struct {
	Id          string `db:"id"`
	First_name  string `db:"first_name"`
	Second_name string `db:"second_name"`
	Middle_name string `db:"middle_name"`
}
