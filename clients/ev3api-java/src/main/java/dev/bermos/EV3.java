package dev.bermos;

import dev.bermos.proto.ev3.*;
import io.grpc.Grpc;
import io.grpc.InsecureChannelCredentials;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

public class EV3 {
    private final SoundGrpc.SoundBlockingStub soundBlockingStub;
    private final PowerGrpc.PowerBlockingStub powerBlockingStub;
    private final ButtonGrpc.ButtonBlockingStub buttonBlockingStub;
    private final LedGrpc.LedBlockingStub ledBlockingStub;

    private final Empty empty = Empty.newBuilder().build();

    public EV3(String host) {
        ManagedChannel channel = ManagedChannelBuilder.forAddress(host, 9000).usePlaintext().build();

        this.soundBlockingStub = SoundGrpc.newBlockingStub(channel);
        this.powerBlockingStub = PowerGrpc.newBlockingStub(channel);
        this.buttonBlockingStub = ButtonGrpc.newBlockingStub(channel);
        this.ledBlockingStub = LedGrpc.newBlockingStub(channel);
    }

    public void beep() {
        soundBlockingStub.beep(empty);
    }

    public boolean button() {
        return buttonBlockingStub.pressed(empty).getPressed();
    }

    public int current() {
        return (int) powerBlockingStub.current(empty).getCurrent();
    }

    public void flash() {
        EV3Led ledRequest = EV3Led.newBuilder().build();
        ledBlockingStub.flash(ledRequest);
    }

    public void flash(String color) {
        EV3Led ledRequest = EV3Led.newBuilder().setColor(color).build();
        ledBlockingStub.flash(ledRequest);
    }

    public void led(String color) {
        EV3Led ledRequest = EV3Led.newBuilder().setColor(color).build();
        ledBlockingStub.flash(ledRequest);
    }

    public void led(String side, String color) {
        EV3Led ledRequest = EV3Led.newBuilder().setSide(side).setColor(color).build();
        ledBlockingStub.flash(ledRequest);
    }

    public void led_off() {
        EV3Led ledRequest = EV3Led.newBuilder().setColor("off").build();
        ledBlockingStub.flash(ledRequest);
    }

    public int max_voltage() {
        return (int) powerBlockingStub.maxVoltage(empty).getMaxVoltage();
    }

    public int min_voltage() {
        return (int) powerBlockingStub.minVoltage(empty).getMinVoltage();
    }

    public void play_tone(int frequency, int durationMs) {
        Tone tone = Tone.newBuilder().setFrequency(frequency).setDurationMs(durationMs).build();
        soundBlockingStub.playTone(tone);
    }

    public void speak(String content) {
        Text text= Text.newBuilder().setContent(content).build();
        soundBlockingStub.speak(text);
    }

    public String technology() {
        return powerBlockingStub.technology(empty).getTechnology();
    }

    public int voltage() {
        return (int) powerBlockingStub.voltage(empty).getVoltage();
    }
}

