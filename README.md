# Concurrency Development Labs

This repository contains a collection of Go programs designed to explore and demonstrate various concurrency patterns, synchronization primitives, and parallel processing techniques.

## Project Structure

The repository is organized into several labs, each focusing on a specific concurrency concept:

*   **`lab3/` & `lab4/`**: **Barrier Synchronization**
    *   Implementations of barriers to synchronize multiple goroutines at specific points.
*   **`lab5/`**: **Dining Philosophers Problem**
    *   Solutions to the classic synchronization problem, exploring deadlock and resource sharing.
*   **`lab6/`**: **Producer-Consumer Problem**
    *   A demonstration of the producer-consumer pattern using Go channels and WaitGroups to coordinate work between multiple producing and consuming goroutines.
*   **`lab/project lab/`**: **Conway's Game of Life**
    *   A parallelized implementation of Conway's Game of Life using the [Ebiten](https://ebiten.org/) game library.
    *   Demonstrates data parallelism by splitting the grid update across multiple worker goroutines.

## Prerequisites

*   [Go](https://golang.org/dl/) (version 1.17 or later recommended)

## How to Run

### Standard Labs (Lab 3 - Lab 6)

To run any of the standalone lab files, navigate to the file's directory (or use the full path) and use `go run`.

**Example: Running the Producer-Consumer Lab**

```bash
cd lab/lab6
go run procon.go
```

**Example: Running the Dining Philosophers Lab**

```bash
cd lab/lab5
go run "dinPhil(1).go"
# Note: Quotes might be needed for filenames with parentheses depending on your shell.
```

### Game of Life Project

The Game of Life project is a module that depends on the Ebiten library.

1.  Navigate to the project directory:
    ```bash
    cd "lab/project lab"
    ```

2.  Ensure dependencies are installed:
    ```bash
    go mod tidy
    ```

3.  Run the application:
    ```bash
    go run .
    ```

## License

This project is licensed under the GNU General Public License v3 (GPLv3). See [LICENSE.md](LICENSE.md) for details.