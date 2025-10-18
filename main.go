package main

import (
	"fmt"
	"math"
	"net/http"
)

var (
	task1 = "Learn new tech daily"
	task2 = "Increase my income"
	task3 = "Plan a Namibia tour"

	allTasks = []string{task1, task2, task3}
)

func main() {
	http.HandleFunc("/", welcomeTodos)
	http.HandleFunc("/show-tasks", showTasks)
	http.HandleFunc("/add-task", addTask)

	http.ListenAndServe(":9000", nil)

	/* Start of Format Specifier Topic

	str := "Hello, world!"
	num1 := 42
	num2 := 3.14159
	num3 := 1234567890
	boolean := true
	char := 'A'

	// %T: print type of the value
	fmt.Printf("%T\n", str)
	// %v: print default format for each type
	fmt.Printf("%v\n", str)

	// String format specifiers
	fmt.Printf("%s\n", str)         // %s: print string
	fmt.Printf("%q\n", str)         // %q: print quoted string
	fmt.Printf("%x\n", []byte(str)) // %x: print hex encoding of bytes
	fmt.Printf("%X\n", []byte(str)) // %X: print uppercase hex encoding of bytes

	// Integer format specifiers
	fmt.Printf("%d\n", num1) // %d: print decimal integer
	fmt.Printf("%b\n", num1) // %b: print binary integer
	fmt.Printf("%o\n", num1) // %o: print octal integer
	fmt.Printf("%x\n", num1) // %x: print hex encoding of integer
	fmt.Printf("%X\n", num1) // %X: print uppercase hex encoding of integer

	fmt.Printf("%c\n", char) // %c: print character

	// Floating-point format specifiers
	fmt.Printf("%f\n", num2) // %f: print floating-point number
	fmt.Printf("%e\n", num2) // %e: print scientific notation of floating-point       number
	fmt.Printf("%E\n", num2) // %E: print scientific notation of floating-point number with uppercase E
	fmt.Printf("%g\n", num2) // %g: print floating-point number in decimal or scientific notation, depending on the value
	fmt.Printf("%G\n", num2) // %G: print floating-point number in decimal or scientific notation, depending on the value with uppercase E

	// Width and precision
	fmt.Printf("|%5d|\n", num1)    // %5d: print decimal integer with minimum width of 5 characters
	fmt.Printf("|%-5d|\n", num1)   // %-5d: print decimal integer with minimum width of 5 characters, left-justified
	fmt.Printf("|%5.2f|\n", num2)  // %5.2f: print floating-point number with minimum width of 5 characters and 2 digits after the decimal point
	fmt.Printf("|%-5.2f|\n", num2) // %-5.2f: print floating-point number with minimum width of 5 characters and 2 digits after the decimal point, left-justified
	// Boolean format specifiers
	fmt.Printf("%t\n", boolean) // %t: print boolean value
	// Pointer format specifier
	fmt.Printf("%p\n", &num3) // %p: print pointer address

	End of Format Specifier Topic*/

	// // example interface
	// tempMain()

}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

type geometry interface {
	area() float64
	perim() float64
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func tempMain() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)
}

func welcomeTodos(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Welcome todos app!"
	fmt.Fprintln(writer, greeting)
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Todos List: ")
	fmt.Fprintln(writer)
	for _, task := range allTasks {
		fmt.Fprintln(writer, task)
	}
}

func addTask(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Use POST Method", http.StatusMethodNotAllowed)
		return
	}

	task := request.FormValue("task")
	if task == "" {
		http.Error(writer, "Task is required", http.StatusBadRequest)
		return
	}

	allTasks = append(allTasks, task)

	showTasks(writer, request)
}
