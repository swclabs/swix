from tools import tools

# IMDB Formula:
# Weighted Rating (WR) = ((v/(v+m))xR)+((m/(v+m))xC)
# v is the number of votes for the movie;
# m is the minimum votes required to be listed in the chart;
# R is the average rating of the movie;
# C is the mean vote across the whole report.


class Recommender(object):
    def __init__(self):
        self.__version = '0.0.1'
        self.tools = tools

    def content_base_filtering(self):
        print(
            f"[RECOMMENDER] version: {self.__version} Content Base Filtering process ...")
        return

    def collaborative_filtering(self):
        return

    def export(self):
        return


recommender = Recommender()
