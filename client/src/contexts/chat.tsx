import React, { Context, useState, createContext, useEffect } from 'react';
import { Author, Message, ChatServiceInterface } from '../services/ChatService';

export interface ChatContextInterface {
  author?: Author;
  messages: Message[];
  error?: string;
  join: (username: string) => void;
  send: (message: string) => void;
}

export interface ChatContextProps {
  service: ChatServiceInterface;
  children: JSX.Element;
}

const ChatContext = createContext<Context<ChatContextInterface>>(null);

ChatContext.displayName = 'ChatContext';

export function ChatContextProvider(props: ChatContextProps): JSX.Element {
  const [messages, setMessages] = useState<Message[]>([]);
  const [author, setAuthor] = useState<Author>(null);

  useEffect(() => {
    props.service.messages.subscribe((next) => {
      setMessages([...next]);
    });

    props.service.author.subscribe((next) => {
      setAuthor(next);
    });

    return () => {
      props.service.messages.unsubscribe();
      props.service.author.unsubscribe();
    };
  }, []);

  return (
    <ChatContext.Provider value={{
      author,
      messages,
      join: props.service.join,
      send: props.service.send
    }}>
      {props.children}
    </ChatContext.Provider>
  );
}

export default ChatContext;
