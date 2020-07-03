# PiMenu
Make Raspberry Pi Menus easily to display on a screen. Note: This needs to be compiled on a Raspberry Pi due to the libraries not being available for use on other systems. If anyone knows how to install them on different operating systems to allow cross-compilation then please get in touch. 

## Running the application
In the release folder you will find a binary (called Menecs) which is the main application. You'll need to customise the json files in the screens folder (and the folder will need to be next to the binary, along with the images folder). You'll also need to install and run the fbcp application https://github.com/juj/fbcp-ili9341. Every time a screen is changed the json file is reread so you can hot-swap json files. 

## Prerequisites for development
You'll need to install the following prerequisites before getting the OpenVG library

```
sudo apt install libjpeg8-dev indent libfreetype6-dev ttf-dejavu-core -y
sudo apt install libraspberrypi-dev raspberrypi-kernel-headers -y
sudo apt install golang-google-grpc-dev -y
sudo apt install python3-grpcio python3-grpc-tools python3-protobuf -y
sudo apt install -y libavdevice-dev libavfilter-dev libswscale-dev libavcodec-dev libavformat-dev libswresample-dev libavutil-dev
```

Then install the dependencies

```
go get golang.org/x/exp/io/i2c
go get github.com/ajstarks/openvg
go get github.com/golang/protobuf/proto
go get golang.org/x/net/context
go get google.golang.org/grpc
go get github.com/nfnt/resize

go get github.com/giorgisio/goav/avcodec
go get github.com/giorgisio/goav/avformat
go get github.com/giorgisio/goav/avutil

go get github.com/electrocatstudios/FTXXXX_Touchscreen_Driver
```
## Setting up a Python Client

If you only wish to consume the screen service then you will need to install the following prerequisites on the Pi:
```
sudo apt install python3-grpcio python3-grpc-tools python3-protobuf -y
```
Note: While it is available as a Pip wheel it is not recommended as it takes significantly longer, and when tested on a Pi zero appeared to grind to a halt (process abandoned after 2 hours - and a lot of heat coming off the Pi!)

An example script is included in the python_client_example folder. You will need to copy the screen_service_pb2 files into any project you instigate as these are the descriptors of the classes you'll need. The following script is a simple hello world example:

```
import grpc
import screenservice_pb2
import screenservice_pb2_grpc

channel = grpc.insecure_channel("localhost:7777")
stub = screenservice_pb2_grpc.ScreenServerStub(channel)
payload = screenservice_pb2.ScreenRequest()

payload.line1 = "test sending message"
payload.line2 = "another line of message"
# etc... (up to 5 lines available)
payload.length = 5 # How long (in seconds) to show - must be integer
payload.showCountdown = 2 # When to show the countdown timer in seconds remaining

stub.SendScreen(payload)
```

If the ScreenService app is running on the same device it should connect with no problem and the new screen should be visible for 5 seconds. Replace "localhost" with any IP address for remote connection.


## Setting up a custom screen
To set up a custom screen we need to create json files which will represent the menu items on the device. The only required one is main.json which will describe the landing menu page. 

## Generating Python Code
You will need the python grpc tools

```
pip3 install grpcio-tools
```

## Work in Progress
No support is offered at the moment for the code because it's still in development. We will update this when it's ready to use

## TODO 
We are working on the following:

Support for other file types
Adding keyboard and pin entry screens
Creating an online web app for generating screens
Allowing remote administration of images and screens