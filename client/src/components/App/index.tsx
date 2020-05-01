import React, { useRef } from 'react';
import Main from '../Main';
import { ChatContextProvider } from '../../contexts/chat';
import ChatService, { ChatServiceInterface } from '../../services/ChatService';

function App(): JSX.Element {
  const chatService = useRef<ChatServiceInterface>(new ChatService());

  return (
    <ChatContextProvider service={chatService.current}>
      <Main />
    </ChatContextProvider>
  )
}

export default App;
