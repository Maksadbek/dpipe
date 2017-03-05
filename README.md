# dpipe
Tool to pipeline data I/O

## Requirements

Read data
- Read data from the given CSV file `hotels.csv`. The first line is a header
   which describes all field names.

Validate data
- A hotel name may only contain UTF-8 characters.
- The hotel URL must be valid (please come up with a good definition of "valid").
- Hotel ratings are given as a number from 0 to 5 stars. There may be no negative numbers.

Write the valid data in *two* of the following formats of your choice:  
- XML, JSON, YAML, HTML, SQLite Database, or your own custom format.  
- The output must be in the same directory as the input.

Bonus tasks
* Make the tool extensible to new output formats
* We care more about code quality (readability, software architecture)
  than about performance - although fast execution is a plus.
* Unit tests would be nice
* Add options to sort/group the data before writing it

## Build

## How to use
