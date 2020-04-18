// A function that takes in one parameter and outputs an array that shows
// the divisiility of numbers from 1 to the inputted argument inclusive

function yugioh(integer) {
    // Assigns the range of values from 1 to argument integer (inclusive) to oneToInt
    let oneToInt = Array.from(Array(integer + 1).keys()).slice(1);
    let l = oneToInt.length;
    for (let i = 0; i < l; i++) {

        // for numbers divisible by 2, 2 and 3, 2 and 5, 2, 3 and 5
        if ((oneToInt[i] % 2) == 0) {
            if (oneToInt[i] % 3 == 0) {
                if (oneToInt[i] % 5 == 0) {
                    oneToInt[i] = 'yu-gi-oh';
                } else {
                    oneToInt[i] = 'yu-gi';
                }
            } else if (oneToInt[i] % 5 == 0) {
                oneToInt[i] = 'yu-oh';
            } else {
                oneToInt[i] = 'yu';
            }

            // for numbers divisible by 3, 3 and 5
        } else if (oneToInt[i] % 3 == 0) {
            if (oneToInt[i] % 5 == 0) {
                oneToInt[i] = 'gi-oh';
            } else {
                oneToInt[i] = "gi";
            }

            // for numbers divisible by 5
        } else if (oneToInt[i] % 5 == 0) {
            oneToInt[i] = 'oh';

            // for numbers not divisible by 2, 3 or 5
        } else {

        }
    }
    console.log(oneToInt); // Use console.dir(oneToInt, { "maxArrayLength": null }); to see all array elements if over 100 elements
    return oneToInt;
}
yugioh(100);