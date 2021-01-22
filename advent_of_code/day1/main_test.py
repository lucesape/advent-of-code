import unittest

from .main import check_sum, cast_to_int, make_pairs, main


class TestDay1(unittest.TestCase):
    def test_check_sum(self) -> None:
        input_data: Tuple[int, int] = 2019, 1
        expected_output: Tuple[int, int, int] = 2019, 1, 2019
        self.assertEqual(check_sum(*input_data), expected_output)

    def test_make_pairs(self) -> None:
        input_data: List[int] = [1, 2, 3]
        expected_output: List[List[int]] = [[1, 2], [1, 3], [2, 3]]
        self.assertEqual(make_pairs(data=input_data), expected_output)
     