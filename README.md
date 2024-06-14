# KATE McELHANEY 
### TAKE HOME TEST FOR FETCH 
## Getting Started:
1. Ensure you have GO installed
2. Open your terminal and make a new directory 
3. Run the command in your terminal to open VSCode with your new directory and Clone GitHub Repository: https://github.com/K8MacEl/Fetch-Take-Home.git
4. In your terminal run "go mod init receipts"
5. Then run "go get github.com/google/uuid" to install the UUID package
6. Start the server by typing in your terminal "go run main.go"
7. The terminal should run and you should recieve a prompt to go to server 8080, to do this in your browser enter: "http://localhost:8080/receipts/process"

## Testing the Code:
1. To test the code you can use the data entered in main_test.go
2. In the terminal run "go test"


## Technologies Used:
This was written entirely in Go

## Function of Code

This allows for a JSON object to get a set number of points based on the following rules:

* One point for every alphanumeric character in the retailer name.
* 50 points if the total is a round dollar amount with no cents.
* 25 points if the total is a multiple of 0.25.
* 5 points for every two items on the receipt.
* If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
* 6 points if the day in the purchase date is odd.
* 10 points if the time of purchase is after 2:00pm and before 4:00pm.


