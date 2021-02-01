class Policy(object):

    def __init__(
        self, 
        lower_limit: int = 0, 
        upper_limit: int = 0, 
        character: str = "", 
        password: str = ""
    ) -> None:
        self.lower_limit = lower_limit
        self.upper_limit = upper_limit
        self.character = character
        self.password = password

    def is_password_valid(self) -> bool:
        return self.lower_limit <= self.password.count(self.character) <= self.upper_limit



def make_policy(line: str) -> Policy:

    p = Policy()
    split_line = line.split()

    p.lower_limit = int(split_line[0].split("-")[0])
    p.upper_limit = int(split_line[0].split("-")[1])
    p.character = split_line[1][0]
    p.password = split_line[-1]

    return p

def main(file: str):
    pass
