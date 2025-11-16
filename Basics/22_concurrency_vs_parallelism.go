package main

import "fmt"

func main() {
	fmt.Println("🧠 Concurrency vs Parallelism")
}

/*

# 🧠 CONCURRENCY VS PARALLELISM

---

## ⚡ 1. CONCURRENCY

Concurrency = **Doing many things at once (in progress)**  
➡️ Tasks **overlap in time**, but not necessarily execute simultaneously.

🧩 Example:
A single-core CPU switches between multiple processes very fast using **context switching**.

So, it seems like multitasking — but only one task runs at any given moment.

🧠 Analogy:
> A single person cooking multiple dishes — preparing one, then another, switching between them quickly.

---

## ⚙️ 2. PARALLELISM

Parallelism = **Actually doing many things at the same time**  
➡️ Requires **multiple CPU cores** or processors.

Each core executes a different task **simultaneously**.

🧠 Analogy:
> Several chefs cooking different dishes at the same time.

---

## 🔍 3. DIFFERENCE TABLE

| Feature | Concurrency | Parallelism |
|----------|--------------|-------------|
| Definition | Multiple tasks making progress together | Multiple tasks executing at the same time |
| Hardware | Can occur on single-core CPU | Requires multi-core CPU |
| Goal | Better resource utilization | Faster execution |
| Achieved By | Context Switching | Multiple cores/threads |
| Example | OS switching between tasks | CPUs processing tasks simultaneously |

---

## 💡 4. KEY INSIGHT

➡️ All parallel systems are concurrent,  
but **not all concurrent systems are parallel**.

Concurrency is about **structure** — managing many tasks.  
Parallelism is about **execution** — doing them simultaneously.

---

## 🧠 5. SUMMARY

| Concept | Description |
|----------|--------------|
| Context Switching | Enables concurrency |
| Concurrency | Multiple tasks in progress |
| Parallelism | Multiple tasks at the same time |

Together, they form the foundation of **modern multitasking systems**.

*/
