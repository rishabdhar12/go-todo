---

# Todo CLI App

This is a simple Todo CLI (Command Line Interface) application built using Golang. It allows you to manage your tasks right from your terminal!

## Features

- **Add Tasks:** Easily add tasks to your todo list.
- **List Tasks:** View all your tasks in one place.
- **Complete Tasks:** Mark tasks as completed.
- **Delete Tasks:** Remove tasks from the list.

## Installation

1. **Clone the Repository:**

    ```bash
    git clone https://github.com/rishabdhar12/go-todo.git
    ```

2. **Build the Executable:**

    ```bash
    cd go-todo 
    go build -o todo main.go
    ```

## Usage

- **Add a Task:**

    ```bash
    ./todo -a
    ```

- **List all Tasks:**

    ```bash
    ./todo -l
    ```

- **Complete a Task:**

    ```bash
    ./todo -u 
    ```

- **Delete a Task:**

    ```bash
    ./todo -d
    ```

- **Help:**

    ```bash
    ./todo -h
    ```
 
- **Version:**

    ```bash
    ./todo -v
    ```

## Current Version

- 1.0.0


## Contributing

Contributions are welcome! If you have any improvements or new features in mind, feel free to submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

---
