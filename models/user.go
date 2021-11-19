package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

var (
	UserList map[string]*User
)

func init() {
	orm.RegisterModel(new(User))

	// UserList = make(map[string]*User)
	// u := User{"user_11111", "astaxie", "11111", "male", 20, "Singapore", "astaxie@gmail.com"}
	// UserList["user_11111"] = &u

}

type User struct {
	Id       string `orm:"pk"`
	Username string
	Password string
	Gender   string
	Age      int
	Address  string
	Email    string
}

func AddUser(u User) string {
	o := orm.NewOrm()
	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	UserList[u.Id] = &u
	_, err := o.Insert(&u)
	if err != nil {
		return "User not created "
	} else {
		return u.Id
	}

}

func GetUser(uid string) (u *User, err error) {
	o := orm.NewOrm()
	user := User{Id: uid}

	rsp := o.Read(&user)
	fmt.Println(user)
	if rsp == orm.ErrNoRows {
		return nil, errors.New("User not exists")
	} else if rsp == orm.ErrMissPK {
		return nil, errors.New("User not exists")
	} else {
		return &user, nil
	}

}

func GetAllUsers() map[string]*User {
	o := orm.NewOrm()
	//var userslist map[string]*User
	var userslist map[string]*User = make(map[string]*User)
	var user []User
	result, _ := o.Raw(("select * from user")).QueryRows(&user)

	fmt.Println(result, user[0].Id)

	for _, u := range user {
		userslist[u.Id] = &User{
			Id:       u.Id,
			Username: u.Username,
			Password: u.Password,
			Gender:   u.Gender,
			Age:      u.Age,
			Address:  u.Address,
			Email:    u.Email,
		}
	}

	return userslist

}

func UpdateUser(uid string, uu *User) (a *User, err error) {

	o := orm.NewOrm()

	user := User{Id: uid}

	rsp := o.Read((&user))

	if rsp == orm.ErrNoRows {
		return nil, errors.New("User not exists")
	} else if rsp == orm.ErrMissPK {
		return nil, errors.New("User not exists")
	} else {

		_, err := o.Raw("update user set username = " + "'" + uu.Username + "'" + ", password =" + "'" + uu.Password + "'" + ", Age =" + "'" + strconv.Itoa(uu.Age) + "'" + ", address = " + "'" + uu.Address + "'" + ", email = " + "'" + uu.Email + "'" + ", gender = " + "'" + uu.Gender + "'" + " where id = " + "'" + uid + "'").Exec()

		if err != nil {
			return nil, errors.New("some problem in db")
		}
		return &user, nil
	}

	// if u, ok := UserList[uid]; ok {
	// 	if uu.Username != "" {
	// 		u.Username = uu.Username
	// 	}
	// 	if uu.Password != "" {
	// 		u.Password = uu.Password
	// 	}
	// 	if uu.Age != 0 {
	// 		u.Age = uu.Age
	// 	}
	// 	if uu.Address != "" {
	// 		u.Address = uu.Address
	// 	}
	// 	if uu.Gender != "" {
	// 		u.Gender = uu.Gender
	// 	}
	// 	if uu.Email != "" {
	// 		u.Email = uu.Email
	// 	}
	// 	return u, nil
	// }
	//return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {

	o := orm.NewOrm()

	var user User

	rsp := o.Raw("select password from user where username = " + "'" + username + "'").QueryRow(&user)

	fmt.Println(rsp, user)

	if rsp == orm.ErrNoRows {
		return false
	} else if rsp == orm.ErrMissPK {
		return false
	} else {
		if user.Password == password {
			return true
		}
	}

	// for _, u := range UserList {
	// 	if u.Username == username && u.Password == password {
	// 		return true
	// 	}
	// }
	return false
}

func DeleteUser(uid string) {
	o := orm.NewOrm()

	rsp, _ := o.Raw("delete from user where id = " + "'" + uid + "'").Exec()

	fmt.Println(rsp)
}
