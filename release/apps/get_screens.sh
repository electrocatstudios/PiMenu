
#!/bin/sh

# Get the example screen files
curl https://menecs.electrocatstudios.com/zips/screens.zip --output screens.zip
unzip -o screens.zip
mkdir screens
mv *.json screens/
rm screens.zip

curl https://menecs.electrocatstudios.com/zips/menecs_images.zip --output menecs_images.zip
unzip -o menecs_images.zip
rm menecs_images.zip

sudo systemctl restart pimenu.service