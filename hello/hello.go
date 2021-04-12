package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"sync"
	"testing"
	"time"
)

func GetData() (int, int) {

	return 100, 200
}

//使用指针更改值
func modify(sls *[3]int) {
	(*sls)[0] = 90
}

//if else
func testif(age int) {
	if age := 20; age > 10 {
		fmt.Printf("age is : %d", age)

	} else {
		fmt.Printf("he is not ")
	}
}

//for循环
func getResult(args ...int) bool {
	for _, i := range args {
		if i < 60 {
			return false
		}
	}
	return true
}

//if语句
func set_data(x int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var arr [10]int
	arr[x] = 88
}

//结构体加方法
type Profile struct {
	name   string
	age    int
	gender string
	mother *Profile
	father *Profile
}

func (person *Profile) FmtProfile() {
	person.age += 1
}

//结构体的继承
type company struct {
	companyName string
	companyAddr string
}

type staff struct {
	name     string
	age      int
	gender   string
	position string
	company
}

//接口
type Good interface {
	settleAccount() int
	orderInfo() string
}

/* 结构体实现 */
type Phone struct {
	name     string
	quantity int
	price    int
}

/* 结构体实现 */
type FreeGift struct {
	name     string
	quantity int
	price    int
}

//为结构体加方法
func (phone Phone) settleAccount() int {
	return phone.quantity * phone.price
}

func (phone Phone) orderInfo() string {
	return "您要购买" + strconv.Itoa(phone.quantity) + "个" + phone.name + "计：" + strconv.Itoa(phone.settleAccount()) + "元"
}

func (gift FreeGift) settleAccount() int {
	return 0
}
func (gift FreeGift) orderInfo() string {
	return "您要购买" + strconv.Itoa(gift.quantity) + "个" +
		gift.name + "计：" + strconv.Itoa(gift.settleAccount()) + "元"

}

//定义一个添加到购物车的方法
func calculateAllPrice(goods []Good) int {
	var allPrice int
	for _, good := range goods {
		fmt.Println(good.orderInfo())
		allPrice += good.settleAccount()
	}
	return allPrice
}

type Person1 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr string `json:"addr,omitempty"`
}

//使用tag
func getTag() {
	p1 := Person1{
		Name: "Jack",
		Age:  22,
	}

	data1, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data1)

	p2 := Person1{
		Name: "Jack",
		Age:  22,
		Addr: "China",
	}

	data2, err := json.Marshal(p2)
	if err != nil {
		panic(err)
	}

	// p2 则会打印所有
	fmt.Printf("%s\n", data2)

}

//获取tag
func getTagName() {
	p2 := Person1{
		Name: "Jack",
		Age:  22,
		Addr: "China",
	}
	field, _ := reflect.TypeOf(p2).FieldByName("Name")
	tag := field.Tag
	fmt.Println(tag)

}

//使用tag
type Person struct {
	Name   string `label:"Name is: "`
	Age    int    `label:"Age is: "`
	Gender string `label:"Gender is: " default:"unknown"`
}

//tag练习
func Print(obj interface{}) error {
	// 取 Value
	v := reflect.ValueOf(obj)

	// 解析字段
	for i := 0; i < v.NumField(); i++ {

		// 取tag
		field := v.Type().Field(i)
		tag := field.Tag

		// 解析label 和 default
		label := tag.Get("label")
		defaultValue := tag.Get("default")

		value := fmt.Sprintf("%v", v.Field(i))
		if value == "" {
			// 如果没有指定值，则用默认值替代
			value = defaultValue
		}

		fmt.Println(label + value)
	}

	return nil
}

// type switch
func findType(i interface{}) {
	switch x := i.(type) {
	case string:
		fmt.Println(x, "is string")
	case int:
		fmt.Println(x, "is int")
	default:
		fmt.Println(x, "is not matched")
	}
}

func getType() {
	var i interface{} = 10
	t1, status := i.(int)

	fmt.Println(t1, "-", status)

	fmt.Println("=============")

	t2, status1 := i.(string)

	fmt.Println(t2, "_", status1)

	findType(status)
}

//调用函数的隐式转换
type Phoneq interface {
	call()
}

type iPhone struct {
	name string
}

func (phone iPhone) call() {
	fmt.Println("Hello, iPhone.")
}

func (phone iPhone) send_wechat() {
	fmt.Println("Hello, Wechat.")
}

func howToGetSend() {
	phone := iPhone{name: "ming's iphone"}
	phone.call()
	phone.send_wechat()
}

//
func reflectTest() {
	var age4 int = 23
	fmt.Printf("类型为：%T, 值为：%v", age4, age4)
	t := reflect.TypeOf(age4)
	v := reflect.ValueOf(&age4)

	fmt.Printf("\ntype对象类型为：%T\n", t)
	fmt.Printf("value对象类型为：%T\n", v)
	fmt.Println("是否可写：\n", v.CanSet())
	v2 := v.Elem()
	v2.SetInt(12)
	fmt.Println(age4)
}

func test1() {
	//定义一个值
	// var lean interface{} = "xishi"
	m := Person{}
	//反射,基础类型
	t := reflect.TypeOf(m)

	fmt.Println("Type: ", t)
	fmt.Println("Kind: ", t.Kind())

	v := reflect.ValueOf(&m)

	fmt.Println("\nxxxx Type: ", v.Type())
	fmt.Println("Kind: \n", v.Kind())

	fmt.Println("Type: ", v.Elem().Type())
	fmt.Println("Kind: ", v.Elem().Kind())
}

func floatToFloat() {
	var score float64 = 99.9
	a := reflect.ValueOf(score)
	a1 := a.Float()
	fmt.Printf("%T is :%v", a1, a1)
	fmt.Println(a1)
}

func slicetest() {
	var numList []int = []int{1, 2, 3, 4}

	v1 := reflect.ValueOf(numList)
	fmt.Printf("转换前， type: %T, value: %v \n", v1, v1)

	// Slice 函数接收两个参数
	v2 := v1.Slice(0, 2)
	fmt.Printf("转换后， type: %T, value: %v \n", v2, v2)
}

type Person4 struct {
	name string
	age  int
}

func (p Person4) SayBye() string {
	return "bye"
}

func (p Person4) SayHello() string {
	return "hello"
}

func useMethod() {
	p := &Person4{"写代码的明哥", 27}

	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	// 也可以这样来遍历
	for i := 0; i < v.NumMethod(); i++ {
		fmt.Printf("第 %d 个方法：%v , 调用结果： %v\n", i+1, t.Method(i).Name, v.Elem().Method(i).Call(nil))
	}
}

type pers struct {
}

func (p pers) SelfIntroduction(name string, age int) {
	fmt.Printf("Hello, my name is %s and i'm %d years old.", name, age)
}

func sum(args ...int) int {
	var result int
	for _, v := range args {
		result += v
	}
	return result
}
func Sum(args ...int) int {
	fmt.Println(args)
	result := sum(args...)
	return result

}

//协程测试
func testgo(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println("hello go ", name)
		time.Sleep(10 * time.Millisecond)
	}
}

func testTowGo() {
	pipline := make(chan int)
	go func() {
		fmt.Println("发送数据 10")
		pipline <- 100
	}()
	go func() {
		num := <-pipline
		fmt.Printf("接受数据： %d", num)
	}()
	time.Sleep(time.Second)
}

// type Sender = chan<- int
// type Receiver = <-chan int

func fibonacci(mychan chan int) {
	n := cap(mychan)
	x, y := 1, 1
	for i := 0; i < n; i++ {
		mychan <- x
		x, y = y, x+y
	}
	close(mychan)
}

func increment(ch chan bool, x *int) {
	ch <- true
	*x = *x + 1
	<-ch
}

func worker(x int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("worker %d :%d \n", x, i)

	}
}

func add(count *int, wg *sync.WaitGroup, lock *sync.Mutex) {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		*count = *count + 1
		lock.Unlock()
	}
	wg.Done()
}

func lockAndUlock() {
	lock := &sync.RWMutex{}
	lock.Lock()

	for i := 0; i < 4; i++ {
		go func(i int) {
			fmt.Printf("the %d start \n", i)
			lock.RLock()
			fmt.Printf("the %d get the lock , sleep 1s , release the lock\n", i)
			time.Sleep(time.Second)
			lock.RUnlock()
		}(i)

	}

	time.Sleep(time.Second * 2)
	fmt.Println("prepare to release wlock , wlock is not in loading")
	lock.Unlock()

	lock.Lock()
	fmt.Println("the process exit")
	lock.Unlock()

}

func canNotExit() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("exit")
			default:
				fmt.Println("ing")

			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("is ok , notice stop ")

	stop <- true
	time.Sleep(time.Second * 5)

}

func monitor(ch chan bool, num int) {
	for {
		select {
		case v := <-ch:
			fmt.Printf("监控器%v ， 接受通道值%v \n", num, v)
			return
		default:
			fmt.Printf("监控器 %v， 正在监控\n", num)
			time.Sleep(2 * time.Second)

		}
	}
}

var inter time.Duration

// func init() {
// 	flag.DurationVar(&inter, "inter", 1* time.Second, "间隔")

// }

type TestTable struct {
	x int
	y int
}

func TestAdd(t *testing.T) {
	tables := []TestTable{
		{1, 2},
		{3, 4},
		{5, 6},
		{7, 8},
	}

	for _, i := range tables {
		res := Add(i.x, i.y)
		if res == (i.x + i.y) {
			t.Log("ok")

		} else {
			t.Fatal("wrong")
		}

	}
}

