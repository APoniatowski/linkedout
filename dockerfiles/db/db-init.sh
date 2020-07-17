#!/bin/bash

FILECOUNT=$(ls -l /linkeout/db | wc -l)
if [ FILECOUNT -eq 0 ]
then
    echo "Creating linkedout DB"
    mongo --shell db.getSiblingDB('linkedout')
    echo "creating Welcome collection"
    mongo --shell db.createCollection("welcome")
    echo "creating Experience collection"
    mongo --shell db.createCollection("experience")
    echo "creating certifications collection"
    mongo --shell db.createCollection("certifications")
    echo "creating skills collection"
    mongo --shell db.createCollection("skills")
    echo "creating accomplishments collection"
    mongo --shell db.createCollection("accomplishments")
    echo "creating recommendations collection"
    mongo --shell db.createCollection("recommendations")
    echo "creating projects collection"
    mongo --shell db.createCollection("projects")
    echo "creating faq collection"
    mongo --shell db.createCollection("faq")
    echo "creating contact collection"
    mongo --shell db.createCollection("contact")
    echo "creating settings collection"
    mongo --shell db.createCollection("settings")
    echo "creating metrics collection"
    mongo --shell db.createCollection("metrics")
    echo "Database and collections have been created."
else
    echo "Directory not empty..."
    echo "Assuming DB is stored here."
    echo "Loading"
fi