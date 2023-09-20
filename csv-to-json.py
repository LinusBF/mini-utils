import sys
import csv
import json

# Based on https://www.geeksforgeeks.org/convert-csv-to-json-using-python/
def make_json(csvFilePath, jsonFilePath, keyColumn):
    data = {}
    with open(csvFilePath, encoding='utf-8') as csvf:
        csvReader = csv.DictReader(csvf, delimiter=';')
        for rows in csvReader:
            print(rows)
            key = rows[keyColumn]
            data[key] = rows
    with open(jsonFilePath, 'w', encoding='utf-8') as jsonf:
        jsonf.write(json.dumps(data, indent=4))

make_json(sys.argv[1], sys.argv[2], sys.argv[3])
