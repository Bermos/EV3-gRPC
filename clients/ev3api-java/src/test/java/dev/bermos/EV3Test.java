package dev.bermos;

import org.junit.Before;
import org.junit.Test;

public class EV3Test {
    private static EV3 ev3;

    @Before
    public void before() {
        ev3 = new EV3("10.0.100.98");
    }

    @Test
    public void beep() {
        ev3.beep();
    }

    @Test
    public void button() {
        System.out.println(ev3.button());
    }

    @Test
    public void current() {
        System.out.println(ev3.current());
    }

    @Test
    public void flash() {
        ev3.flash();
    }

    @Test
    public void flash_color() {
        ev3.flash("green");
    }

    @Test
    public void led() {
        ev3.led("left", "green");
    }

    @Test
    public void led_off() {
        ev3.led_off();
    }

    @Test
    public void max_voltage() {
        System.out.println(ev3.max_voltage());
    }

    @Test
    public void min_voltage() {
        System.out.println(ev3.min_voltage());
    }

    @Test
    public void play_tone() {
        ev3.play_tone(220, 1000);
    }

    @Test
    public void speak() {
        ev3.speak("Hello, I am a robot");
    }

    @Test
    public void technology() {
        System.out.println(ev3.technology());
    }

    @Test
    public void voltage() {
        System.out.println(ev3.voltage());
    }
}
