import grpc

import screenservice_pb2
import screenservice_pb2_grpc

if __name__ == "__main__":
    print("Starting to send screen")
    channel = grpc.insecure_channel("192.168.0.100:7777")
    stub = screenservice_pb2_grpc.ScreenServerStub(channel)
    payload = screenservice_pb2.ScreenRequest()
   
    payload.line1 = "test sending message"
    payload.line2 = "another line of message"
    payload.length = 5
    payload.showCountdown = 2

    ret = stub.SendScreen(payload)
    print("Finished sending screen")