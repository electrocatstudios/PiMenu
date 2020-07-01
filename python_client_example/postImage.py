from picamera import PiCamera
import time 
from io import BytesIO
from PIL import Image

import grpc

import screenservice_pb2
import screenservice_pb2_grpc

def image_to_byte_array(image):
  imgByteArr = BytesIO()
  image.save(imgByteArr, format='JPEG')
  imgByteArr = imgByteArr.getvalue()
  return imgByteArr

if __name__ == "__main__":
    print("Starting app")
    
    camera = PiCamera()
    camera.start_preview()
    channel = grpc.insecure_channel("192.168.0.100:7777")
    
    stub = screenservice_pb2_grpc.ScreenServerStub(channel)
    payload = screenservice_pb2.ScreenImage()

    time.sleep(2) # Suggested to let the camera settle after start_preview()

    stream = BytesIO()
    filename = "latest.jpg"

    print("Starting capturing")
    
    for pic in camera.capture_continuous(stream, 'jpeg', use_video_port=False, burst=True):
        image = Image.open(stream)
        image.save(filename, format='JPEG')

        with open(filename, mode='rb') as file:
            fileContent = file.read()
        
        payload.image_data = fileContent
        payload.height = 480
        payload.width = 800

        stream.truncate()
        stream.seek(0)

        # For benchmarking    
        # cur_time = time.time()

        ret = stub.SendImage(payload)
        if ret.status != "ok":
            print("ERROR during send")

        # elapsed = time.time() - cur_time
        # print("{} elapsed".format(elapsed))