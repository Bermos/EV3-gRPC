package dev.bermos;

import org.junit.Before;
import org.junit.Test;

import static org.junit.Assert.*;

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
}
