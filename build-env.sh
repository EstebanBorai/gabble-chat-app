SERVER_HOST=""
SERVER_PORT=""
LOG_LEVEL=""

echo "Enter the server host: "
read SERVER_HOST

echo "Enter the server port: "
read SERVER_PORT

echo "Enter the log level: "
echo "0 = None, 1 = Error, 2 = Warning, 3 = Info"
read LOG_LEVEL

echo -e "SERVER_HOST=$SERVER_HOST" \
  "\nSERVER_PORT=$SERVER_PORT" \
  "\nLOG_LEVEL=$LOG_LEVEL" > .env
