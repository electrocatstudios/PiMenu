# PiMenu
Make Raspberry Pi Menus easily to display on a screen. Note: This needs to be compiled on a Raspberry Pi due to the libraries not being available for use on other systems. If anyone knows how to install them on different operating systems to allow cross-compilation then please get in touch. 

## Prerequisites for development
You'll need to install the following prerequisites before getting the OpenVG library

```
sudo apt-get install libjpeg8-dev indent libfreetype6-dev ttf-dejavu-core -y
sudo apt-get install libraspberrypi-dev raspberrypi-kernel-headers -y
```

Then install the dependencies

```
go get golang.org/x/exp/io/i2c
go get github.com/ajstarks/openvg
```

## Setting up a custom screen
To set up a custom screen we need to create json files which will represent the menu items on the device. The only required one is main.json which will describe the landing menu page. 

## Generating Python Code
You will need the python grpc tools

```
pip3 install grpcio-tools
```

## Work in Progress
No support is offered at the moment for the code because it's still in development. We will update this when it's ready to use

