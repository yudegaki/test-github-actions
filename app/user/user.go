package user

type User struct {
	Id   int
	Name string
	Age  int
}

func (u *User) SetId(id int) {
	u.Id = id
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) SetAge(age int) {
	u.Age = age
}

func GenerateUser(id int, name string, age int) (u User) {
	u.SetId(id)
	u.SetName(name)
	u.SetAge(age)

	return u
}
