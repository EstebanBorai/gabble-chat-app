import React, { useState, useContext } from 'react';
import ChatContext, { ChatContext as IChatContext } from '../../contexts/chat';

function Input(): JSX.Element {
  const [inputValue, setInputValue] = useState('');
  const { isConnected, sendMessage } = useContext<IChatContext>(ChatContext);

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>): void => {
    const { target: { value } } = event;

    setInputValue(value);
  }

  const handleSend = (event: React.FormEvent<HTMLFormElement>): void => {
    event.preventDefault();
    sendMessage(inputValue);
    setInputValue('');
  }

  return (
    <footer>
      <form action="" onSubmit={handleSend}>
        <label htmlFor="chat-input">
          <input
            name="chat-input"
            type="text"
            value={inputValue}
            disabled={!isConnected}
            onChange={handleChange}
          />
        </label>
        <button type="submit">Send</button>
      </form>
    </footer>
  );
}

export default Input;
