build:
	GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap main.go

zip:
	zip bootstrap.zip bootstrap

docker-build:
	docker build --platform linux/amd64 -t lambda-functions:latest .

run:
	docker run \
	-v ~/.aws-lambda-rie:/aws-lambda \
	-p 9000:8080 \
	--entrypoint /aws-lambda/aws-lambda-rie \
	lambda-functions:latest ./main
