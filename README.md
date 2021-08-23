# beamer
Automatically generate job parameter options from [GCP Dataflow Templates](https://github.com/GoogleCloudPlatform/DataflowTemplates)


## Usage

```
usage: beamer <command> <args>

These are the only two commands available:
	init    setup beamer for the current directory
        gen     generates a new job config
        run     runs the generated job config on GCP

Examples:
        - Generates a job config for BigQueryToDatastore
        $ beamer gen BigQueryToDatastore

        - Run the job config for BigQueryToDatastore migration
        $ beamer run BigQueryToDatastore
```

## Installation

```shell
$ make install
```
