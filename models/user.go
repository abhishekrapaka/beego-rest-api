package models

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	//"golang.org/x/crypto/bcrypt"

	"github.com/beego/beego/v2/client/orm"
	"golang.org/x/crypto/bcrypt"
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AddUser(u User) string {
	o := orm.NewOrm()
	fmt.Println(u)
	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	//UserList[u.Id] = &u
	u.Password, _ = HashPassword(u.Password)
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

		var username, password, gender, email, address string
		var age int

		if uu.Username != "" {
			username = uu.Username
		} else {
			username = user.Username
		}
		if uu.Password != "" {
			password = uu.Password
		} else {
			password = user.Password
		}
		if uu.Age != 0 {
			age = uu.Age
		} else {
			age = user.Age
		}
		if uu.Address != "" {
			address = uu.Address
		} else {
			address = user.Address
		}
		if uu.Gender != "" {
			gender = uu.Gender
		} else {
			gender = user.Gender
		}
		if uu.Email != "" {
			email = uu.Email
		} else {
			email = user.Email
		}
		_, err := o.Raw("update user set username = " + "'" + username + "'" + ", password =" + "'" + password + "'" + ", Age =" + strconv.Itoa(age) + ", address = " + "'" + address + "'" + ", email = " + "'" + email + "'" + ", gender = " + "'" + gender + "'" + " where id = " + "'" + uid + "'").Exec()

		if err != nil {
			return nil, errors.New("some problem in db")
		}

		o.Read((&user))

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
		if CheckPasswordHash(password, user.Password) {
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
