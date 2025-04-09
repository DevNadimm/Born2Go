// 🟦 Package declaration
package main

import "fmt"

// 🔸 Constant (Compile time constant)
// এটি code segment বা constant pool এ সংরক্ষিত থাকে
const pie = 3.1416

// 🔸 Global variable (initialized)
// এটি Data Segment-এ সংরক্ষিত হয়
var r = 10

// 🔹 Function: Prints which function is called
// এই ফাংশনের কোড থাকে Code Segment-এ
// প্যারামিটার name → Stack Segment-এ যায়
func printFuncName(name string) {
	fmt.Println(name, "is executed")
}

// 🔹 Function: Calculates the area of a circle
// radius এবং answer → Stack Segment-এ রাখা হয়
func areaOfACircle(radius int) {
	printFuncName("areaOfACircle") // অন্য ফাংশন কল
	answer := pie * float64(radius) * float64(radius)
	fmt.Println("Area of the circle:", answer)
}

// 🔹 init() function: Go তে এটি main() এর আগে রান হয়
func init() {
	fmt.Println("Hello, Go!")
}

// 🔹 main() function: Execution এর Entry Point
func main() {
	areaOfACircle(r) // r → Data Segment থেকে নেয়া হয়
}

/*
------------------------------------------------------
🧠 Internal Memory Explanation (In Bangla)
------------------------------------------------------

📌 Compile Phase:
-----------------
✅ কোড কম্পাইল হলে:
- সকল ফাংশনের কোড → Code Segment এ যায়
- গ্লোবাল ভ্যারিয়েবল r → Data Segment এ
- pie → কম্পাইল টাইম কনস্ট্যান্ট, Code Segment/Constant Pool এ থাকে
- Executable ফাইল তৈরি হয়

📌 Execution Phase:
-------------------
1. init() → প্রথমে রান হয় → "Hello, Go!" প্রিন্ট
2. main() → এরপর রান হয়
3. main() → areaOfACircle(10) কল করে
4. areaOfACircle → printFuncName("areaOfACircle") কল করে
5. শেষে area হিসেব করে প্রিন্ট হয়

📌 Memory Segments:
-------------------

🔹 Code Segment:
- Function bodies → main, init, areaOfACircle, printFuncName

🔹 Data Segment:
- Global variable: var r = 10

🔹 Stack Segment:
- Function কল হলে প্যারামিটার ও লোকাল ভ্যারিয়েবল এখানে যায়
  • main → (কিছু নেই)
  • areaOfACircle → radius, answer
  • printFuncName → name

🔹 Heap Segment:
- এই কোডে ব্যবহৃত হয়নি
- হিপে ডেটা রাখার জন্য make, new, append ইত্যাদি দরকার

📌 Output:
-----------
Hello, Go!  
areaOfACircle is executed  
Area of the circle: 314.16

------------------------------------------------------
*/

