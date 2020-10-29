/*
* This code piece demonstrates how to create "Class" for javascript
* */

var Animal = {
    New: function() {
        var animal = {};
        animal.name = "animal";
        animal.say = function(){ console.log("Animal says: ...") };
        animal.sleep = function () {
            console.log("The " + animal.name + " sleeps")
        };
        return animal
    }

};


var Fox = {
    age: 1,  // Instance variable

    New: function () {
        var fox = Animal.New();
        fox.name = "Fox";
        fox.say = function(){ console.log("Fox says: Ling Ling")};
        var dna = "ATCG";  // Private variable
        fox.dna = function(){ return dna };
        return fox
    }
};

var fox = Fox.New();
fox.sleep();  // Call inherited method
fox.say();  // Call its own method
console.log("Fox DNA: " + fox.dna());
console.log("Age of the fox:" + fox.age);  // Not available
console.log("Shared age of Fox:" + Fox.age);
Fox.age += 1;
console.log("Edited age:" + Fox.age);

