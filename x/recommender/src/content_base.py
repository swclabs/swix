from recommender import recommender

if __name__ == '__main__':
    recommender.tools.read_csv("data/movies_metadata.csv")
    recommender.content_base_filtering()
    print(recommender.tools.metadata.head(8))
