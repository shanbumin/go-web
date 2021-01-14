package my

import (
	"log"
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

//为已经存在的用户创建一个新的session
func (user *User) CreateSession() (session Session, err error) {

	session.Uuid=createUUID()
	session.Email=user.Email
	session.UserId=user.Id
	session.CreatedAt=time.Now()
	rs, err :=Db.Exec("insert into sessions (uuid, email, user_id, created_at) values (?,?,?,?)",session.Uuid,session.Email,session.UserId,session.CreatedAt)
	if err != nil {
		log.Fatalln(err)
	}
	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	session.Id=int(id)
	return
}

// Get the session for an existing user
func (user *User) Session() (session Session, err error) {
	session = Session{}
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = ?", user.Id).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

//检测当前客户端携带的cookie，在数据库中是否存在有效的session
func (session *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = ?", session.Uuid).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}



//session表的uuid删除对应的session记录
func (session *Session) DeleteByUUID() (err error) {
	statement := "delete from sessions where uuid = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(session.Uuid)
	return
}

//根据session信息获取登录的用户信息
func (session *Session) User() (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = ?", session.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

// Delete all sessions from database
func SessionDeleteAll() (err error) {
	statement := "delete from sessions"
	_, err = Db.Exec(statement)
	return
}

//注册用户的时候，需要添加进来
func (user *User) Create() (err error) {

	user.Uuid=createUUID()
	user.Password=Encrypt(user.Password)
	user.CreatedAt=time.Now()

	rs,err:=Db.Exec("INSERT INTO users (uuid,name,email,password,created_at) values (?,?,?,?,?)",user.Uuid,user.Name,user.Email,user.Password,user.CreatedAt)
	if err !=nil{
		 log.Fatalln(err)
	}
	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}
	user.Id=int(id)
	return
}




// Delete user from database
func (user *User) Delete() (err error) {
	statement := "delete from users where id = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	return
}

// Update user information in the database
func (user *User) Update() (err error) {
	statement := "update users set name = $2, email = $3 where id = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id, user.Name, user.Email)
	return
}

// Delete all users from database
func UserDeleteAll() (err error) {
	statement := "delete from users"
	_, err = Db.Exec(statement)
	return
}

// Get all users in the database and returns it
func Users() (users []User, err error) {
	rows, err := Db.Query("SELECT id, uuid, name, email, password, created_at FROM users")
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

//根据邮箱获取用户信息
func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = ?", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// Get a single user given the UUID
func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE uuid = ?", uuid).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}
