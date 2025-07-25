package main

import "fmt"

// type Residence struct {
//  Province string
//  City     string
// }

type User struct {
	// 成员变量
	Id				int
	Score			float64
	Name, address	string
	Father			*User	//结构体嵌套自身会形成无限循环,但是可以嵌套自己的指针,因为指针只是一个地址,不会产生循环嵌套
	residence		struct {
		Province	string
		City		string
	}
}

type TreeNode struct {
	Data		int
	LeftChild	*TreeNode
	RightChild	*TreeNode
}

// 成员方法
func (mi User) hello() {
	fmt.Println("My name is", mi.Name)
}

func main() {
	var u User
	u.Score = 100
	u = User{Id: 32, address: "中国", Name: "乔瑟夫"}
	fmt.Println(u.Name)
	u = User{32, 34.9, "斯黛拉", "中国", nil,
		struct {
			Province string
			City	 string
		}{},
	}
	u.hello()

	ue := u
	ue.Name = "张三"
	_ = ue

	// 匿名结构体(该结构体只使用一次)
	var student struct {
		Name string
		Age	 int
	}
	student.Name = "Joseph"
	student.Age = 18
	u.Name = student.Name
	u2 := User{}
	u2.address = "上海"

	u3 := &u2 // 取址符号,u3是*User类型
	fmt.Println(u3.Name)

	u4 := new(User) // new先创建空的结构体,再返回其指针,u4是*User类型
	u4.Name = "Jodie"
	u4.Father = u3
	fmt.Printf("%+v\n", u4)
	fmt.Println(u4.Father.Name)

	u5 := User{
		Name: "Tom",
		residence: struct {
			Province string
			City	 string
		}{Province: "广东", City: "广州"},
	}
	u5.residence.Province = "湖南"
	u5.residence.City = "长沙"
	_ = u5
}
