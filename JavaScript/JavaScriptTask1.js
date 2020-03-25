//A program to display all odd numbers from 1 to 200 (inclusive)

const name = "Emeka Elo-Chukwuma";

const courses = ['HTML', 'CSS', 'JavaScript', 'GO', 'Design'];
const oddnumbers = [];
//Display name and courses
console.log(name);
console.log(courses);

console.log("Number of courses is " + courses.length + " which is an odd number.")

//find and display all odd numbers from 1 to 200 (inclusive)
for (i = 0; i < 200; i++) {
    if (i % 2 != 0) {
        oddnumbers.push(i);
    }
}
console.log(oddnumbers)