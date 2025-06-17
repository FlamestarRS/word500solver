from itertools import product, permutations

# Define the letter groups
group1 = []
group2 = []
group3 = []
group4 = []

alphabet = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z']

base_invalid = ['j', 'q', 'x', 'z']

inval_first_letter = [] + base_invalid
inval_second_letter = [] + base_invalid
inval_third_letter = [] + base_invalid
inval_fourth_letter = [] + base_invalid
inval_fifth_letter = [] + base_invalid

solution_letters = []

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