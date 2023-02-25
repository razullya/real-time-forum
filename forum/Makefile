all:
	docker image build -f Dockerfile . -t forum
	docker run -p 8080:8080 --name cont forum
	
