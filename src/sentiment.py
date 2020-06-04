from textblob import TextBlob


class JANLP:
    def __init__(self, feedbackFile):
        self.feedbackFile = feedbackFile
        self.lines = ""

        with open(self.feedbackFile, "r") as f: self.lines = f.readlines()
    def getSentiment(self):
        print(self.lines)


if __name__ == "__main__":
    Jastics = JANLP("/tmp/jm_feedback")
    Jastics.getSentiment()
