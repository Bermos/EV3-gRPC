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

    public int gyro(boolean reset) {
        int gyroValue = (int) sensorsBlockingStub.gyro(empty).getNumValue();

        if (reset) {
            gyro_reset();
        }

        return gyroValue;
    }

    public void gyro_reset() {
        sensorsBlockingStub.gyroReset(empty);
    }

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

    public void on(int lSpeed, int rSpeed) {
        OnParams params = OnParams.newBuilder().setLSpeed(lSpeed).setRSpeed(rSpeed).build();
        motorsBlockingStub.on(params);
    }

    public void onForDegrees(int lSpeed, int rSpeed) {
        OnParams params = OnParams.newBuilder().setLSpeed(lSpeed).setRSpeed(rSpeed).build();
        motorsBlockingStub.on(params);
    }

    public void onForRotations(int lSpeed, int rSpeed) {
        OnParams params = OnParams.newBuilder().setLSpeed(lSpeed).setRSpeed(rSpeed).build();
        motorsBlockingStub.on(params);
    }

    public void onForSeconds(int lSpeed, int rSpeed) {
        OnParams params = OnParams.newBuilder().setLSpeed(lSpeed).setRSpeed(rSpeed).build();
        motorsBlockingStub.on(params);
    }

    public int sonic() {
        return (int) sensorsBlockingStub.sonic(empty).getNumValue();
    }

    public void stop(boolean brake) {
        StopParams params = StopParams.newBuilder().setBreak(brake).build();
        motorsBlockingStub.stop(params);
    }

    public boolean waitUntilNotMoving() {
        return motorsBlockingStub.waitUntilNotMoving(empty).getNotMoving();
    }
}
