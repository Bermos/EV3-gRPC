package dev.bermos;

import dev.bermos.proto.buggy.*;

public class Buggy extends EV3 {
    public enum StopAction {
        BRAKE("brake"),
        COAST("coast"),
        HOLD("hold");

        private final String name;

        StopAction(String stopAction) {
            this.name = stopAction;
        }

        @Override
        public String toString() {
            return this.name;
        }
    }

    private final MotorsGrpc.MotorsBlockingStub motorsBlockingStub;

    private final SensorsGrpc.SensorsBlockingStub sensorsBlockingStub;
    private final Empty empty = Empty.newBuilder().build();

    public Buggy(String host) {
        super(host);

        motorsBlockingStub = MotorsGrpc.newBlockingStub(channel);
        sensorsBlockingStub = SensorsGrpc.newBlockingStub(channel);
    }

    public Buggy(String host, int port) {
        super(host, port);

        motorsBlockingStub = MotorsGrpc.newBlockingStub(channel);
        sensorsBlockingStub = SensorsGrpc.newBlockingStub(channel);
    }

    /**
     * @return gyro sensor angle in degrees
     */
    public int gyro() {
        return gyro(false);
    }

    /**
     * @param reset true to reset the gyro angle to 0 after reading
     * @return gyro sensor angle in degrees
     */
    public int gyro(boolean reset) {
        int gyroValue = (int) sensorsBlockingStub.gyro(empty).getNumValue();

        if (reset) {
            gyro_reset();
        }

        return gyroValue;
    }

    /**
     * Reset the gyro angle to 0 degrees
     */
    public void gyro_reset() {
        sensorsBlockingStub.gyroReset(empty);
    }

    /**
     * Sets left motor parameters and return tacho count
     * @param speed in percent of max speed
     * @param rampUp in milliseconds
     * @param rampDown in milliseconds
     * @param stopAction
     * @param reset true resets motor tacho count to 0
     * @return motor tacho count
     */
    public int left(int speed, int rampUp, int rampDown, StopAction stopAction, boolean reset) {
        MotorParams params = MotorParams.newBuilder()
                .setSpeed(speed)
                .setRampUp(rampUp)
                .setRampDown(rampDown)
                .setStop(stopAction.toString())
                .setReset(reset).build();
        MotorState state = motorsBlockingStub.left(params);
        return state.getPosition();
    }

    /**
     * Sets right motor parameters and return tacho count
     * @param speed in percent of max speed
     * @param rampUp in milliseconds
     * @param rampDown in milliseconds
     * @param stopAction
     * @param reset true resets motor tacho count to 0
     * @return motor tacho count
     */
    public int right(int speed, int rampUp, int rampDown, StopAction stopAction, boolean reset) {
        MotorParams params = MotorParams.newBuilder()
                .setSpeed(speed)
                .setRampUp(rampUp)
                .setRampDown(rampDown)
                .setStop(stopAction.toString())
                .setReset(reset).build();
        MotorState state = motorsBlockingStub.right(params);
        return state.getPosition();
    }

    /**
     * Turn on both motors
     * @param lSpeed as percentage of max speed
     * @param rSpeed as percentage of max speed
     */
    public void on(int lSpeed, int rSpeed) {
        OnParams params = OnParams.newBuilder().setLSpeed(lSpeed).setRSpeed(rSpeed).build();
        motorsBlockingStub.on(params);
    }

    /**
     * Turn on both motors for the given degrees of rotation.
     * If the left speed is not equal to the right speed (i.e., the robot will
     * turn), the motor on the outside of the turn will rotate for the full
     * degrees while the motor on the inside will have its requested
     * distance calculated according to the expected turn.
     *
     * @param lSpeed as percentage of max speed
     * @param rSpeed as percentage of max speed
     * @param degrees to turn the wheels
     * @param brake true to brake
     */
    public void onForDegrees(int lSpeed, int rSpeed, int degrees, boolean brake) {
        OnParams params = OnParams.newBuilder()
                .setLSpeed(lSpeed)
                .setRSpeed(rSpeed)
                .setDegrees(degrees)
                .setBreak(brake).build();
        motorsBlockingStub.onForDegrees(params);
    }

    /**
     * Turn on both motors for the given counts of rotation.
     * If the left speed is not equal to the right speed (i.e., the robot will
     * turn), the motor on the outside of the turn will rotate for the full
     * rotations while the motor on the inside will have its requested
     * distance calculated according to the expected turn.
     *
     * @param lSpeed as percentage of max speed
     * @param rSpeed as percentage of max speed
     * @param rotations to turn the wheels
     * @param brake true to brake
     */
    public void onForRotations(int lSpeed, int rSpeed, int rotations, boolean brake) {
        OnParams params = OnParams.newBuilder()
                .setLSpeed(lSpeed)
                .setRSpeed(rSpeed)
                .setRotations(rotations)
                .setBreak(brake).build();
        motorsBlockingStub.onForRotations(params);
    }

    /**
     * Turn on both motors for the given duration in seconds.
     *
     * @param lSpeed as percentage of max speed
     * @param rSpeed as percentage of max speed
     * @param seconds to turn the wheels
     * @param brake true to brake
     */
    public void onForSeconds(int lSpeed, int rSpeed, int seconds, boolean brake) {
        OnParams params = OnParams.newBuilder()
                .setLSpeed(lSpeed)
                .setRSpeed(rSpeed)
                .setSeconds(seconds)
                .setBreak(brake).build();
        motorsBlockingStub.onForSeconds(params);
    }

    /**
     * @return sonic sensor distance in centimeters
     */
    public int sonic() {
        return (int) sensorsBlockingStub.sonic(empty).getNumValue();
    }

    /**
     * Stop both motors
     * @param brake true to brake
     */
    public void stop(boolean brake) {
        StopParams params = StopParams.newBuilder().setBreak(brake).build();
        motorsBlockingStub.stop(params);
    }

    /**
     * Wait until both motors stop moving
     * @return true if the motors stopped moving before reaching the timeout
     */
    public boolean waitUntilNotMoving() {
        return motorsBlockingStub.waitUntilNotMoving(empty).getNotMoving();
    }
}
