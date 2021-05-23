# A Go Service

## Objective

Just a REST API written in Golang for practice. Modeled heavily on [gorsk](https://github.com/ribice/gorsk). 

This REST API is intended to be deployed to AWS as the user facing component of a larger group of microservices.

For READ operations, the API reads directly from a database. 
For WRITE operations, the API does simple validation before formatting and passing the data along to an AWS Kinesis stream. The stream is the trigger for the [lambda component](https://github.com/konstantinfarrell/go-example-lambda)


## Usage

	make

