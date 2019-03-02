Build the image

    docker build -t go-pubg .

Run the image 

    docker run --rm -i -t -p 8080:8080 go-pubg

Run container with docker-compose
    
    docker-compose up
    
Stop container with docker-compose

    docker-compose down