import re


# implemented this in python because I felt like I was going insane
# trying to troubleshoot golang. Turned out to be useful because
# golang didn't support positive lookaheads in regex

def parseDigitOrName(match: str):
    if match.isdigit():
        return int(match)
    elif match == "one":
        return 1
    elif match == "two":
        return 2
    elif match == "three":
        return 3
    elif match == "four":
        return 4
    elif match == "five":
        return 5
    elif match == "six":
        return 6
    elif match == "seven":
        return 7
    elif match == "eight":
        return 8
    elif match == "nine":
        return 9
        

with open("./input.txt") as file:
    count = 0
    for line in file:
        matches = re.findall("(?=([0-9]|one|two|three|four|five|six|seven|eight|nine))", line,)
        firstDigit = None
        lastDigit = None
        for match in matches:
            lastDigit = parseDigitOrName(match)
            if firstDigit == None:
                firstDigit = lastDigit
        count += firstDigit*10 + lastDigit

    print("The count is " + str(count))
