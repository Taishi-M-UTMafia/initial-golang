package main

// ---------------------------------------------
// Cards
// ---------------------------------------------
// func main() {
// cards := newDeck()

// // receiver on function
// // cards.print()

// hand, remainingCards := deal(cards, 5)

// hand.print()
// remainingCards.print()

// cards.saveToFile("MyCards")

// cards := newDeckFromFile("MyCards")
// cards.print()

// cards := newDeck()
// cards.shuffle()
// cards.print()
// }

// ---------------------------------------------
// Structs
// ---------------------------------------------
// func main() {
// alex := person{firstName: "Alex", lastName: "Anderson"}
// var alex person

// alex.firstName = "Alex"
// alex.lastName = "Anderson"

// fmt.Println(alex)
// fmt.Printf("%+v", alex)

// jim := person{
// 	firstName: "Jim",
// 	lastName:  "Party",
// 	contactInfo: contactInfo{
// 		email:   "jimpgmail.com",
// 		zipCode: 94000,
// 	},
// }

// jimPointer := &jim
// jimPointer.updateName("Jimmy")

// updateNameのreceiverのtypeは*personな一方で
// jimの方はただのpersonだが、そこはgoがよしなにやってくれる
// jim.updateName("Jimmy")
// jim.print()
// }

// ---------------------------------------------
// Maps
// ---------------------------------------------
// func main() {
// structと似てる！

// colors := map[string]string{
// 	"red":   "#ff0000",
// 	"green": "#4327as",
// }

// var colors map[string]string

// colors := make(map[string]string)
// // structだとcolors.white = hogeみたいな感じになるという相違点
// // keyとvalueの肩が指定されている(structはされてない)
// colors["white"] = "#fff"

// delete(colors, "white")

// fmt.Println(colors)
// }

// ---------------------------------------------
// Interfaces
// ---------------------------------------------
// type bot interface {
// 	getGreeting() string
// }

// type englishBot struct{}
// type spanishBot struct{}

// func main() {
// 	eb := englishBot{}
// 	sb := spanishBot{}

// 	printGreeting(eb)
// 	printGreeting(sb)
// }

// func printGreeting(b bot) {
// 	fmt.Println(b.getGreeting())
// }

// func (englishBot) getGreeting() string {
// 	return "Hi there!"
// }

// func (spanishBot) getGreeting() string {
// 	return "Hola!"
// }

// ---------------------------------------------
// Interfaces with HTTP requests
// ---------------------------------------------

// type logWriter struct{}

// func main() {
// resp, err := http.Get("https://google.com")
// if err != nil {
// 	fmt.Println("Error:", err)
// 	os.Exit(1)
// }

// fmt.Println(resp)

// bs := make([]byte, 999999)
// // resp.Bodyはio.ReadCloserインターフェイスをtypeに持つ
// // Read(byte[]) (int, error)メソッド、Close() (error)メソッドが使える
// resp.Body.Read(bs)
// fmt.Println(string(bs))

// interface Writer { Write(p byte[]) (n int, err error)}
// func Copy(dst Writer, src Reader)
// io.Copy(os.Stdout, resp.Body)

// 	lw := logWriter{}
// 	io.Copy(lw, resp.Body)
// }

// func (logWriter) Write(bs []byte) (int, error) {
// 	// interfaceのメソッドは必ずしも正しくなくても良い？
// 	// return 1 nil (-> miss!)
// 	fmt.Println(string(bs))
// 	fmt.Println("Just wrote this many bytes", len(bs))
// 	return len(bs), nil
// }

// ---------------------------------------------
// Interfaces with HTTP requests Assignments
// ---------------------------------------------
// type triangle struct {
// 	height float64
// 	base   float64
// }
// type square struct {
// 	sideLength float64
// }

// type shape interface {
// 	getArea() float64
// }

// func main() {
// 	t := triangle{
// 		height: 20.2,
// 		base:   3.6,
// 	}
// 	s := square{
// 		sideLength: 8.4,
// 	}
// 	printArea(t)
// 	printArea(s)
// }

// func (t triangle) getArea() float64 {
// 	return t.height * t.base / 2
// }

// func (s square) getArea() float64 {
// 	return s.sideLength * s.sideLength
// }

// func printArea(s shape) {
// 	fmt.Println(s.getArea())
// }

// func main() {
// 	// go run main.go hoge.txtとターミナルで打つと、
// 	// ファイル内容がPrintlnされるプログラム
// 	// hint -> os.Args, os.Open
// 	fmt.Println(os.Args)
// }
