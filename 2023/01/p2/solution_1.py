import re

with open("../input.txt", "r") as f:

	input_data = f.readlines()

def word_to_number(input_data):
	results = []
	number_dict = {
		'oneight': '18',
		'threeight': '38',
		'fiveight': '58',
		'nineight': '98',
		'twone': '21',
		'sevenine': '79',
		'eighthree': '83',
		'eightwo': '82',
		'one': '1',
		'two': '2',
		'three': '3',
		'four': '4',
		'five': '5',
		'six': '6',
		'seven': '7',
		'eight': '8',
		'nine': '9',
	}

	regex = re.compile("(%s)" % "|".join(map(re.escape, number_dict.keys())))

	for item in input_data:
		results.append(regex.sub(lambda mo: number_dict[mo.string[mo.start():mo.end()]], item))

	return results


def extract_callibration_value(input_data):
  
	results = []
  	
	for s in input_data:
		numbers = [char for char in s if char.isdigit()]

		if numbers:
			first = numbers[0]
			last = numbers[-1]

			results.append(int(first + last))
 
	return results

number_conversion = word_to_number(input_data)
calibration_values = extract_callibration_value(number_conversion)

total = sum(calibration_values)

print(total)