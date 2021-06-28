# GO lambda API 

Requirements:
- AWS account with API Gateway and Lambda access.

First, setup GO from [Here](https://golang.org/doc/install)

For local run,

`copy the .env.example to .env and set necessary configurations and then run:`
 
```sh
go get 
go run .
```

> Note: "go get" fetches all the required dependencies and "go run ." searches for the entrypoint/main function in the current directory which is 'main.go' in this case and runs it.

with this, the server will be available on 'localhost:8081'

`For dev/prod server deployment,

`From AWS Lambda Console, create a lambda preferably named '{stage}_{projectname}' and set the handler to 'main'`

`From AWS Gateway Console, create a REST API with lambda proxy integration, link it with the lambda then deploy the required endpoints.`

After this cloud setup, run:

```sh
go build ./main.go
zip -j ./build.zip ./main
aws lambda update-function-code --zip-file=fileb://build.zip \
   --function-name="CREATED_FUNCTION_NAME"
```

or Simply

```sh
sh ./scripts/deploy.sh
```

So, that's how we create a binary build of the project which gives out a 'main' file which will be the handler we previously mentioned while creating the lambda.
We zip the compiled files and push to the function.

Then, the function can be invoked from the Gateway endpoint got from the deployment of the APIs previously.
