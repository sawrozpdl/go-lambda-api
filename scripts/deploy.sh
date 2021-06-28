go build ./main.go

zip -j ./build.zip ./main 

aws lambda update-function-code --zip-file=fileb://build.zip \
   --function-name="${STAGE}_go_api"
