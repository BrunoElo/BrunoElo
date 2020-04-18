// Interest calculator that displays an array of object containing rate and interest properties

// Declared objects containing principle and time properties
let interestObj1 = { principal: 2500, time: 1.8 };
let interestObj2 = { principal: 1000, time: 5 };
let interestObj3 = { principal: 3000, time: 1 };
let interestObj4 = { principal: 2000, time: 3 };

// A data array of the four declared objects
let data = [interestObj1, interestObj2, interestObj3, interestObj4];

// Function that calculates interest
function interestCalculator(data) {
    // Applicable rates for the objects in the array is determined and included in the objects
    for (let i in data) {
        if (data[i].principal >= 2500 && (data[i].time > 1 && data[i].time < 3)) {
            data[i].rate = 3;
        } else if (data[i].principal >= 2500 && data[i].time >= 3) {
            data[i].rate = 4;
        } else if (data[i].principal < 2500 || data[i].time <= 1) {
            data[i].rate = 2;
        } else {
            data[i].rate = 1;
        }
    }
    // Interest for each object is calculated using their corresponding property values
    // and then included in the objects in the data array
    let interestData = [];
    for (let i in data) {
        data[i].interest = (data[i].principal * data[i].rate * data[i].time) / 100
        // Contents of the data array (objects) are transferred to the interestData array
        interestData.push(data[i]);
    }
    console.log(interestData);
    return interestData
}
interestCalculator(data);


// Interest calculator that displays an array of object containing rate and interest properties
// This method has array disconnection to avoid modification of another array if one is modified (see line 62)
/* Redundant declarations...already done above
// Declared objects containing principle and time properties
let interestObj1 = { principal: 2500, time: 1.8 };
let interestObj2 = { principal: 1000, time: 5 };
let interestObj3 = { principal: 3000, time: 1 };
let interestObj4 = { principal: 2000, time: 3 };

// A data array of the four declared objects
let data = [interestObj1, interestObj2, interestObj3, interestObj4];*/

// Function that calculates interest
function interestCalculator(data) {

    /*
    // Using this method keeps the referencing between the arrays interestData and data which modifies the other if one is modified.
    let interestData = []
    for (let i in data) {
        interestData.push(data[i]);
    }*/

    // This method disconnects the referencing
    let interestData = JSON.parse(JSON.stringify(data));

    // Applicable rates for the objects in the array is determined and included in the objects
    for (let i in interestData) {
        if (interestData[i].principal >= 2500 && (interestData[i].time > 1 && interestData[i].time < 3)) {
            interestData[i].rate = 3;
        } else if (interestData[i].principal >= 2500 && interestData[i].time >= 3) {
            interestData[i].rate = 4;
        } else if (interestData[i].principal < 2500 || interestData[i].time <= 1) {
            interestData[i].rate = 2;
        } else {
            interestData[i].rate = 1;
        }
    }
    // Interest for each object is calculated using their corresponding property values
    // and then included in the objects in the interestData array

    for (let i in interestData) {
        interestData[i].interest = (interestData[i].principal * interestData[i].rate * interestData[i].time) / 100
    }
    console.log(interestData);

    return interestData
}
interestCalculator(data);