// Golangはクラス内の変数とメソッドを別にもつ
// クラス内変数 -> struct
// メソッドの集合 -> interface

package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //匿名フィールド
	school string
	loan   float32
}

type Employee struct {
	Human   //匿名フィールド
	company string
	money   float32
}

//HumanにSayHiメソッドを実装します。
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//HumanにSingメソッドを実装します。
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

//EmployeeはHumanのSayHiメソッドをオーバーロードします。
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

// Interface MenはHuman,StudentおよびEmployeeに実装されます。
// この３つの型はこの２つのメソッドを実装するからです。
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	//Men型の変数iを定義します。
	var i Men

	//iにはStudentを保存できます。
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	//iにはEmployeeを保存することもできます。
	i = tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	//sliceのMenを定義します。
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	//この３つはどれも異なる要素ですが、同じインターフェースを実装しています。
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}
}
