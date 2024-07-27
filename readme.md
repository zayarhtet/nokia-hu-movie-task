# Nokia Movie Console Application

## Overview

This Go application is a console-based program for managing and listing movies. It supports various interactive commands and flags for filtering and sorting movies based on different criteria.

## Interactive Commands

The application supports the following interactive commands:

- **l**: List movies with optional filters and sorting options.
- **a -p**: Add a new person (actor or director) to the database.
- **a -m**: Add a new movie to the database with its details.

Additional commands and features:

- **help**: Displays detailed explanations of commands and their usage.
- **d -p "<person-name>"**: Deletes a person with double quoted from the database and removes them from associated movies if applicable.

## Flags

The following flags can be used with the **l** command for listing movies:

- **-v**: List movies with detailed information including actors.
- **-t "regex"**: Filter movies by title using a regular expression.
- **-d "regex"**: Filter movies by director's name using a regular expression.
- **-a "regex"**: Filter movies by actor's name using a regular expression.
- **-la**: Sort movies in ascending order by length.
- **-ld**: Sort movies in descending order by length.

## Execution

1. **Clone the Repository:**

```bash
git clone <repository-url>
cd <repository-directory>
```

2. **Install Dependencies:**

Ensure you have Go installed on your system. If not, follow the official installation guide.

3. **Run the Application:**

Enter the root directory of the project. (The directory where this readme is located)

```bash
go run ,
```

Follow the prompts to interact with the application using the supported commands and flags.

## Testing

To run unit tests and view code coverage:

1. **Run Tests:**

```bash
go test -v ./test/...
```
This command runs all tests in the current directory and subdirectories, displaying detailed test output with test case names and results.