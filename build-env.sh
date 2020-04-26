SERVER_HOST=""
SERVER_PORT=""
DEFAULT_SOCKET_PATH=""

echo "Enter the server host: "
read SERVER_HOST

echo "Enter the server port: "
read SERVER_PORT

echo "Enter the default socket path: "
read DEFAULT_SOCKET_PATH

echo -e "SERVER_HOST=$SERVER_HOST" \
  "\nSERVER_PORT=$SERVER_PORT" \
  "\nDEFAULT_SOCKET_PATH=$DEFAULT_SOCKET_PATH" > .env
