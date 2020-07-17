# Replace with the values on your pi
# this allows development on a different machine
# Requires access keys to be set up on the devices
PI_ADDR=
FOLDER=/home/pi/menecs
# FOLDER=/home/pi/camera_send

ssh pi@$PI_ADDR "sudo systemctl stop pimenu.service"
scp release/Menecs pi@$PI_ADDR:$FOLDER/

# scp -r screens/rpi_touch/*.json pi@$PI_ADDR:$FOLDER/screens/
scp -r web/ pi@$PI_ADDR:$FOLDER/

ssh pi@$PI_ADDR "sudo systemctl start pimenu.service"
# scp -r python_client_example/ pi@$PI_ADDR:$FOLDER/

# scp -r images/ pi@$PI_ADDR:$FOLDER/
# scp -r service/ pi@$PI_ADDR:$FOLDER/
