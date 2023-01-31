# Email-finder

This project is written in Go. Its main purpose is to expose API's to find a specific text from a Full Text Search dataset containing email information.

## Getting Started

It will expose an API to search for body content within a Full Text Search engine that contains email data.

## Prerequisites

ZincSearch should already be filled with data in `email_index` index.

## Installation

Follow the next steps:

- Go to the `src` folder.
- Execute the command:
    ```
    go build .
    ```
- Run the executable file:
    ```
    main.go.exe
    ```

It will expose a server running on http://localhost:3000.

## Usage

In any browser you can now make a request to:
```
GET - http://localhost:3000/emailfinder/search/[text_to_search]
```
For example:
```
http://localhost:3000/emailfinder/search/apte
```

It will return a collection of results in JSON format.

```
[
  {
    "id": String,
    "subject": String,
    "from": String,
    "to": String,
    "body": String,
    "date": String
  },
  ...
]
```
