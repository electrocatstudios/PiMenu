all:
	protoc -I screenservice/ --go_out=plugins=grpc:screenservice screenservice/screenservice.proto
	go build -o Menecs

python:
	python3 -m grpc_tools.protoc -I screenservice/ --python_out=bindings/. --grpc_python_out=bindings/. screenservice/screenservice.proto

clean:
	rm -rf Menecs