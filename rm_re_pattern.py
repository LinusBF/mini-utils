# This script filters out any line in a file that contains the specified RegEx pattern
# Run with `python rm_re_pattern.py input.txt output.txt`
import sys
import re

# Swedish SSN pattern
pattern = "((((18|19|20)[0-9][0-9])(((01|03|05|07|08|10|12)(0[1-9]|1[0-9]|2[0-9]|3[0-1]))|((04|06|09|11)(0[1-9]|1[0-9]|2[0-9]|30))|((02)(0[1-9]|1[0-9]|2[0-8])))|(((18|19|20)(04|08|12|16|20|24|28|32|36|40|44|48|52|56|60|64|68|72|76|80|84|88|92|96)(0229)))|(20000229)))(00[1-9]|0[1-9][0-9]|[1-9][0-9][0-9])[0-9]|((((18|19|20)[0-9][0-9])(((00|01|03|05|07|08|10|12)(6[0-9]|7[0-9]|8[0-9]|9[0-1]))|((04|06|09|11)(6[0-9]|7[0-9]|8[0-9]|90))|((02)(6[0-9]|7[0-9]|8[0-8])))|(((18|19|20)(04|08|12|16|20|24|28|32|36|40|44|48|52|56|60|64|68|72|76|80|84|88|92|96)(0289)))|(20000289)))(00[1-9]|0[1-9][0-9]|[1-9][0-9][0-9])[0-9]"


def test_line(line):
  return (re.search(pattern, line) is None)


with open(sys.argv[1], 'r') as f_in:
  with open(sys.argv[2], 'w') as f_out:
    for line in f_in:
      if test_line(line):
        f_out.write(line)
      else:
        print("Skipped line: " + line, end="")
