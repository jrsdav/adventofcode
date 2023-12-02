number_dict = {
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

def get_first_number(line):
	for i in range(len(item)):
		substring = item[:i+1]
		for word, digit in number_dict.items():
			if digit in substring:
				return digit
			if word in substring:
				return number_dict[word]
			

def get_last_number(line):
	for i in range(len(item), 0, -1):
		substring = item[i-1:]
		for word, digit in number_dict.items():
			if digit in substring:
				return digit
			if word in substring:
				return number_dict[word]

results = []

with open("../input.txt", "r") as f:
	input_data = f.readlines()
	for item in input_data:
		first_number = get_first_number(item)
		last_number = get_last_number(item)
		results.append(int(first_number + last_number))	
	
print(sum(results))

def test_get_first_number_int():
	item = "4nineeightseven2"
	assert get_first_number(item) == 4

def test_get_first_number_string():
	item = "two1nine"
	assert get_first_number(item) == 2

def test_get_last_number_int():
	item = "4nineeightseven2"
	assert get_last_number(item) == 2

def test_get_last_number_string():
	item = "two1nine"
	assert get_last_number(item) == 9	

def test_edge_cases():
	item = "3fiveeightoneightg"
	assert get_last_number(item) == 8