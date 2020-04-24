import React, { createContext, useState, useEffect } from 'react';
import io from 'socket.io-client';

export interface ChatContext {
  messages: any[];
  isConnected: boolean;
  sendMessage: (event:string, message: string) => void;
}

export interface ChatContextProps {
  children: JSX.Element;
}

const Chat = createContext<ChatContext | undefined>(undefined);

export function ChatContextProvider(props: ChatContextProps): JSX.Element {
  // Socket implementation must be written into a service
  // this is a test approach
  const [isConnected, setConnected] = useState(false);
  const [messages, setMessages] = useState([]);
  const socket = io(`ws://${process.env.WEB_SOCKET_HOST}:${process.env.WEB_SOCKET_PORT}`);

  useEffect(() => {
    socket.on('message', (data: any) => {
      console.log('Received message from server:', data);
      setMessages([...messages, data]);
    });

    socket.on('connect', () => {
      setConnected(socket.connected);
    });

    socket.on('connect', () => {
      setConnected(socket.connected);
    });
  }, []);

  const sendMessage = (event: string, message: string): void => {
    if (isConnected) {
      socket.emit(event, message);
    }
  };

  return (
    <Chat.Provider value={{
      isConnected,
      messages,
      sendMessage
    }}>
      {props.children}
    </Chat.Provider>
  );
}

export default Chat;
