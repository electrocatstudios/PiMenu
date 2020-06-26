# Replace with the values on your pi
# this allows development on a different machine
# Requires access keys to be set up on the devices
PI_ADDR=192.168.0.100
FOLDER=/home/pi/go/src/github.com/electrocatstudios/PiMenu

scp *.go pi@$PI_ADDR:$FOLDER/
scp Makefile pi@$PI_ADDR:$FOLDER/
scp -r screens/ pi@$PI_ADDR:$FOLDER/
scp -r bindings/ pi@$PI_ADDR:$FOLDER/

scp -r screenservice/ pi@$PI_ADDR:$FOLDER/

## TODO: Add in auto build after deployment - need to stop service first
# ssh  pi@$PI_ADDR "cd $FOLDER && make"
ssh  pi@$PI_ADDR "cd $FOLDER && make python"