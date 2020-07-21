#!/bin/bash

# FILECOUNT=$(ls -l /linkedout/db | wc -l)
if [[ $(ls -l /linkedout/db | wc -l) -eq 0 ]]
then
    echo "Creating linkedout DB"
    mongo redirect --eval 'db.getSiblingDB("linkedout")'
    echo "creating Welcome collection"
    mongo redirect --eval 'db.createCollection("welcome")'
    echo "creating Experience collection"
    mongo redirect --eval 'db.createCollection("experience")'
    echo "creating certifications collection"
    mongo redirect --eval 'db.createCollection("certifications")'
    echo "creating skills collection"
    mongo redirect --eval 'db.createCollection("skills")'
    echo "creating accomplishments collection"
    mongo redirect --eval 'db.createCollection("accomplishments")'
    echo "creating recommendations collection"
    mongo redirect --eval 'db.createCollection("recommendations")'
    echo "creating projects collection"
    mongo redirect --eval 'db.createCollection("projects")'
    echo "creating faq collection"
    mongo redirect --eval 'db.createCollection("faq")'
    echo "creating contact collection"
    mongo redirect --eval 'db.createCollection("contact")'
    echo "creating settings collection"
    mongo redirect --eval 'db.createCollection("settings")'
    echo "creating metrics collection"
    mongo redirect --eval 'db.createCollection("metrics")'
    echo "Database and collections have been created."
fi