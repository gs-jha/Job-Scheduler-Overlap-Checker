# Job-Scheduler-Overlap-Checker

An API service to check overlap between two ranges

## Prerequisites

- Go 1.21.5 or later
- Git
- Docker (for containerized deployment)
- Make (for running build pipeline tasks)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/gs-jha/Job-Scheduler-Overlap-Checker
   cd Job-Scheduler-Overlap-Checker
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Build the application:
   ```bash
   make build
   ```

## Running the Service

Start the server:
```bash
make run
```

The service will be available at `http://localhost:8080`.

## API Endpoint

**POST /api/v1/check-overlap**

**Request Body** (JSON):
```json
{
  "range1": {
    "start": "2025-07-16T12:00:00Z",
    "end": "2025-07-16T13:00:00Z"
  },
  "range2": {
    "start": "2025-07-16T12:30:00Z",
    "end": "2025-07-16T13:30:00Z"
  }
}
```
