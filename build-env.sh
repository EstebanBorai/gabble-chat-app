SERVER_HOST=""
SERVER_PORT=""

echo "Enter the server host: "
read SERVER_HOST

echo "Enter the server port: "
read SERVER_PORT

echo -e "SERVER_HOST=$SERVER_HOST" \
  "\nSERVER_PORT=$SERVER_PORT" > .env
