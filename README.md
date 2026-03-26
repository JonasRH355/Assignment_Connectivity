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
Create directed or undirected graphs
Add and remove vertices
Add and remove edges (or arcs)
<br>

### 🔍 Graph Traversal
DFS (Depth-First Search)
BFS (Breadth-First Search)
<br>

### 🔗 Connectivity Analysis
Check if the graph is connected
Identify maximal strongly connected subgraphs

<br>

---

### 🔁 Transitive Closure
Direct closure: vertices reachable from a given node
Inverse closure: vertices that can reach a given node

<br>

---

## 🛠️ Technologies Used
Language: Go (Golang)
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
2. Remove vertex
3. Add edge
4. Remove edge
5. DFS traversal
6. BFS traversal
7. Transitive closure
8. Check connectivity
9. Strongly connected components
0. Exit

<br>

---

## 🧠 Algorithms Implemented
<br>
Depth-First Search (DFS)
Explores as far as possible along each branch before backtracking.
Breadth-First Search (BFS)
Explores nodes level by level.
Strongly Connected Components (SCC)
Implemented using Kosaraju’s Algorithm:
DFS to determine finishing order
Graph reversal
DFS on reversed graph
Transitive Closure
Computed using DFS/BFS from a given vertex
Inverse closure uses graph reversal

<br>

---

## 🎯 Objectives
Apply graph theory concepts in practice
Understand traversal and connectivity algorithms
Build a user-friendly interactive system
Ensure code organization and maintainability

<br>

---

## 🎥 Demonstration

A video demonstration is included showing:

Program execution
Feature walkthrough
Explanation of implementation decisions

<br>

---

## 👥 Team
Jônatas Régis Hellmann
Leonardo de Borba Cardoso

<br>

---

## 📌 Notes
The program is fully standalone and requires no external libraries
Designed for ease of use via command-line interface
