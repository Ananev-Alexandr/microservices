docker build -t my-golang-app .
docker run -it --rm --name my-running-app -p 50051:50051 my-golang-app
