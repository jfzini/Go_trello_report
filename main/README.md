# Go Report Generator

This Go program generates a report based on actions and cards data from a JSON file. It filters actions based on specific criteria, creates a report dictionary from the filtered actions and cards, and writes the report to a text file.

## Features

- **Action Filtering**: Filters actions based on type, date, and specific data fields.
- **Report Generation**: Creates a detailed report string from the filtered actions and associated cards.
- **File Operations**: Reads from a JSON file and writes the generated report to a text file.

## Prerequisites

- Go  1.16 or later
- A JSON file exported from Trello with the expected structure.

## Usage

1. Ensure you have a JSON file with the Trello Board data in the same directory as the Go program.
2. Change the constants `JSONPath` and `DoneColumnID` in `main.go` to the path of the JSON file and the ID of the "Done" column, respectively.
3. Run the Go program:
  
  ```bash
  go run main.go
  ```
  
4. It will automatically read from the JSON file, filter the actions, generate a report, and write the report to `report.txt`.

## Example JSON Structure

The JSON file should be structured as follows:

```json
{
  "actions": [
    {
      "type": "updateCard",
      "data": {
        "checkItem": {
          "state": "complete",
          "name": "Task Name"
        },
        "listAfter": {
          "id": "65d3be0d2d74d5884cf6c4e7"
        },
        "card": {
          "id": "card_id"
        }
      },
      "date": "2024-02-22T15:04:05.000Z"
    }
  ],
  "cards": [
    {
      "id": "card_id",
      "labels": [
        {
          "color": "black_light",
          "name": "Label Name"
        }
      ],
      "name": "Card Name"
    }
  ]
}
```

## Report Format

The generated report will be formatted as follows:

```
Label Name:
    - Card Name
        - Task Name
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.
