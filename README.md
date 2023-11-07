# Go Programming Assignment - Ideal Gas Law, Decryption, and Parity

 This repository contains a Go programming assignment that focuses on implementing and testing essential functions related to the ideal gas law, decryption, and parity checking. This assignment was part of my CS 311 course and was designed to help practice Go programming fundamentals.

## Table of Contents

- [Installation](#installation)
- [Functions](#functions) 

## Installation

Follow these steps to install and run the code on your local machine:

1. Install Go: 
   If Go is not already installed on your machine, you need to install it. Visit the official Go    
   Downloads page at https://golang.org/dl/ and download the installer for your operating system. Follow the installation instructions provided for your OS.
2. Clone the Repository:
   Clone this GitHub repository to your local machine using the following command:
   git clone https://github.com/RudyTheCoder/idealGasLaw-Decryption-Parity.git 
3. Navigate to the Project Directory: 
   Change your current working directory to the project folder using the cd command:
   cd idealGasLaw-Decryption-Parity
4. Run the Main Program:
   To execute the main program, use the following command:
   go run hw1.go 
   This command will run the Go program located in hw1.go, and you should see the output demonstrating the functionality of the three functions (CalcPressure, Decode, and OddParity) as specified in the assignment.


# Functions

CalcPressure

The CalcPressure function calculates the pressure of an ideal gas using the ideal gas law. It takes three parameters: volume (v), molarity (n), and temperature (t), all of type float64. The function returns the pressure in pascals (Pa) and is calculated using the formula: p = n * R * t / v, where R is the Ideal Gas Constant (8.3144598 J/(mol·K)). An example result for CalcPressure(1.0, 1.0, 298.15) should return approximately 2478.9562 pascals.

Decode

The Decode function decrypts a message using a provided code map, transforming an encrypted message into its original form. It takes two parameters: the encrypted message (e) and a map (c) of byte to byte, representing the code for decryption. The function returns the decrypted message as a slice of bytes ([]byte). To decrypt a character 'a', the function replaces it with the corresponding byte from the code map (e.g., c['a']). If the code map does not contain a character, it leaves it unchanged. For example, Decode([]byte("hello"), c) with c['e'] = 'u', c['h'] = 'f', c['l'] = 'n', c['o'] = 'y' should return the slice of bytes ['f', 'u', 'n', 'n', 'y']`.

OddParity

The OddParity function takes a slice of integers and returns two boolean values. The first boolean indicates whether the list has an odd number of 1's (true for odd, false for even), and the second boolean indicates whether all values in the list are either 0's or 1's.