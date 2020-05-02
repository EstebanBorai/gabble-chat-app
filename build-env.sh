SERVER_HOST=""
SERVER_PORT=""
LOG_LEVEL=""
CLIENT_HOST=""
CLIENT_PORT=""

echo "Enter the server host: "
read SERVER_HOST

echo "Enter the server port: "
read SERVER_PORT

echo "Enter the client host: "
read CLIENT_HOST

echo "Enter the client port: "
read CLIENT_PORT

echo "Enter the log level: "
echo "0 = None, 1 = Error, 2 = Warning, 3 = Info"
read LOG_LEVEL

echo -e "SERVER_HOST=$SERVER_HOST" \
  "\nSERVER_PORT=$SERVER_PORT" \
  "\nCLIENT_HOST=$CLIENT_HOST" \
  "\nCLIENT_PORT=$CLIENT_PORT" \
  "\nLOG_LEVEL=$LOG_LEVEL" > .env
