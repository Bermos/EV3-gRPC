import grpc

import ev3_pb2_grpc
import ev3_pb2


class EV3:

    def __init__(self, host_address: str, port: int = 9000):
        self.channel = grpc.insecure_channel(f"{host_address}:{port}")

        self.sound_stub = ev3_pb2_grpc.SoundStub(self.channel)
        self.power_stub = ev3_pb2_grpc.PowerStub(self.channel)
        self.button_stub = ev3_pb2_grpc.ButtonStub(self.channel)
        self.led_stub = ev3_pb2_grpc.LedStub(self.channel)

        self.ev3_empty = ev3_pb2.Empty()

    def beep(self) -> None:
        """
        The EV3 will beep.
        """

        self.sound_stub.Beep(self.ev3_empty)

    def button(self) -> bool:
        return self.button_stub.Pressed(self.ev3_empty).pressed

    def current(self) -> int:
        return int(self.power_stub.Current(self.ev3_empty).current)

    def flash(self, color: str = 'amber') -> None:
        params = ev3_pb2.EV3Led(color=color)
        self.led_stub.Flash(params)

    def led(self, side: str, color: str) -> None:
        params = ev3_pb2.EV3Led(side=side, color=color)
        self.led_stub.Led(params)

    def led_off(self) -> None:
        self.led_stub.LedOff(self.ev3_empty)

    def max_voltage(self) -> int:
        return int(self.power_stub.MaxVoltage(self.ev3_empty).max_voltage)

    def min_voltage(self) -> int:
        return int(self.power_stub.MinVoltage(self.ev3_empty).min_voltage)

    def play_tone(self, frequency: int, duration_ms: int) -> None:
        params = ev3_pb2.Tone(frequency=frequency, duration_ms=duration_ms)
        self.sound_stub.PlayTone(params)

    def speak(self, content: str) -> None:
        params = ev3_pb2.Text(content=content)
        self.sound_stub.Speak(params)

    def technology(self) -> str:
        return self.power_stub.Technology(self.ev3_empty).technology

    def voltage(self) -> int:
        return int(self.power_stub.Voltage(self.ev3_empty).voltage)
