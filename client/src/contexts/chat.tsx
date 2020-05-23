import React, { useState, createContext, useEffect, useRef, useCallback } from 'react';
import ChatService, { Message } from '../services/chat';

export interface IChatContext {
  messages: Message[];
  isConnected: boolean;
  error?: string;
  connect: () => void;
  send: (message: string) => void;
}

export interface ChatContextProps {
  children: JSX.Element;
}

const ChatContext = createContext<IChatContext>(null);

ChatContext.displayName = 'ChatContext';

export function ChatContextProvider(props: ChatContextProps): JSX.Element {
  const { current: chatService } = useRef(new ChatService());
  const [messages, setMessages] = useState<Message[]>([]);
  const [isConnected, setConnected] = useState(chatService.isConnected.getValue());

  useEffect(() => {
    const streamSubs = chatService.stream.subscribe((next) => {
      setMessages([...messages, next]);
    });

    const isConnectedSubs = chatService.isConnected.subscribe((conn: boolean) => {
      setConnected(conn);
    });

    return () => {
      isConnectedSubs.unsubscribe();
      streamSubs.unsubscribe();
      chatService.disconnect();
    };
  }, []);

  const connect = useCallback((): void => {
    chatService.connect('ws://127.0.0.1:4200');
  }, []);

  const send = useCallback((message: string): void => {
    chatService.send(message);
  }, []);

  return (
    <ChatContext.Provider value={{
      isConnected,
      messages,
      connect,
      send
    }}>
      {props.children}
    </ChatContext.Provider>
  );
}

export default ChatContext;
