declare namespace NodeJS {
  export interface ProcessEnv {
    NODE_ENV: 'development' | 'production';

    /**
     * The host of the SocketIO server
     * Eg: 127.0.0.1
     */
    WEB_SOCKET_HOST: string;

    /**
     * The port of the SocketIO server
     * Eg: 8000
     */
    WEB_SOCKET_PORT: string;
  }
}
