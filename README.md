# tdl-go

A command-line to-do list application written in Go that helps you manage tasks with due dates, stored in a MySQL database.

## Table of Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [Features](#features)
- [Usage](#usage)
- [Configuration](#configuration)
- [Docker](#docker)
- [Contributing](#contributing)
- [Questions](#questions)

## Installation

**Prerequisites:**
- Go 1.25.0 or later
- MySQL database server

**Build from source:**

```bash
git clone https://github.com/krosengr4/tdl-go.git
cd tdl-go
go build -o tdl-go
```

## Quick Start

1. Create your database and tasks table in MySQL:

```sql
CREATE DATABASE todo_db;
USE todo_db;

CREATE TABLE tasks (
    task_id INT AUTO_INCREMENT PRIMARY KEY,
    description VARCHAR(255) NOT NULL,
    completed BOOLEAN DEFAULT 0,
    due_date DATE NOT NULL
);
```

2. Copy the example environment file and configure your database credentials:

```bash
cp .env.example .env
```

3. Edit `.env` with your MySQL credentials:

```
DB_USERNAME=your_username
DB_PASSWORD=your_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=todo_db
```

4. Run the application:

```bash
./tdl-go
```

## Features

- **Add Tasks** - Create new tasks with descriptions and due dates
- **Check Off Tasks** - Mark tasks as completed
- **View All Tasks** - Display all tasks sorted by due date
- **Filter by Status** - View only pending or completed tasks
- **Delete Tasks** - Remove tasks from your list
- **Persistent Storage** - Tasks stored in MySQL database

## Usage

Launch the application and interact with the menu:

```
-----WELCOME TO YOUR TO DO LIST-----
__________________________________________________

---OPTIONS---
____________________
1 - Add A New Task
2 - Check Off A Task
3 - View All Tasks
4 - View All Pending Tasks
5 - View All Completed Tasks
6 - Delete A Task
0 - Exit

Enter option:
```

### Adding a Task

Select option `1`, then enter your task description and due date:

```
Enter task description: Finish project report
Enter due date (mm-dd-yyyy): 02-15-2026
```

### Viewing Tasks

Tasks are displayed with their ID, description, due date, and status:

```
ID: 1 | Description: Finish project report | Due: 02-15-2026 | Status: ❌ Pending
ID: 2 | Description: Call dentist | Due: 02-10-2026 | Status: ✅ Completed
```

### Checking Off or Deleting Tasks

Select option `2` or `6`, then enter the task ID. Enter `0` to go back without making changes.

## Configuration

The application uses environment variables for database configuration. Create a `.env` file in the project root:

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_USERNAME` | MySQL username | `root` |
| `DB_PASSWORD` | MySQL password | - |
| `DB_HOST` | Database host | `localhost` |
| `DB_PORT` | Database port | `3306` |
| `DB_NAME` | Database name | `todo_db` |

The application will fall back to system environment variables if no `.env` file is found.

## Docker

Build and run with Docker:

```bash
docker build -t tdl-go .
docker run -it --env-file .env tdl-go
```

**Note:** Ensure your MySQL database is accessible from within the Docker container (use network host or configure appropriate networking).

## Contributing

**Please contribute to this project:**

- [Submit Bugs and Request Features you'd like to see Implemented](https://github.com/krosengr4/tdl-go/issues)

## Questions

- [Link to my GitHub Profile](https://github.com/krosengr4)

- For any additional questions, email me at rosenkev4@gmail.com
