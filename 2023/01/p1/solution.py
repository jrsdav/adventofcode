input = open("../input.txt", "r")

def extract_callibration_value(input):
  
	results = []
  	
	for s in input:
		numbers = [char for char in s if char.isdigit()]

		if numbers:
			first = numbers[0]
			last = numbers[-1]

			results.append(int(first + last)) 
 
	return results
	
total = sum(extract_callibration_value(input))

print(total)
		