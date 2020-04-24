import React, { useContext } from 'react';
import ChatContext, { ChatContext as IChatContext } from '../../contexts/chat';
import Input from './input';

function Chat(): JSX.Element {
  const { isConnected, messages } = useContext(ChatContext);

  return (
    <div>
      <header>
        {isConnected ? 'Connected' : 'Not Connected'}
      </header>
      <ul>
        {messages?.map((message: string) => (
          <li>{message}</li>
        ))}
      </ul>
      <Input />
    </div>
  );
};

export default Chat;
