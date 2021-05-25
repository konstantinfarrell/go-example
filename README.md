# A Go Service

## Objective

Just a REST API written in Golang for practice. Modeled heavily on [gorsk](https://github.com/ribice/gorsk). 

This REST API is intended to be deployed to AWS as the user facing component of a larger group of microservices.

For READ operations, the API reads directly from a database. 
For WRITE operations, the API does simple validation before formatting and passing the data along to an AWS Kinesis stream. The stream is the trigger for the [lambda component](https://github.com/konstantinfarrell/go-example-lambda)


## Setup

The following are environment variables necessary to run the API

Database:

	DB_USER
	DB_PASS
	DB_NAME
	DB_PORT
	DB_ADDR

AWS:

	AWS_REGION
	KINESIS_STREAM_NAME
	KINESIS_PARTITION_KEY

The following environment variables are NOT currently required, but will be later on.


## Usage

	make

