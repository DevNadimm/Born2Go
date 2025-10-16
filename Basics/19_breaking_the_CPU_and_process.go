package main

import "fmt"

func main() {
	fmt.Println("Hello, CPU World!")
}

/*
===============================================================
🧠 CPU AND PROCESS EXECUTION — SYSTEM OVERVIEW
===============================================================

📌 1. CPU STRUCTURE
---------------------------------------------------------------
A CPU (Central Processing Unit) is composed of three major components:

1️⃣ ALU — Arithmetic Logic Unit  
   • Performs arithmetic (add, subtract, multiply, divide)
     and logical (AND, OR, NOT, compare) operations.

2️⃣ CU — Control Unit  
   • Directs and coordinates all CPU operations.
   • Controls data flow between CPU, memory, and I/O devices.

3️⃣ Register Set — Temporary high-speed storage inside the CPU.  
   • Used to hold instructions, addresses, and intermediate results.

---------------------------------------------------------------
📍 2. IMPORTANT REGISTERS
---------------------------------------------------------------
- **PC (Program Counter):**
  Holds the address of the next instruction to execute.

- **IR (Instruction Register):**
  Stores the current instruction fetched from memory.

- **SP (Stack Pointer):**
  Points to the top of the stack in memory (used for function calls, returns, etc.).

- **BP (Base Pointer):**
  Points to the base of the current stack frame (helps in accessing local variables).

---------------------------------------------------------------
⚙️ 3. INSTRUCTION EXECUTION CYCLE
---------------------------------------------------------------
The CPU follows the **Fetch–Decode–Execute Cycle**, which repeats continuously:

1️⃣ **Fetch**  
    - The Program Counter (PC) points to the next instruction in memory (RAM).  
    - The instruction is fetched and loaded into the Instruction Register (IR).

2️⃣ **Increment PC**  
    - PC = PC + 1 → prepares for the next instruction.

3️⃣ **Decode**  
    - The Control Unit decodes the instruction in the IR.  
    - Breaks it into parts (e.g., operands and operator like `2 3 +`).

4️⃣ **Execute**  
    - The CU sends the decoded operation to the ALU.  
    - The ALU performs the computation or logic.  
    - The result is stored in a register or memory location.

5️⃣ **Repeat**  
    - The cycle continues for every instruction in the program.

---------------------------------------------------------------
💾 4. PROCESS AND MEMORY ORGANIZATION
---------------------------------------------------------------
When a program is executed, it becomes a **process**, which is a
combination of **CPU execution + memory management**.

Memory layout of a process typically includes:
   - **Code Segment:** Contains executable instructions.
   - **Data Segment:** Stores global and static variables.
   - **Heap:** Used for dynamic memory allocation.
   - **Stack:** Stores function calls, local variables, and return addresses.

🧩 The **Stack Pointer (SP)** keeps track of the current top of the stack.
🧩 The **Base Pointer (BP)** helps locate variables in the current stack frame.

Together, these define a **virtual computer** (or logical CPU)
for each process running on the system.

===============================================================
*/
