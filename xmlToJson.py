import sys
import json
from xmljson import parker
import xml.etree.ElementTree as etree

with open(sys.argv[1], 'r') as xml_file:
    xml_tree = etree.parse(xml_file)

json_data = json.dumps(parker.data(xml_tree.getroot()), ensure_ascii=False, indent=2)

with open(sys.argv[2], 'w') as out:
    out.write(json_data)
