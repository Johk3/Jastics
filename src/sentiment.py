from textblob import TextBlob


class JANLP:
    """JANLP, provides statistical analysis for jastonmatter.com"""
    def __init__(self, feedback_file):
        self.feedback_file = feedback_file
        self.lines = []

        with open(self.feedback_file, "r") as f:
            self.lines = f.read().split("-+--++--+-+--++--+--+-+--+")

    def get_sentiment(self):
        """Produce the sentiment values for each feedback given"""
        polarityAvg = 0
        subjectivityAvg = 0

        for line in self.lines:
            context = TextBlob(line)
            polarityAvg += context.sentiment.polarity
            subjectivityAvg += context.sentiment.subjectivity

        polarityAvg = polarityAvg/len(self.lines)*1000
        subjectivityAvg = subjectivityAvg/len(self.lines)*1000

        print(f"{round(polarityAvg, 2)}% Positive, and {round(subjectivityAvg, 2)}% Subjective")


if __name__ == "__main__":
    Jastics = JANLP("/tmp/jm_feedback")
    Jastics.get_sentiment()
