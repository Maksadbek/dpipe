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

## Build & Install
All dependency packages can be restored with `godep restore`

Make by default builds and installs binary into $GOPATH/bin directory
```bash
$ make
```

Run tests and vet 
```bash
$ make vet
$ make test
```

## Configure

## How to use

## Add new filters, inputs, outputs
All filters, inputs and outputs are made as a plugins.

### Inputs
Inputs must be located in `inputs/<name of input>` directory
In order to add new input, the follwoing steps must be made:
- Create a directory and implement your input decoder, 
implement `dpipe.Input` interface, it is located in `inputs.go` file.
- Implement `LoadConf` method that loads configuration settings.
- Create `init` function to add the instance of your input with its name into global map of inputs: 
```
func init() {
	inputs.Add("csv", &CSV{})
}
```
- Import your input in `inputs/all/all.go` - this runs the `init` function in your input plugin
- Configure your input in `config.toml`

### Outputs
Outputs must be located in `outputs/<name of output>` directory
In order to add new output, the follwoing steps must be made:
- Create a directory and implement your output encoder 
implement `dpipe.Output` interface, it is located in `outputs.go` file.
- Implement `LoadConf` method that loads configuration settings.
- Create `init` function to add the instance of your output with its name into global map of outputs: 
```
func init() {
	inputs.Add("json", &JSON{})
}
```
- Import your input in `inputs/all/all.go` - this runs the `init` function in your input plugin
- Configure your input in `config.toml`

### Filters
Filters are added like Outputs and Inputs
- Implement `dpipe.Filter` interface
- Set up which field to filter in config toml:
```
[filters]
        [filters.encodingUTF8]
                enabled = true # if enabled is True, this filter is active
                field = "name" # field that filter must be applied

```
