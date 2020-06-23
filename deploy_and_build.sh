# Replace with the values on your pi
# this allows development on a different machine
# Requires access keys to be set up on the devices
PI_ADDR=<IPADDRESS>
FOLDER=<PIFOLDER>

scp *.go pi@$PI_ADDR:$FOLDER/
scp Makefile pi@$PI_ADDR:$FOLDER/
scp -r screens/ pi@$PI_ADDR:$FOLDER/

## TODO: Add in auto build after deployment - need to stop service first
ssh  pi@$PI_ADDR "cd /home/pi/menecs/ && make"