# Jastics
Tool to statistically analyze the user feedback of the experimental application TextBreakerAlpha on jastonmatter.com

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

What things you need to install for the software to work properly

```
Go
python3
pip3 install textblob
```

## Deployment

Deploy the program by running the following
```
go run ./main.go
python3 src/sentiment.py
```

## Built With

* [TextBlob](https://textblob.readthedocs.io/en/dev/) - Sentiment Analysis 
* [Python 3.6](https://www.python.org/downloads/release/python-360/) - To further analyze especially the feedback text
* [Go](https://golang.org/) - Provides fast processing, to fetch the data and create graphs 


## Authors

* **Johannes Leppäkorpi** - *Entire project* - [Johk3](https://github.com/Johk3)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

* Be sure to give your own spin to the source code, and have fun!
