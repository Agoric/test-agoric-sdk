/*---
description: 
flags: [onlyStrict]
---*/

if (!(device?.I2C?.default) || !($TESTMC?.config?.i2c))
    throw new Error("Default I2C device and test harness i2c configuration must be provided")

const I2C = device.I2C.default.io;
const defaultConfiguration = {...device.I2C.default, ...$TESTMC.config.i2c};

let i2c = new I2C(defaultConfiguration);

i2c.close();