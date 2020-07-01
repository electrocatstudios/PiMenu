import grpc

import screenservice_pb2
import screenservice_pb2_grpc

if __name__ == "__main__":
    print("Starting to send screen")
    channel = grpc.insecure_channel("192.168.0.100:7777")
    stub = screenservice_pb2_grpc.ScreenServerStub(channel)
    payload = screenservice_pb2.ScreenRequest()
    line1 = screenservice_pb2.Line()
    line1.line_type = "text"
    line1.line_value = "another message sent"
    payload.line1.line_type = "text" 
    payload.line1.line_value =  "another message sent"

    # timeout = screenservice_pb2.Timeout()
    payload.timeout.length = 10
    payload.timeout.showtimeout = 5
    payload.timeout.returnscreen = "menu"
    
    ret = stub.SendScreen(payload)
    print("Finished sending screen")