import unittest

from .helpers import read_file, read_file_line_by_line

class TestHelpers(unittest.TestCase):
    def test_read_file(self) -> None:
        self.assertEqual(
            read_file(filepath="code/helpers/helpers.txt"),
            "this\nis\na\ntest\n"
        )
    def test_read_file_line_by_line(self) -> None:
        self.assertEqual(
            read_file_line_by_line(filepath="code/helpers/helpers.txt"),
            ["this", "is", "a", "test"]
        )
