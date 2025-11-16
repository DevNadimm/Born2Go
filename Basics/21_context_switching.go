package main

import "fmt"

func main() {
	fmt.Println("🔁 Context Switching | PCB | Concurrency")
}

/*

# 🔁 CONTEXT SWITCHING | PCB | CONCURRENCY

## 🧩 1. WHAT IS CONTEXT SWITCHING?

➡️ **Context Switching** is the process of **saving and restoring** the state (context) of a CPU so that execution can be resumed from the same point later.

It allows **multiple processes** to share a single CPU efficiently.

🧠 Think of it like:
> "Pausing one task, saving its progress, and loading another task to continue."

---

## 🗂️ 2. PROCESS CONTROL BLOCK (PCB)

Each process has a **PCB (Process Control Block)** — a data structure maintained by the operating system.

It stores **process-specific information** such as:

| Field | Description |
|--------|--------------|
| Process ID (PID) | Unique identifier for each process |
| Process State | Running, Ready, Waiting |
| Program Counter (PC) | Address of next instruction |
| CPU Registers | Saved values of registers |
| Memory Info | Code, data, stack pointers |
| I/O Info | List of open files, I/O status |

When the CPU switches from one process to another:
1️⃣ OS saves the current process’s state to its PCB.  
2️⃣ OS loads the next process’s state from its PCB.  
3️⃣ Execution continues from the saved point.

---

## ⚙️ 3. STEPS OF CONTEXT SWITCHING

1️⃣ Save context of the current process (registers, PC, etc.)  
2️⃣ Update PCB of the current process.  
3️⃣ Select next process from the ready queue.  
4️⃣ Load its context from PCB.  
5️⃣ Update CPU registers and PC.  
6️⃣ Resume execution.

🔁 This cycle continues repeatedly — giving the illusion of **multitasking**.

---

## ⚡ 4. CONCURRENCY IN OPERATING SYSTEMS

**Concurrency** means multiple tasks **start, run, and complete in overlapping time periods** — but not necessarily simultaneously.

🧩 Example:
While one process waits for I/O, another can use the CPU.

So, concurrency = “**making progress on multiple tasks at once**,” even on a single-core CPU.

---

## 🧠 5. KEY IDEA

Context Switching is the **mechanism** that makes **concurrency** possible in a single CPU.

It gives each process a “virtual CPU” — managed through the PCB.

*/
