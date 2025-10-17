package main

import "fmt"

func main() {
fmt.Println("⚙️ CPU and Process Execution")
}

/*

# ⚙️ CPU AND PROCESS EXECUTION — SYSTEM OVERVIEW

## 🧩 1. CPU STRUCTURE

A CPU consists of three major components:

1️⃣ ALU — Arithmetic Logic Unit
• Performs arithmetic (add, subtract, multiply, divide)
and logical (AND, OR, NOT, compare) operations.

2️⃣ CU — Control Unit
• Directs and coordinates CPU operations.
• Controls data flow between CPU, memory, and I/O devices.

3️⃣ Register Set — High-speed temporary storage.
• Holds instructions, addresses, and intermediate results.

---

## 📍 2. IMPORTANT REGISTERS

• **PC (Program Counter):** Holds address of the next instruction.
• **IR (Instruction Register):** Stores the current instruction.
• **SP (Stack Pointer):** Points to the top of the stack.
• **BP (Base Pointer):** Points to the base of the current stack frame.

---

## ⚙️ 3. INSTRUCTION EXECUTION CYCLE

CPU follows a continuous **Fetch–Decode–Execute Cycle**:

1️⃣ Fetch → Instruction fetched from memory to IR.
2️⃣ Increment PC → PC = PC + 1 (next instruction).
3️⃣ Decode → Control Unit interprets the instruction.
4️⃣ Execute → ALU performs operation or logic.
5️⃣ Repeat → Cycle continues for each instruction.

---

## 💾 4. PROCESS AND MEMORY ORGANIZATION

When a program executes, it becomes a **process** —
a combination of **CPU execution** and **memory management**.

Memory Layout of a Process:
• Code Segment — Executable instructions
• Data Segment — Global/static variables
• Heap — Dynamic memory allocation
• Stack — Function calls, local variables, return addresses

🧠 The Stack Pointer (SP) tracks top of the stack.
🧠 The Base Pointer (BP) helps access local variables in the stack frame.

# Together, they define a “virtual CPU” for every running process.

*/
