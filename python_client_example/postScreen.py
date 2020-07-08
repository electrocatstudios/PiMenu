import grpc

import screenservice_pb2
import screenservice_pb2_grpc

if __name__ == "__main__":
    print("Starting to send screen")
    channel = grpc.insecure_channel("192.168.0.100:7777")
    stub = screenservice_pb2_grpc.ScreenServerStub(channel)
    payload = screenservice_pb2.ScreenRequest()

    payload.line1.line_type = "text" 
    payload.line1.line_value =  "another message sent"
 
    payload.line2.line_type = "gif"
    payload.line2.line_value = "0;0;0;0;gifs/cat.gif"

    payload.timeout.length = 10
    payload.timeout.showtimeout = 5
    payload.timeout.returnscreen = "menu"
    
    ret = stub.SendScreen(payload)
    print("Finished sending screen")