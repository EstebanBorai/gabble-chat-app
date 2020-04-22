import io from 'socket.io-client';
import { useState, useEffect } from 'react';

interface UseSocketIO {
  messages: any[];
  isConnected: boolean;
  sendMessage: (message: string) => void;
}

export const useSocketIO = (url: string): UseSocketIO => {
  const [isConnected, setConnected] = useState(false);
  const [messages, setMessages] = useState([]);

  const socket = io(url);

  socket.on('message', (data: any) => {
    setMessages([...messages, data]);
  });

  useEffect(() => {
    setConnected(socket.connected);
  }, [socket.connected]);

  const sendMessage = (message: string): void => {
    if (message && isConnected) {
      socket.send(message);
    }
  };

  return {
    messages,
    isConnected,
    sendMessage
  };
}
