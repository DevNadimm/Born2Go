package main

import "fmt"

func main() {
	/* 
	=== GO DATA TYPES AND MEMORY ALLOCATION ===
	
	1. INTEGER TYPES (Signed):
	   - int: Platform dependent (32-bit on 32-bit systems, 64-bit on 64-bit systems)
	   - int8: 8 bits (1 byte) | Range: -128 to 127
	   - int16: 16 bits (2 bytes) | Range: -32,768 to 32,767
	   - int32: 32 bits (4 bytes) | Range: -2,147,483,648 to 2,147,483,647
	   - int64: 64 bits (8 bytes) | Range: -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
	   
	   ** How int64 works on 32-bit computer:
	      - On a 32-bit system, int64 still allocates 8 bytes (64 bits)
	      - The CPU handles it using TWO 32-bit registers
	      - Operations are slower because they require multiple instructions
	      - Memory alignment may require padding
	
	2. UNSIGNED INTEGER TYPES:
	   - uint: Platform dependent (32 or 64 bits)
	   - uint8: 8 bits (1 byte) | Range: 0 to 255
	   - uint16: 16 bits (2 bytes) | Range: 0 to 65,535
	   - uint32: 32 bits (4 bytes) | Range: 0 to 4,294,967,295
	   - uint64: 64 bits (8 bytes) | Range: 0 to 18,446,744,073,709,551,615
	
	3. FLOATING POINT TYPES:
	   - float32: 32 bits (4 bytes) | IEEE-754 single precision
	   - float64: 64 bits (8 bytes) | IEEE-754 double precision
	   
	   ** float64 on 32-bit computer:
	      - Still allocates 8 bytes
	      - Uses FPU (Floating Point Unit) if available
	      - Otherwise handled through software emulation (slower)
	
	4. BOOLEAN TYPE:
	   - bool: 1 byte (8 bits) | Values: true or false
	   - Note: Only uses 1 bit logically, but allocates 1 byte for addressability
	
	5. BYTE TYPE:
	   - byte: Alias for uint8 | 1 byte (8 bits) | Range: 0 to 255
	   - Used for raw binary data
	
	6. RUNE TYPE:
	   - rune: Alias for int32 | 4 bytes (32 bits)
	   - Represents a Unicode code point
	   - Range: -2,147,483,648 to 2,147,483,647 (but typically 0 to 0x10FFFF for valid Unicode)
	
	7. STRING TYPE:
	   - string: Variable size
	   - Internal structure: 16 bytes on 64-bit (or 8 bytes on 32-bit)
	     * Pointer to byte array: 8 bytes (or 4 bytes on 32-bit)
	     * Length integer: 8 bytes (or 4 bytes on 32-bit)
	   - Actual string data stored separately in memory
	   - Immutable in Go
	
	8. Printf() and %T FORMAT SPECIFIER:
	   - fmt.Printf(): Formatted print function
	   - %T: Prints the TYPE of a variable
	   - Other common format specifiers:
	     * %v: Default format
	     * %d: Decimal integer
	     * %f: Floating point
	     * %s: String
	     * %t: Boolean
	     * %c: Character (rune)
	     * %p: Pointer address
	     * %b: Binary
	     * %x: Hexadecimal
	*/
	
	// Examples:
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	
	var ui8 uint8 = 255
	var ui16 uint16 = 65535
	var ui32 uint32 = 4294967295
	var ui64 uint64 = 18446744073709551615
	
	var f32 float32 = 3.14159
	var f64 float64 = 3.141592653589793
	
	var b bool = true
	var bt byte = 65 // ASCII 'A'
	var r rune = '♥' // Unicode heart
	var s string = "Hello, Go!"
	
	fmt.Printf("int8: %d, Type: %T\n", i8, i8)
	fmt.Printf("int16: %d, Type: %T\n", i16, i16)
	fmt.Printf("int32: %d, Type: %T\n", i32, i32)
	fmt.Printf("int64: %d, Type: %T\n", i64, i64)
	
	fmt.Printf("\nuint8: %d, Type: %T\n", ui8, ui8)
	fmt.Printf("uint16: %d, Type: %T\n", ui16, ui16)
	fmt.Printf("uint32: %d, Type: %T\n", ui32, ui32)
	fmt.Printf("uint64: %d, Type: %T\n", ui64, ui64)
	
	fmt.Printf("\nfloat32: %f, Type: %T\n", f32, f32)
	fmt.Printf("float64: %f, Type: %T\n", f64, f64)
	
	fmt.Printf("\nbool: %t, Type: %T\n", b, b)
	fmt.Printf("byte: %d (char: %c), Type: %T\n", bt, bt, bt)
	fmt.Printf("rune: %d (char: %c), Type: %T\n", r, r, r)
	fmt.Printf("string: %s, Type: %T\n", s, s)
}