from itertools import product, permutations

# Define the letter groups
guess1 = []
guess2 = ["mouch"]
guess3 = ["blank"]
guess4 = ["wefts"]
guess5 = []
guess6 = []
guess7 = []
guess8 = []

res1 = [0, 0, 5]
res2 = [2, 1, 2]
res3 = [0, 0, 5]
res4 = [0, 2, 3]
res5 = []
res6 = []
res7 = []
res8 = []

alphabet = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z']

base_invalid = ['j', 'q', 'x', 'z']

inval_first_letter = [] + base_invalid
inval_second_letter = [] + base_invalid
inval_third_letter = [] + base_invalid
inval_fourth_letter = [] + base_invalid
inval_fifth_letter = [] + base_invalid

list_of_lists = [inval_first_letter, inval_second_letter, inval_third_letter, inval_fourth_letter, inval_fifth_letter]

solution_letters = []

for i in range(0,5):
    if res1[2] == 5: # letters are not in solution
        inval_first_letter.append(guess1[i])
        inval_second_letter.append(guess1[i])
        inval_third_letter.append(guess1[i])
        inval_fourth_letter.append(guess1[i])
        inval_fifth_letter.append(guess1[i])
    if res1[0] == 0 and res1[2] != 5: # letters cannot be where they are
        list_of_lists[i].append(guess1[i])
    if res1[2] == 0: # letters are in solution
        solution_letters.append(guess1[i])
        if res1[0] == 0: # letters cannot be where they are, letters are in solution
            list_of_lists[i].append(guess1[i])

    # add logic for green != 0, yellow == 0, red != 0
    # letters are where they are, or not in solution

        

# Generate all base combinations
base_combinations = list(product(group1, group2, group3, group4))

valid_words = set()
invalid_words = set()

# Create permutations with the constraints
for combo in base_combinations:
    for perm in permutations(combo):
        if perm[0] not in inval_first_letter and perm[1] not in inval_second_letter and perm[2] not in inval_third_letter and perm[3] not in inval_fourth_letter and perm[4] not in inval_fifth_letter:
            valid_words.add(''.join(perm))
        else: invalid_words.add(''.join(perm))

# Load English words from nltk
import nltk
from nltk.corpus import words

# Filter for 5-letter real English words
english_words = set(word.lower() for word in words.words() if len(word) == 5)

# Intersect to find valid real words
real_words = valid_words & english_words

# Output
print("Total invalid words:", len(invalid_words))
print("Total valid words:", len(valid_words))
print(f"Total valid combinations (after constraints): {len(valid_words)}")
print(f"Number of real English words: {len(real_words)}")
print("Sample real words:", list(real_words))

test_word = "clear"
if test_word in real_words:
    print("real words")
if test_word in english_words:
    print("english words")
if test_word in valid_words:
    print("valid words")
if test_word in invalid_words:
    print("invalid words")