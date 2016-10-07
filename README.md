# Code Evaluation

This goal of this project is to generate a report in CSV (Comma Seperated Value) format.

The end report will look like

```
Timestamp,Location,Price
2016-01-02T10:00:00-0700,Unit A,23.00
2016-01-02T10:00:00-0700,Unit B,21.00
2016-01-02T11:00:00-0700,Unit A,24.10
2016-01-02T11:00:00-0700,Unit B,13.04
```

The report should be writen out to `./reports/report.csv` and ordered by
Timestamp, Location. The final application should be able to generate the report with a single command.

### Process

The prices are available at `http://localhost:4000/prices?location_id=LOCATION_ID`

Before collecting the prices you will need the list of locations. These are obtained at
`http://localhost:4000/locations`

Both services return their data as JSON. The service is a bit unreliable, so be
prepared for an intermittent error.

### Tools

There is a Gemfile which you can modify as needed with any gem you need. Any
version of Ruby is fine.

### Getting started

* create your own git branch. submit all commits to this branch, not master.
* start the server `./server`
* start coding
* commit as needed
