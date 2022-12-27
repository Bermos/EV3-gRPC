package dev.bermos;

import org.junit.Before;
import org.junit.Test;

public class BuggyTest {
    private static Buggy buggy;

    @Before
    public void before() {
        buggy = new Buggy("10.0.100.98");
    }

    @Test
    public void gyro() {
        System.out.println(buggy.gyro());
    }

    @Test
    public void gyro_reset() {
        buggy.gyro_reset();
    }

    @Test
    public void left() {
        buggy.left(20, 5000, 5000, Buggy.StopAction.BRAKE, false);
    }

    @Test
    public void right() {
        buggy.right(-20, 5000, 5000, Buggy.StopAction.BRAKE, false);
    }

    @Test
    public void on() {
        buggy.on(5, 10);
    }

    @Test
    public void onForDegrees() {
        buggy.onForDegrees(10, -10, 20);
    }

    @Test
    public void onForRotations() {
        buggy.onForRotations(-5, 5, 2);
    }

    @Test
    public void onForSeconds() {
        buggy.onForSeconds(10, 5, 5);
    }

    @Test
    public void sonic() {
        System.out.println(buggy.sonic());
    }

    @Test
    public void stop() {
        buggy.stop(true);
    }

    @Test
    public void waitUntilNotMoving() {
        System.out.println(buggy.waitUntilNotMoving());
    }
}
