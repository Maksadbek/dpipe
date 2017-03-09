# dpipe
Tool to pipeline data I/O

## Requirements

Read data
- [x] Read data from the given CSV file `hotels.csv`. The first line is a header
   which describes all field names.

Validate data - validations are made as plugins, they are pluggable.
- [x] A hotel name may only contain UTF-8 characters.
- [x] The hotel URL must be valid (please come up with a good definition of "valid").
- [x] Hotel ratings are given as a number from 0 to 5 stars. There may be no negative numbers.


Write the valid data in *two* of the following formats of your choice:  
- [x] XML, JSON, YAML, HTML, SQLite Database, or your own custom format. XML and JSON is done 
- [x] The output must be in the same directory as the input. By default output and input are same directories, it can be configured.

Bonus tasks
- [x] Make the tool extensible to new output formats
- [ ] We care more about code quality (readability, software architecture)
  than about performance - although fast execution is a plus - no performance optimizations are made
- [x] Unit tests would be nice - only plugins are covered with tests
- [x] Add options to sort/group the data before writing it. Aggregators are made as a plugins, it can be pluggable, by default
sorting plugin is implemented.

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

Set up inputs and outputs
Each input and outputs require file that it will read or write into.
Sample config is here:
```toml
# inputs, available inputs: csv
[inputs]
        [inputs.csv] # csv input
                # default file name is hotels.csv
                file = "data/hotels.csv"
# outputs, available outputs: json, xml
[outputs]
        [outputs.json] # json output
                # default file name is hotels.json
                file = "data/hotels.json"
        [outputs.xml] # xml output
                # default file name is hotels.xml
                file = "data/hotels.xml"
# filters aka validators, available filters: encodintUTF8, range, url
# set which filter must check which field
[filters]
        [filters.encodingUTF8]
                enabled = true
                field = "name" # field name to validate
        [filters.range]
                enabled = true
                field = "stars" # field name to validate
                min = 0         # minimum value
                max = 5         # maximal value
        [filters.url]
                enabled = true  # filter is enabled
                field = "uri"   # field name to validate
# aggregators, available aggregators: sorting
[aggregators]
        [aggregators.sorting]
                enabled = true  # aggregation is enabled
                field = "stars" # availabe fields are: stars, name, phone

```

## How to use
```
$ make
go get github.com/tools/godep
godep restore
go install ./...
$ dpipe
2017/03/09 18:24:17 DPIPE
2017/03/09 18:24:17 I! registered inputs: [csv]
2017/03/09 18:24:17 I! registered outputs: [json xml]
2017/03/09 18:24:17 I! registered filters: [encodingUTF8 range url]
2017/03/09 18:24:17 I! registered aggregators: [sorting]
2017/03/09 18:24:17 E! invalid hotel data, skipping
2017/03/09 18:24:17 I! finished processing, stats:
2017/03/09 18:24:17 I! failed to write:		 0
2017/03/09 18:24:17 I! succeed to write:	 7998
2017/03/09 18:24:17 I! validation fails:	 1
2017/03/09 18:24:17 I! received:		 4000
2017/03/09 18:24:17 I! aggregated:		 3999
2017/03/09 18:24:17 I! failed aggreations:	 0
2017/03/09 18:24:17 I! aggreation errors:	 0
```

## How to use with Docker
1. Build a binary:
```make build-for-docker```
2. Build a docker image
```make build-docker-image```
After that **dpipe** will be available in your docker images list.
3. Run, mount **data** volume where input files are kept
```
$ sudo docker run -i -v /data:/app/data dpipe
2017/03/09 18:24:17 DPIPE
2017/03/09 18:24:17 I! registered inputs: [csv]
2017/03/09 18:24:17 I! registered outputs: [json xml]
2017/03/09 18:24:17 I! registered filters: [encodingUTF8 range url]
2017/03/09 18:24:17 I! registered aggregators: [sorting]
2017/03/09 18:24:17 E! invalid hotel data, skipping
2017/03/09 18:24:17 I! finished processing, stats:
2017/03/09 18:24:17 I! failed to write:		 0
2017/03/09 18:24:17 I! succeed to write:	 7998
2017/03/09 18:24:17 I! validation fails:	 1
2017/03/09 18:24:17 I! received:		 4000
2017/03/09 18:24:17 I! aggregated:		 3999
2017/03/09 18:24:17 I! failed aggreations:	 0
2017/03/09 18:24:17 I! aggreation errors:	 0
```

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
