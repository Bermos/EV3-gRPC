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

    def beep(self) -> None:
        """ The EV3 will beep """
        self.sound_stub.Beep(ev3_pb2.Empty())

    def button(self) -> bool:
        """ :return: True if a button is currently pressed or has been pressed in the last 3 seconds """
        return self.button_stub.Pressed(ev3_pb2.Empty()).pressed

    def current(self) -> int:
        """ :return: measured battery current in micro ampere """
        return int(self.power_stub.Current(ev3_pb2.Empty()).current)

    def flash(self, color: str = 'amber') -> None:
        """ Flashes the LEDs
        :param color: to use [off|red|yellow|orange|amber|lime|green], default 'amber'
        """
        params = ev3_pb2.EV3Led(color=color)
        self.led_stub.Flash(params)

    def led(self, side: str, color: str) -> None:
        """ Set the given color for the given LED(s)
        :param side: to set as string [left|right|all]
        :param color: to set as string [off|red|yellow|orange|amber|lime|green]
        """
        params = ev3_pb2.EV3Led(side=side, color=color)
        self.led_stub.Led(params)

    def led_off(self) -> None:
        """ Turn off all LEDs """
        self.led_stub.LedOff(ev3_pb2.Empty())

    def max_voltage(self) -> int:
        """ :return: max battery voltage in micro volts """
        return int(self.power_stub.MaxVoltage(ev3_pb2.Empty()).max_voltage)

    def min_voltage(self) -> int:
        """ :return: min battery voltage in micro volts """
        return int(self.power_stub.MinVoltage(ev3_pb2.Empty()).min_voltage)

    def play_tone(self, frequency: int, duration_ms: int) -> None:
        """ Play a tone
        :param frequency: of the tone in herz
        :param duration_ms: of the tone in milliseconds
        """
        params = ev3_pb2.Tone(frequency=frequency, duration_ms=duration_ms)
        self.sound_stub.PlayTone(params)

    def speak(self, content: str) -> None:
        """ Speak the given content
        :param content: to be spoken
        """
        params = ev3_pb2.Text(content=content)
        self.sound_stub.Speak(params)

    def technology(self) -> str:
        """ :return: battery technology """
        return self.power_stub.Technology(ev3_pb2.Empty()).technology

    def voltage(self) -> int:
        """ :return: measured battery voltage in micro volts """
        return int(self.power_stub.Voltage(ev3_pb2.Empty()).voltage) * 1000
