import buggy_pb2
import buggy_pb2_grpc
from EV3 import EV3


class Buggy(EV3):

    def __init__(self, host_address: str, port: int = 9000):
        super().__init__(host_address, port)

        self.motors_stub = buggy_pb2_grpc.MotorsStub(self.channel)
        self.sensors_stub = buggy_pb2_grpc.SensorsStub(self.channel)

        self.buggy_empty = buggy_pb2.Empty

    def gyro(self) -> int:
        """ :return: gyro sensor angle in degrees """
        return int(self.sensors_stub.Gyro(self.buggy_empty))
