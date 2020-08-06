if [ "$(docker ps -q -f name=jewerly-api)" ]; then
    if [ ! "$(docker ps -aq -f status=exited -f name=jewerly-api)" ]; then
        docker stop jewerly-api
    fi
fi

#docker run -e HOST -e ACCESS_KEY -e SECRET_KEY --rm -d --publish 8000:8000 --network jewerly --name jewerly-api jewerly:0.1
docker run -e HOST -e ACCESS_KEY -e SECRET_KEY --rm -d --publish 8000:8000 --name jewerly-api jewerly-api:0.1