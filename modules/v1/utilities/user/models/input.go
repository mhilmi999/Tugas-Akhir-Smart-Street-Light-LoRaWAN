package models

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}
type RegisterInput struct {
	Nama     string `json:"nama" form:"nama" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type Received struct{
	First AntaresDetail `json:"m2m:cin"`
}

type AntaresDetail struct{
	Rn  string `json:"rn"`
	Ty  int    `json:"ty"`
	Ri  string `json:"ri"`
	Pi  string `json:"pi"`
	Ct  string `json:"ct"`
	Lt  string `json:"lt"`
	St  int    `json:"st"`
	Cnf string `json:"cnf"`
	Cs  int    `json:"cs"`
	Con DataReceived `json:"con"`
}

type DataReceived struct{
	Voltage string `json:"voltage"`
	Current string `json:"current"`
	Power   string `json:"power"`
}