import unittest

from advent_of_code.day2.part1.main import Policy, make_policy

class TestDay2Part1(unittest.TestCase):

    def test_make_policy(self):
        input_data: str = "1-4 q: abcqrs"
        expected_output: Policy = Policy(lower_limit=1, upper_limit=4, character="q", password="abcqrs")
        self.assertEqual(make_policy(line=input_data), expected_output)
