# swim-api
A Rest API written in Go

### Usage
- Clone repository: ``git clone https://github.com/l4rma/swim-api.git``
- Compile code: ``make build``
- Run API: ``make run``
```

#### Example requests
```bash
# Add swimmer
curl localhost:8080/swimmers/add -H "application/json" -d '{"name":"Lars","age":35}'
# Get all swimmers
curl localhost:8080/swimmers | jq
```

### Todo:
- [x] Create MVP with inmemory DB
- [ ] Add controller for update function
- [ ] Add a repository layer for a SQL db
- [ ] Make serverless with lambda and aurora db

### Tables
#### Swimmers

| Field Name | Type        | Description                            |
|------------|-------------|----------------------------------------|
| `ID`       | `string`    | Unique identifier for the swimmer.     |
| `Name`     | `string`    | Full name of the swimmer.              |
| `Age`      | `int`       | Age of the swimmer.                    |
| `CreatedAt`| `time.Time` | Timestamp when the swimmer was created.|
| `IsActive` | `bool`      | Indicates if the swimmer is active.    |

#### Sessions

| Field Name | Type        | Description                                   |
|------------|-------------|-----------------------------------------------|
| `ID`       | `string`    | Unique identifier for the session.           |
| `SwimmerID`| `string`    | Foreign key linking to the swimmer's `ID`.    |
| `Date`     | `time.Time` | Date of the session in `time.Time`.           |
| `Distance` | `int`       | Total distance swam in meters.               |
| `Duration` | `int`       | Total duration of the session in minutes.     |
| `Intensity`| `string`    | Intensity level (e.g., "low", "moderate").    |
| `Style`    | `string`    | Swimming style (e.g., "freestyle").           |
| `Notes`    | `string`    | Additional notes about the session.           |

