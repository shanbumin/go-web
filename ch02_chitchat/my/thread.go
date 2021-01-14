package my

import (
	"log"
	"time"
)


type Thread struct {
	Id        int
	Uuid      string  //12a43db6-b8b9-4845-6151-21b57a4973c6
	Topic     string  //30岁之前必须得结婚吗
	UserId    int     //1
	CreatedAt time.Time //2019-02-11 13:21:14.725847
}

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

//格式化帖子的创建时间
func (thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}
//格式化回复记录的创建时间
func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

//获取某个帖子的回复数量
func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("SELECT count(*) FROM posts where thread_id = ?", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}












// get posts to a thread
func (thread *Thread) Posts() (posts []Post, err error) {
	rows, err := Db.Query("SELECT id, uuid, body, user_id, thread_id, created_at FROM posts where thread_id = ?", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

// Create a new thread
func (user *User) CreateThread(topic string) (conv Thread, err error) {
	conv.Uuid=createUUID()
	conv.UserId=user.Id
	conv.CreatedAt=time.Now()
	conv.Topic=topic
	rs,err:=Db.Exec("insert into threads (uuid, topic, user_id, created_at) values (?,?,?,?)",conv.Uuid,conv.Topic,conv.UserId,conv.CreatedAt)
	if err!=nil{
		log.Fatalln(err)
		return
	}
	id,err:=rs.LastInsertId()
	if err!=nil{
		log.Fatalln(err)
		return
	}
	conv.Id=int(id)
	return
}

// Create a new post to a thread
func (user *User) CreatePost(conv Thread, body string) (post Post, err error) {
	post.Uuid=createUUID()
	post.Body=body
	post.UserId=user.Id
	post.ThreadId=conv.Id
	post.CreatedAt=time.Now()
	rs,err:=Db.Exec("insert into posts (uuid, body, user_id, thread_id, created_at) values (?, ?, ?, ?,?)",post.Uuid,post.Body,post.UserId,post.ThreadId,post.CreatedAt)
	if err !=nil{
		log.Fatalln(err)
		 return
	}
	id,err:=rs.LastInsertId()
	if err !=nil{
		log.Fatalln(err)
		return
	}
	post.Id=int(id)
	return
}


// 从数据库里面按照创建时间取出所有帖子并将其返回
func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Thread{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt); err != nil {
			return
		}
		threads = append(threads, conv)
	}
	rows.Close()
	return
}

// Get a thread by the UUID
func ThreadByUUID(uuid string) (conv Thread, err error) {
	conv = Thread{}
	err = Db.QueryRow("SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid = ?", uuid).
		Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}

// Get the user who started this thread
//{{ .User.Name }}  模板中使用这个动作会找到这里的
func (thread *Thread) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = ?", thread.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

// Get the user who wrote the post
func (post *Post) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = ?", post.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}
