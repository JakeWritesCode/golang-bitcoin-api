# bitcoin-price-api
This project is a bitcoin price API created as a learning exercise in golang.

## Loading an initial dataset.
To prevent having to hit an API endpoint forever to load data into the database its a good idea
to load some historical data from a file.
The file I've used can be found at: https://www.kaggle.com/datasets/mczielinski/bitcoin-historical-data
That should get you everything up to March 2021.

To use:
 - Download the dataset.
 - Run the compiled code with the flag `-parse path/to/your/file.csv`.

## Running the API
To run the API:
 - Run the compiled code with no flags. The server will start at port 10000.
