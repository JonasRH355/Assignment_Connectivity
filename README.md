# Assignment_Connectivity
## 📖 Description

This project is an implementation of a graph manipulation and analysis tool, developed as part of a Computer Networks / Graph Theory assignment.

The program allows users to:

- Create and manipulate directed or undirected graphs
- Dynamically add/remove vertices and edges
- Traverse the graph using DFS (Depth-First Search) and BFS (Breadth-First Search)
- Compute transitive closure (direct and inverse)
- Verify if the graph is connected
- Identify Strongly Connected Components (SCCs)
<br>

---

## 🚀 Features
<br>

### 🧱 Graph Construction
Create directed or undirected graphs <br>
Add and remove vertices <br>
Add and remove edges (or arcs)
<br>

### 🔍 Graph Traversal
DFS (Depth-First Search) <br>
BFS (Breadth-First Search) <br>
<br>

### 🔗 Connectivity Analysis
Check if the graph is connected <br>
Identify maximal strongly connected subgraphs <br>

<br>

---

### 🔁 Transitive Closure
Direct closure: vertices reachable from a given node <br>
Inverse closure: vertices that can reach a given node <br>

<br>

---

## 🛠️ Technologies Used
Language: Go (Golang) <br>
No external dependencies (standard library only)

<br>

---

## 📂 Project Structure
. <br>
├── main.go          # Program entry point <br>
├── graph.go         # Graph structure and operations <br>
├── algorithms.go    # DFS, BFS, SCC, etc. <br>
└── README.md <br>

<br>

---

## ▶️ How to Run

1. Clone the repository
```
git clone git@github.com:JonasRH355/Assignment_Connectivity.git
cd Assignment_Connectivity
```

2. Build the executable
```
go build -o graph_app
```

3. Run the program
```
./graph_app
```
<br> OU <br>

1. Clone the repository
```
git clone git@github.com:JonasRH355/Assignment_Connectivity.git
cd Assignment_Connectivity
```

2. Build and execute 
```
go run .
```

<br>

---

## 🧪 Usage
The program provides an interactive menu:
<br>
1. Add vertex
2. Add edge
3. DFS traversal
4. BFS traversal
5. Transitive closure
6. Check connectivity
7. Print graph
8. Remove vertex
9. Remove edge
0. Exit

<br>

---

## 🧠 Algorithms Implemented
<br>
Depth-First Search (DFS) <br>
Explores as far as possible along each branch before backtracking. <br>
Breadth-First Search (BFS) <br>
Explores nodes level by level. <br>
Strongly Connected Components (SCC) <br>
Implemented using Kosaraju’s Algorithm: <br>
DFS to determine finishing order <br>
Graph reversal <br>
DFS on reversed graph <br>
Transitive Closure <br>
Computed using DFS/BFS from a given vertex <br>
Inverse closure uses graph reversal <br>

<br>

---

## 🎯 Objectives
Apply graph theory concepts in practice <br>
Understand traversal and connectivity algorithms <br>
Build a user-friendly interactive system <br>
Ensure code organization and maintainability <br>

<br>

---

## 👥 Team
Jônatas Régis Hellmann <br>
Leonardo de Borba Cardoso

<br>

---

## 📌 Notes
The program is fully standalone and requires no external libraries <br>
Designed for ease of use via command-line interface
