# BuyIt - Live Search with Beego and Elasticsearch

## Overview
BuyIt is a web application implemented using the Beego framework and Go programming language, featuring a live search functionality where data comes from Elasticsearch.

## Features
- Real-time search functionality
- Elasticsearch integration
- Built with Beego framework
- RESTful API endpoints

## Prerequisites
- Go 1.16 or higher
- Beego framework
- Elasticsearch 7.x
- Git

## Installation
```bash
# Clone the repository
git clone https://github.com/uzzalcse/buyit.git

# Navigate to project directory
cd buyit

# Install dependencies
go mod tidy
```

## Configuration
1. Set up Elasticsearch connection in `conf/app.conf`
2. Configure database settings if required
3. Update any environment-specific configurations

## Running the Application
```bash
# Build the applicaiton
docker compose build
# Run the server
docker compose up
# Stop the server
docker compose down
```
The application will be available at `http://localhost:8080`


## Sample Configuration
Here's a sample `app.conf` configuration:

```ini
# Application Settings
appname = buyit
httpport = 8080
runmode = dev

# Elasticsearch Configuration
ES_LOCAL_API_KEY = your-api-key
ES_LOCAL_URL = http://elasticsearch:9200

# without docker
ES_LOCAL_URL = http://localhost:9200

```

## Contributing
Pull requests are welcome. For major changes, please open an issue first.