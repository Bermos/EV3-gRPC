import buggy_pb2
import buggy_pb2_grpc
from EV3 import EV3


class Buggy(EV3):

    def __init__(self, host_address: str, port: int = 9000):
        super().__init__(host_address, port)

        self.motors_stub = buggy_pb2_grpc.MotorsStub(self.channel)
        self.sensors_stub = buggy_pb2_grpc.SensorsStub(self.channel)

        self.buggy_empty = buggy_pb2.Empty

    def gyro(self, reset: bool = False) -> int:
        """
        :param reset: True to reset the gyro angle to 0 after reading, default False
        :return: gyro sensor angle in degrees
        """
        gyro_value = self.sensors_stub.Gyro(self.buggy_empty)

        if reset:
            self.gyro_reset()

        return gyro_value

    def gyro_reset(self) -> None:
        """ Reset the gyro angle to 0 degrees """
        self.sensors_stub.GyroReset(self.buggy_empty)

    def left(self, speed: int = 10, ramp_up: int = 5000, ramp_down: int = 5000, stop: str = 'hold',
             reset: bool = False) -> int:
        """
        Set left motor parameters and return tacho count
        :param speed: in percent of max speed, default 10%
        :param ramp_up: in milliseconds [0 .. 60'000] default 5000ms
        :param ramp_down: in milliseconds [0 .. 60'000] default 5000ms
        :param stop: stop action as string [coast|brake|hold] default 'hold'
        :param reset: true resets motor tacho count to 0
        :return: motor tacho count
        """
        params = buggy_pb2.MotorParams(speed=speed, ramp_up=ramp_up, ramp_down=ramp_down, stop=stop, reset=reset)
        return int(self.motors_stub.Left(params).position)

    def right(self, speed: int = 10, ramp_up: int = 5000, ramp_down: int = 5000, stop: str = 'hold',
              reset: bool = False) -> int:
        """
        Set right motor parameters and return tacho count
        :param speed: in percent of max speed (default: 10%)
        :param ramp_up: in milliseconds [0 .. 60'000] (default: 5000ms)
        :param ramp_down: in milliseconds [0 .. 60'000] (default: 5000ms)
        :param stop: stop action as string [coast|brake|hold] (default: 'hold')
        :param reset: true resets motor tacho count to 0 (default: False)
        :return: motor tacho count
        """
        params = buggy_pb2.MotorParams(speed=speed, ramp_up=ramp_up, ramp_down=ramp_down, stop=stop, reset=reset)
        return int(self.motors_stub.Right(params).position)

    def on(self, l_speed: int = 10, r_speed: int = 10) -> None:
        """
        :param l_speed: as percentage of max speed (default: 10%)
        :param r_speed: as percentage of max speed (default: 10%)
        """
        params = buggy_pb2.OnParams(l_speed=l_speed, r_speed=r_speed)
        self.motors_stub.On(params)

    def on_for_degrees(self, l_speed: int = 10, r_speed: int = 10, degrees: int = 360, brake: bool = False) -> None:
        """
        Turn on both motors for the given degrees of rotation.
        If the left speed is not equal to the right speed (i.e., the robot will
        turn), the motor on the outside of the turn will rotate for the full
        degrees while the motor on the inside will have its requested
        distance calculated according to the expected turn.

        :param l_speed: as percentage of max speed (default: 10%)
        :param r_speed: as percentage of max speed (default: 10%)
        :param degrees: to turn the wheels (default: 360)
        :param brake: true to brake (default: False)
        """
        params = buggy_pb2.OnParams(l_speed=l_speed, r_speed=r_speed, degrees=degrees, brake=brake)
        self.motors_stub.OnForDegrees(params)

    def on_for_rotations(self, l_speed: int = 10, r_speed: int = 10, rotations: int = 1, brake: bool = False) -> None:
        """
        Turn on both motors for the given counts of rotation.
        If the left speed is not equal to the right speed (i.e., the robot will
        turn), the motor on the outside of the turn will rotate for the full
        counts while the motor on the inside will have its requested
        distance calculated according to the expected turn.

        :param l_speed: as percentage of max speed (default: 10%)
        :param r_speed: as percentage of max speed (default: 10%)
        :param rotations: to turn the wheels (default: 1)
        :param brake: true to brake (default: False)
        """
        params = buggy_pb2.OnParams(l_speed=l_speed, r_speed=r_speed, rotations=rotations, brake=brake)
        self.motors_stub.OnForRotations(params)

    def on_for_seconds(self, l_speed: int = 10, r_speed: int = 10, seconds: int = 1, brake: bool = False) -> None:
        """
        Turn on both motors for the given duration in seconds.

        :param l_speed: as percentage of max speed (default: 10%)
        :param r_speed: as percentage of max speed (default: 10%)
        :param seconds: to turn the wheels (default: 1)
        :param brake: true to brake (default: False)
        """
        params = buggy_pb2.OnParams(l_speed=l_speed, r_speed=r_speed, seconds=seconds, brake=brake)
        self.motors_stub.OnForRotations(params)

    def sonic(self) -> int:
        """ :return: sonic sensor distance in centimeters """
        return self.sensors_stub.Sonic(self.buggy_empty)

    def stop(self, brake: bool = False):
        """ Stop both motors
        :param brake: true to brake (default: False)
        """
        params = buggy_pb2.StopParams(brake=brake)
        self.motors_stub.Stop(params)

    def wait_until_not_moving(self) -> bool:
        """ Wait until both motors stop moving
        :return: true if the motors stopped moving before reaching the timeout
        """
        return self.motors_stub.WaitUntilNotMoving(self.buggy_empty)
