# ChitChat

This is the simple forum application written with Go for the book "Go Web Programming" from Manning Publication Co. The source code for this application is the basis for the Chapter 2 - Go ChitChat. 

However the code is a reference and is not a 1 - 1 match with the code in the book. There are portions of the code that is more detailed in here than in Chapter 2 (which is a simplified version of the source code here).

Some differences include:

* This version of ChitChat is configurable with a `config.json` file
* This version is able to log to a `chitchat.log` file
* There are test files in this code
* All structs are fully fleshed out (in the book chapter, they are only implied)
* Some of the functions are placed as methods for the struct types instead of being a part of the package
