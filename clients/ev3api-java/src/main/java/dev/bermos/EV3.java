package dev.bermos;

import dev.bermos.proto.ev3.*;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

public class EV3 {
    ManagedChannel channel;
    private final SoundGrpc.SoundBlockingStub soundBlockingStub;
    private final PowerGrpc.PowerBlockingStub powerBlockingStub;
    private final ButtonGrpc.ButtonBlockingStub buttonBlockingStub;
    private final LedGrpc.LedBlockingStub ledBlockingStub;

    private final Empty empty = Empty.newBuilder().build();

    public EV3(String host) {
        channel = ManagedChannelBuilder.forAddress(host, 9000).usePlaintext().build();

        this.soundBlockingStub = SoundGrpc.newBlockingStub(channel);
        this.powerBlockingStub = PowerGrpc.newBlockingStub(channel);
        this.buttonBlockingStub = ButtonGrpc.newBlockingStub(channel);
        this.ledBlockingStub = LedGrpc.newBlockingStub(channel);
    }

    public EV3(String host, int port) {
        channel = ManagedChannelBuilder.forAddress(host, port).usePlaintext().build();

        this.soundBlockingStub = SoundGrpc.newBlockingStub(channel);
        this.powerBlockingStub = PowerGrpc.newBlockingStub(channel);
        this.buttonBlockingStub = ButtonGrpc.newBlockingStub(channel);
        this.ledBlockingStub = LedGrpc.newBlockingStub(channel);
    }

    /**
     * Let the robot beep
     */
    public void beep() {
        soundBlockingStub.beep(empty);
    }

    /**
     * @return true if a button was pressend in the last 3 seconds
     */
    public boolean button() {
        return buttonBlockingStub.pressed(empty).getPressed();
    }

    /**
     * @return measured battery current in mirco-amps
     */
    public int current() {
        return (int) powerBlockingStub.current(empty).getCurrent();
    }

    /**
     * Flashes the LEDs
     */
    public void flash() {
        EV3Led ledRequest = EV3Led.newBuilder().build();
        ledBlockingStub.flash(ledRequest);
    }

    /**
     * Flashes the LEDs in the given color
     * @param color to use [off|red|yellow|orange|amber|lime|green]
     */
    public void flash(String color) {
        EV3Led ledRequest = EV3Led.newBuilder().setColor(color).build();
        ledBlockingStub.flash(ledRequest);
    }

    /**
     * Set the given color for both LEDs
     * @param color to set [off|red|yellow|orange|amber|lime|green]
     */
    public void led(String color) {
        EV3Led ledRequest = EV3Led.newBuilder().setColor(color).build();
        ledBlockingStub.led(ledRequest);
    }

    /**
     * Set the given color for the given LED(s)
     * @param side to set [left|right|all]
     * @param color to set [off|red|yellow|orange|amber|lime|green]
     */
    public void led(String side, String color) {
        EV3Led ledRequest = EV3Led.newBuilder().setSide(side).setColor(color).build();
        ledBlockingStub.led(ledRequest);
    }

    /**
     * Turns off all LEDs
     */
    public void led_off() {
        ledBlockingStub.ledOff(empty);
    }

    /**
     * @return max battery voltage in micro-volts
     */
    public int max_voltage() {
        return (int) powerBlockingStub.maxVoltage(empty).getMaxVoltage();
    }

    /**
     * @return min battery voltage in micro-volts
     */
    public int min_voltage() {
        return (int) powerBlockingStub.minVoltage(empty).getMinVoltage();
    }

    /**
     * Play a tone
     * @param frequency of the tone in herz
     * @param durationMs of the tone in ms
     */
    public void play_tone(int frequency, int durationMs) {
        Tone tone = Tone.newBuilder().setFrequency(frequency).setDurationMs(durationMs).build();
        soundBlockingStub.playTone(tone);
    }

    /**
     * Speak the given content
     * @param content to be spoken
     */
    public void speak(String content) {
        Text text= Text.newBuilder().setContent(content).build();
        soundBlockingStub.speak(text);
    }

    /**
     * @return battery technology
     */
    public String technology() {
        return powerBlockingStub.technology(empty).getTechnology();
    }

    /**
     * @return measured battery voltage in micro-volts
     */
    public int voltage() {
        return (int) powerBlockingStub.voltage(empty).getVoltage();
    }
}

