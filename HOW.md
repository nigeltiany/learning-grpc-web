docker build -t echo-server .

docker tag echo-server gcr.io/confab-cloud/echo-server:v1.0.2

docker push gcr.io/confab-cloud/echo-server:v1.0.2