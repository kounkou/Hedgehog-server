#!/bin/bash

# Wait for MongoDB to start
echo "Waiting for MongoDB to start..."
sleep 5

# Import the data.json file into the questionDB database and questions collection
mongoimport --host localhost --port 27017 --db questionDB --collection programmingQuestions --file /data/data.json


echo "Database and collection initialized with data!"
